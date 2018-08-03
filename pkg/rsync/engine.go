package rsync

import (
	"bufio"
	"bytes"
	"crypto/sha1"
	"io"
	"math"

	"github.com/pkg/errors"
)

// BlockHash represents a pair of weak and strong hash for a base block.
type BlockHash struct {
	// Weak is the weak hash for the block.
	Weak uint32
	// Strong is the strong hash for the block.
	Strong [sha1.Size]byte
}

// Signature represents an rsync base signature. It encodes the block size used
// to generate the signature, the size of the last block in the signature (which
// may be smaller than a full block), and the hashes for the blocks of the file.
type Signature struct {
	// BlockSize is the block size used to compute the signature.
	BlockSize uint64
	// LastBlockSize is the size of the last block in the signature.
	LastBlockSize uint64
	// Hashes are the hashes of the blocks in the base.
	Hashes []BlockHash
}

// isZeroValue indicates whether or not the Operation has its zero value.
func (s Signature) isZeroValue() bool {
	return s.BlockSize == 0 && s.LastBlockSize == 0 && len(s.Hashes) == 0
}

// ensureValid verifies that signature invariants are respected.
func (s Signature) ensureValid() error {
	// If the block size is 0, then the last block size should also be 0 and
	// there shouldn't be any hashes.
	if s.BlockSize == 0 {
		if s.LastBlockSize != 0 {
			return errors.New("block size of 0 with non-0 last block size")
		} else if len(s.Hashes) != 0 {
			return errors.New("block size of 0 with non-0 number of hashes")
		}
		return nil
	}

	// If block size is non-0, then the last block size should be non-0 but less
	// than or equal to the block size.
	if s.LastBlockSize == 0 {
		return errors.New("non-0 block size with last block size of 0")
	} else if s.LastBlockSize > s.BlockSize {
		return errors.New("last block size greater than block size")
	}

	// If the block size is non-0, then a non-zero number of blocks should have
	// been hashed.
	if len(s.Hashes) == 0 {
		return errors.New("non-0 block size with no block hashes")
	}

	// Success.
	return nil
}

// Operation represents an rsync operation, which can be either a data operation
// or a block operation.
type Operation struct {
	// Data contains data for data operations. If its length is 0, the operation
	// is assumed to be a non-data operation. Operation transmitters and
	// receivers may thus treat a length-0 buffer as semantically equivalent to
	// a nil buffer and utilize that fact to efficiently re-use buffer capacity,
	// e.g. by truncating the buffer and doing a gob receive into it.
	Data []byte
	// Start is the 0-indexed starting block for block operations.
	Start uint64
	// Count is the number of blocks for block operations.
	Count uint64
}

// isZeroValue indicates whether or not the Operation has its zero value.
func (o Operation) isZeroValue() bool {
	return len(o.Data) == 0 && o.Start == 0 && o.Count == 0
}

// ensureValid verifies that operation invariants are respected.
func (o Operation) ensureValid() error {
	if len(o.Data) > 0 {
		if o.Start != 0 {
			return errors.New("data operation with non-0 block start index")
		} else if o.Count != 0 {
			return errors.New("data operation with non-0 block count")
		}
	} else if o.Count == 0 {
		return errors.New("block operation with 0 block count")
	}
	return nil
}

const (
	// minimumOptimalBlockSize is the minimum block size that will be returned
	// by OptimalBlockSizeForBaseLength. It has to be chosen so that it is at
	// least a few orders of magnitude larger than the size of a BlockHash.
	minimumOptimalBlockSize = 1 << 10
	// maximumOptimalBlockSize is the maximum block size that will be returned
	// by OptimalBlockSizeForBaseLength. It mostly just needs to be bounded by
	// what can fit into a reasonably sized in-memory buffer, particularly if
	// multiple rsync engines are running. maximumBlockSize also needs to be
	// less than or equal to (2^32)-1 for the weak hash algorithm to work.
	maximumOptimalBlockSize = 1 << 16
	// DefaultBlockSize is the default block size that will be used if a zero
	// value is passed into Engine.Signature for the blockSize parameter.
	DefaultBlockSize = 1 << 13
	// DefaultMaximumDataOperationSize is the default maximum data size
	// permitted per operation. The optimal value for this isn't at all
	// correlated with block size - it's just what's reasonable to hold
	// in-memory and pass over the wire in a single transmission. This value
	// will be used if a zero value is passed into Engine.Deltafy for the
	// maxDataOpSize parameter.
	DefaultMaximumDataOperationSize = 1 << 16
)

// OptimalBlockSizeForBaseLength uses a simpler heuristic to choose a block
// size based on the base length. It starts by choosing the optimal block length
// using the formula given in the rsync thesis. It then enforces that the block
// size is within a sensible range.
// TODO: Should we add rounding to "nice" values, e.g. the nearest multiple of
// 1024 bytes? Would this improve read throughput?
func OptimalBlockSizeForBaseLength(baseLength uint64) uint64 {
	// Compute the optimal block length (see the rsync thesis) assuming one
	// change per file.
	result := uint64(math.Sqrt(24.0 * float64(baseLength)))

	// Ensure it's within the allowed range.
	if result < minimumOptimalBlockSize {
		result = minimumOptimalBlockSize
	} else if result > maximumOptimalBlockSize {
		result = maximumOptimalBlockSize
	}

	// Done.
	return result
}

// OptimalBlockSizeForBase is a convenience function that will determine the
// optimal block size for a base that implements io.Seeker. It calls down to
// OptimalBlockSizeForBaseLength. After determining the base's length, it will
// attempt to reset the base to its original position.
func OptimalBlockSizeForBase(base io.Seeker) (uint64, error) {
	if currentOffset, err := base.Seek(0, io.SeekCurrent); err != nil {
		return 0, errors.Wrap(err, "unable to determine current base offset")
	} else if currentOffset < 0 {
		return 0, errors.Wrap(err, "seek return negative starting location")
	} else if length, err := base.Seek(0, io.SeekEnd); err != nil {
		return 0, errors.Wrap(err, "unable to compute base length")
	} else if length < 0 {
		return 0, errors.New("seek returned negative offset")
	} else if _, err = base.Seek(currentOffset, io.SeekStart); err != nil {
		return 0, errors.Wrap(err, "unable to reset base")
	} else {
		return OptimalBlockSizeForBaseLength(uint64(length)), nil
	}
}

// OperationTransmitter transmits an operation. Operation data buffers are
// re-used between calls to the transmitter, so the transmitter should not
// return until it has either transmitted the data buffer (if any) or copied it
// for later transmission.
type OperationTransmitter func(Operation) error

// Engine provides rsync functionality without any notion of transport. It is
// designed to be re-used to avoid heavy buffer allocation.
type Engine struct {
	// buffer is a re-usable buffer that will be used for reading data and
	// setting up operations.
	buffer []byte
	// targetReader is a re-usable bufio.Reader that will be used for delta
	// creation operations.
	targetReader *bufio.Reader
}

// NewEngine creates a new rsync engine.
func NewEngine() *Engine {
	return &Engine{
		targetReader: bufio.NewReader(nil),
	}
}

// bufferWithSize lazily allocates the engine's internal buffer, ensuring that
// it is the required size. The capacity of the internal buffer is retained
// between calls to avoid allocations if possible.
func (e *Engine) bufferWithSize(size uint64) []byte {
	// Check if the buffer currently has the required capacity. If it does, then
	// use that space. Note that we're checking *capacity* - you're allowed to
	// slice a buffer up to its capacity, not just its length. Of course, if you
	// don't own the buffer, you could run into problems with accessing data
	// outside the slice that was given to you, but this buffer is completely
	// internal, so that's not a concern.
	if uint64(cap(e.buffer)) >= size {
		return e.buffer[:size]
	}

	// If we couldn't use our existing buffer, create a new one, but store it
	// for later re-use.
	e.buffer = make([]byte, size)
	return e.buffer
}

const (
	// m is the weak hash modulus. I think they now recommend that it be the
	// largest prime less than 2^16, but this value is fine as well.
	m = 1 << 16
)

// weakHash computes a fast checksum that can be rolled (updated without full
// recomputation). This particular hash is detailed on page 55 of the rsync
// thesis. It is not theoretically optimal, but it's fine for our purposes.
func (e *Engine) weakHash(data []byte, blockSize uint64) (uint32, uint32, uint32) {
	// Compute hash components.
	var r1, r2 uint32
	for i, b := range data {
		r1 += uint32(b)
		r2 += (uint32(blockSize) - uint32(i)) * uint32(b)
	}
	r1 = r1 % m
	r2 = r2 % m

	// Compute the hash.
	result := r1 + m*r2

	// Done.
	return result, r1, r2
}

// rollWeakHash updates the checksum computed by weakHash by adding and removing
// a byte.
func (e *Engine) rollWeakHash(r1, r2 uint32, out, in byte, blockSize uint64) (uint32, uint32, uint32) {
	// Update components.
	r1 = (r1 - uint32(out) + uint32(in)) % m
	r2 = (r2 - uint32(blockSize)*uint32(out) + r1) % m

	// Compute the hash.
	result := r1 + m*r2

	// Done.
	return result, r1, r2
}

// strongHash computes a slow but strong hash for a block of data.
func (e *Engine) strongHash(data []byte) [sha1.Size]byte {
	return sha1.Sum(data)
}

// Signature computes the signature for a base stream. If the provided block
// size is 0, this method will attempt to compute the optimal block size (which
// requires that base implement io.Seeker), and failing that will fall back to a
// default block size.
func (e *Engine) Signature(base io.Reader, blockSize uint64) (Signature, error) {
	// Choose a block size if none is specified. If the base also implements
	// io.Seeker (which most will since they need to for Patch), then use the
	// optimal block size, otherwise use the default.
	if blockSize == 0 {
		if baseSeeker, ok := base.(io.Seeker); ok {
			if s, err := OptimalBlockSizeForBase(baseSeeker); err == nil {
				blockSize = s
			} else {
				blockSize = DefaultBlockSize
			}
		} else {
			blockSize = DefaultBlockSize
		}
	}

	// Create the result.
	result := Signature{
		BlockSize: blockSize,
	}

	// Create a buffer with which to read blocks.
	buffer := e.bufferWithSize(blockSize)

	// Read blocks and append their hashes until we reach EOF.
	eof := false
	for !eof {
		// Read the next block and watch for errors. If we receive io.EOF, then
		// nothing was read, and we should break immediately. This means that
		// the base had a length that was a multiple of the block size. If we
		// receive io.ErrUnexpectedEOF, then something was read but we're still
		// at the end of the file, so we should hash this block but not go
		// through the loop again. All other errors are terminal.
		n, err := io.ReadFull(base, buffer)
		if err == io.EOF {
			result.LastBlockSize = blockSize
			break
		} else if err == io.ErrUnexpectedEOF {
			result.LastBlockSize = uint64(n)
			eof = true
		} else if err != nil {
			return Signature{}, errors.Wrap(err, "unable to read data block")
		}

		// Compute hashes for the the block that was read. For short blocks, we
		// still use the full block size when computing the weak hash. We could
		// alternatively use the short block length, but it doesn't matter - all
		// that matters is that we keep consistency when we compute the short
		// block weak hash when searching in Deltafy.
		weak, _, _ := e.weakHash(buffer[:n], blockSize)
		strong := e.strongHash(buffer[:n])

		// Add the block hash.
		result.Hashes = append(result.Hashes, BlockHash{weak, strong})
	}

	// If there are no hashes, then clear out the block sizes.
	if len(result.Hashes) == 0 {
		result.BlockSize = 0
		result.LastBlockSize = 0
	}

	// Success.
	return result, nil
}

// BytesSignature computes the signature for a byte slice.
func (e *Engine) BytesSignature(base []byte, blockSize uint64) Signature {
	// Perform the signature and watch for errors (which shouldn't be able to
	// occur in-memory).
	result, err := e.Signature(bytes.NewReader(base), blockSize)
	if err != nil {
		panic(errors.Wrap(err, "in-memory signature failure"))
	}

	// Success.
	return result
}

// dualModeReader unifies the io.Reader and io.ByteReader interfaces. It is used
// in deltafy operations to ensure that bytes can be efficiently extracted from
// targets.
type dualModeReader interface {
	io.Reader
	io.ByteReader
}

// min implements simple minimum finding for uint64 values.
func min(a, b uint64) uint64 {
	if a < b {
		return a
	}
	return b
}

// chunkAndTransmitAll is a fast-path routine for simply transmitting all data
// in a target stream. This is used when there are no blocks to match because
// the base stream is empty.
func (e *Engine) chunkAndTransmitAll(target io.Reader, maxDataOpSize uint64, transmit OperationTransmitter) error {
	// Verify that maxDataOpSize is sane.
	if maxDataOpSize == 0 {
		maxDataOpSize = DefaultMaximumDataOperationSize
	}

	// Create a buffer to transmit data operations.
	buffer := e.bufferWithSize(maxDataOpSize)

	// Loop until the entire target has been transmitted as data operations.
	for {
		if n, err := io.ReadFull(target, buffer); err == io.EOF {
			return nil
		} else if err == io.ErrUnexpectedEOF {
			if err = transmit(Operation{Data: buffer[:n]}); err != nil {
				return errors.Wrap(err, "unable to transmit data operation")
			}
			return nil
		} else if err != nil {
			return errors.Wrap(err, "unable to read target")
		} else if err = transmit(Operation{Data: buffer}); err != nil {
			return errors.Wrap(err, "unable to transmit data operation")
		}
	}
}

// Deltafy computes delta operations to reconstitute the target data stream
// using the base stream (based on the provided base signature). It streams
// operations to the provided transmission function. The internal engine buffer
// will be resized to the sum of the maximum data operation size plus the block
// size, and retained for the lifetime of the engine, so a reasonable value
// should be provided. The data buffer passed to the transmission function is
// reused, so the transmission function should transmit or make a copy of the
// data before returning.
func (e *Engine) Deltafy(target io.Reader, base Signature, maxDataOpSize uint64, transmit OperationTransmitter) error {
	// Verify that the signature is sane. We don't control its value, and if its
	// invariants are broken it can cause this method to behave strangely.
	if err := base.ensureValid(); err != nil {
		return errors.Wrap(err, "invalid signature")
	}

	// Verify that the maximum data operation size is sane.
	if maxDataOpSize == 0 {
		maxDataOpSize = DefaultMaximumDataOperationSize
	}

	// If the base is empty, then there's no way we'll find any matching blocks,
	// so just send the entire file.
	if len(base.Hashes) == 0 {
		return e.chunkAndTransmitAll(target, maxDataOpSize, transmit)
	}

	// Create a set of block and data transmitters that efficiently coalesce
	// adjacent block operations and provide data chunking. Some corresponding
	// finalization logic is required at the end of this function.
	var coalescedStart, coalescedCount uint64
	sendBlock := func(index uint64) error {
		if coalescedCount > 0 {
			if coalescedStart+coalescedCount == index {
				coalescedCount += 1
				return nil
			} else if err := transmit(Operation{Start: coalescedStart, Count: coalescedCount}); err != nil {
				return nil
			}
		}
		coalescedStart = index
		coalescedCount = 1
		return nil
	}
	sendData := func(data []byte) error {
		if len(data) > 0 && coalescedCount > 0 {
			if err := transmit(Operation{Start: coalescedStart, Count: coalescedCount}); err != nil {
				return err
			}
			coalescedStart = 0
			coalescedCount = 0
		}
		for len(data) > 0 {
			sendSize := min(uint64(len(data)), maxDataOpSize)
			if err := transmit(Operation{Data: data[:sendSize]}); err != nil {
				return err
			}
			data = data[sendSize:]
		}
		return nil
	}

	// Ensure that the target implements io.Reader and io.ByteReader. If it can
	// do this natively, great! If not, wrap it in our re-usable buffered
	// reader, but ensure that it is released when we're done so that we don't
	// retain it indefinitely.
	bufferedTarget, ok := target.(dualModeReader)
	if !ok {
		e.targetReader.Reset(target)
		bufferedTarget = e.targetReader
		defer func() {
			e.targetReader.Reset(nil)
		}()
	}

	// Create a lookup table that maps weak hashes to all matching block hashes.
	// If the last block is short, we extract it and hold it separately, because
	// when doing match searches, we assume that all blocks in this map have a
	// full block size worth of data.
	//
	// The rsync technical report (see the section titled "Checksum searching")
	// actually advocates a 3-tier search (i.e. an additional 16-bit hash layer
	// before the weak hash), but I think this probably isn't necessary with
	// modern hardware and hashing algorithms.
	//
	// TODO: This is currently a little expensive because it requires a slice
	// allocation for each map entry. I suspect that the collision rate for weak
	// hashes is actually sufficiently low that we could make each map value a
	// fixed array of int that would limit the number of matches we could try,
	// but save us a lot of allocating. We would have to use an int, because
	// invalid values would likely need to be -1. This might be an unnecessary
	// operation though, because this map is only generated for non-empty bases,
	// which typically don't come in large numbers. For a few files, generating
	// these maps with slice values is fine. It also might be a bit slow since
	// each insertion would require a linear search to find the insertion
	// location within the array.
	hashes := base.Hashes
	haveShortLastBlock := false
	var lastBlockIndex uint64
	var shortLastBlock BlockHash
	if base.LastBlockSize != base.BlockSize {
		haveShortLastBlock = true
		lastBlockIndex = uint64(len(hashes) - 1)
		shortLastBlock = hashes[lastBlockIndex]
		hashes = hashes[:lastBlockIndex]
	}
	weakToBlockHashes := make(map[uint32][]uint64, len(hashes))
	for i, h := range hashes {
		weakToBlockHashes[h.Weak] = append(weakToBlockHashes[h.Weak], uint64(i))
	}

	// Create a buffer that we can use to load data and search for matches. We
	// start by filling it with a block's worth of data and then continuously
	// appending bytes until we either fill the buffer (at which point we
	// transmit data preceding the block and truncate) or find a match (at
	// which point we transmit data preceding the block and then transmit the
	// block match). Once we're unable to append a new byte or refill with a
	// full block, we terminate our search and send the remaining data
	// (potentially searching for one last short block match at the end of the
	// buffer).
	//
	// We choose the buffer size to hold a chunk of data of the maximum allowed
	// transmission size and a block of data. This size choice is somewhat
	// arbitrary since we have a data chunking function and could load more data
	// before doing a truncation/transmission, but this is also a reasonable
	// amount of data to hold in memory at any given time. We could choose a
	// larger preceding data chunk size to have less frequent truncations, but
	// (a) truncations are cheap and (b) we'll probably be doing a lot of
	// sequential block matching cycles where we just continuously match blocks
	// at the beginning of the buffer and then refill, so truncations won't be
	// all that common.
	buffer := e.bufferWithSize(maxDataOpSize + base.BlockSize)

	// Track the occupancy of the buffer.
	var occupancy uint64

	// Track the weak hash and its parameters for the block at the end of the
	// buffer.
	var weak, r1, r2 uint32

	// Loop over the contents of the file and search for matches.
	for {
		// If the buffer is empty, then we need to read in a block's worth of
		// data (if possible) and calculate the weak hash and its parameters. If
		// the buffer is non-empty but less than a block's worth of data, then
		// we've broken an invariant in our code. Otherwise, we need to move the
		// search block one byte forward and roll the hash.
		if occupancy == 0 {
			if n, err := io.ReadFull(bufferedTarget, buffer[:base.BlockSize]); err == io.EOF || err == io.ErrUnexpectedEOF {
				occupancy = uint64(n)
				break
			} else if err != nil {
				return errors.Wrap(err, "unable to perform initial buffer fill")
			} else {
				occupancy = base.BlockSize
				weak, r1, r2 = e.weakHash(buffer[:occupancy], base.BlockSize)
			}
		} else if occupancy < base.BlockSize {
			panic("buffer contains less than a block worth of data")
		} else {
			if b, err := bufferedTarget.ReadByte(); err == io.EOF {
				break
			} else if err != nil {
				return errors.Wrap(err, "unable to read target byte")
			} else {
				weak, r1, r2 = e.rollWeakHash(r1, r2, buffer[occupancy-base.BlockSize], b, base.BlockSize)
				buffer[occupancy] = b
				occupancy += 1
			}
		}

		// Look for a block match for the block at the end of the buffer.
		potentials := weakToBlockHashes[weak]
		match := false
		var matchIndex uint64
		if len(potentials) > 0 {
			strong := e.strongHash(buffer[occupancy-base.BlockSize : occupancy])
			for _, p := range potentials {
				if base.Hashes[p].Strong == strong {
					match = true
					matchIndex = p
					break
				}
			}
		}

		// If there's a match, send any data preceding the match and then send
		// the match. Otherwise, if we've reached buffer capacity, send the data
		// preceding the search block.
		if match {
			if err := sendData(buffer[:occupancy-base.BlockSize]); err != nil {
				return errors.Wrap(err, "unable to transmit data preceding match")
			} else if err = sendBlock(matchIndex); err != nil {
				return errors.Wrap(err, "unable to transmit match")
			}
			occupancy = 0
		} else if occupancy == uint64(len(buffer)) {
			if err := sendData(buffer[:occupancy-base.BlockSize]); err != nil {
				return errors.Wrap(err, "unable to transmit data before truncation")
			}
			copy(buffer[:base.BlockSize], buffer[occupancy-base.BlockSize:occupancy])
			occupancy = base.BlockSize
		}
	}

	// If we have a short last block and the occupancy of the buffer is large
	// enough that it could match, then check for a match.
	if haveShortLastBlock && occupancy >= base.LastBlockSize {
		potentialLastBlockMatch := buffer[occupancy-base.LastBlockSize : occupancy]
		// For short blocks, we still use the full block size when computing the
		// weak hash. We could alternatively use the short block length, but it
		// doesn't matter - all that matters is that we keep consistency when we
		// compute the short block weak hash in Signature.
		if w, _, _ := e.weakHash(potentialLastBlockMatch, base.BlockSize); w == shortLastBlock.Weak {
			if e.strongHash(potentialLastBlockMatch) == shortLastBlock.Strong {
				if err := sendData(buffer[:occupancy-base.LastBlockSize]); err != nil {
					return errors.Wrap(err, "unable to transmit data")
				} else if err = sendBlock(lastBlockIndex); err != nil {
					return errors.Wrap(err, "unable to transmit operation")
				}
				occupancy = 0
			}
		}
	}

	// Send any data remaining in the buffer.
	if err := sendData(buffer[:occupancy]); err != nil {
		return errors.Wrap(err, "unable to send final data operation")
	}

	// Send any final pending coalesced operation. This can't be done as a defer
	// because we need to watch for errors.
	if coalescedCount > 0 {
		if err := transmit(Operation{Start: coalescedStart, Count: coalescedCount}); err != nil {
			return errors.Wrap(err, "unable to send final block operation")
		}
	}

	// Success.
	return nil
}

// DeltafyBytes computes delta operations for a byte slice. Unlike the streaming
// Deltafy method, it returns a slice of operations, which should be reasonable
// since the target data can already fit into memory.
func (e *Engine) DeltafyBytes(target []byte, base Signature, maxDataOpSize uint64) []Operation {
	// Create an empty result.
	var delta []Operation

	// Create an operation transmitter to populate the result. Note that we copy
	// any operation data buffers because they are re-used.
	transmit := func(operation Operation) error {
		// Copy the operation's data buffer if necessary.
		if len(operation.Data) > 0 {
			dataCopy := make([]byte, len(operation.Data))
			copy(dataCopy, operation.Data)
			operation.Data = dataCopy
		}

		// Record the operation.
		delta = append(delta, operation)

		// Success.
		return nil
	}

	// Wrap up the bytes in a reader.
	reader := bytes.NewReader(target)

	// Compute the delta and watch for errors (which shouldn't occur for for
	// in-memory data).
	if err := e.Deltafy(reader, base, maxDataOpSize, transmit); err != nil {
		panic(errors.Wrap(err, "in-memory deltafication failure"))
	}

	// Success.
	return delta
}

// Patch applies a single operation against a base stream to reconstitute the
// target into the destination stream.
func (e *Engine) Patch(destination io.Writer, base io.ReadSeeker, signature Signature, operation Operation) error {
	// Verify that the signature is sane. The caller probably does control its
	// value (i.e. it's most likely not coming from the network), but if its
	// invariants are broken it can cause this method to behave strangely.
	if err := signature.ensureValid(); err != nil {
		return errors.Wrap(err, "invalid signature")
	}

	// Verify that the operation is sane.
	if err := operation.ensureValid(); err != nil {
		return errors.Wrap(err, "invalid operation")
	}

	// Handle the operation based on type.
	if len(operation.Data) > 0 {
		// Write data operations directly to the destination.
		if _, err := destination.Write(operation.Data); err != nil {
			return errors.Wrap(err, "unable to write data")
		}
	} else {
		// Seek to the start of the requested block in base.
		// TODO: We should technically validate that operation.Index
		// multiplied by the block size can't overflow an int64. Worst case
		// at the moment it will cause the seek operation to fail.
		if _, err := base.Seek(int64(operation.Start)*int64(signature.BlockSize), io.SeekStart); err != nil {
			return errors.Wrap(err, "unable to seek to base location")
		}

		// Copy the requested number of blocks.
		for c := uint64(0); c < operation.Count; c++ {
			// Compute the size to copy.
			copyLength := signature.BlockSize
			if operation.Start+c == uint64(len(signature.Hashes)-1) {
				copyLength = signature.LastBlockSize
			}

			// Create a buffer of the required size.
			buffer := e.bufferWithSize(copyLength)

			// Copy the block.
			if _, err := io.ReadFull(base, buffer); err != nil {
				return errors.Wrap(err, "unable to read block data")
			} else if _, err = destination.Write(buffer); err != nil {
				return errors.Wrap(err, "unable to write block data")
			}
		}
	}

	// Success.
	return nil
}

// Patch applies a series of operations against a base byte slice to
// reconstitute the target byte slice.
func (e *Engine) PatchBytes(base []byte, signature Signature, delta []Operation) ([]byte, error) {
	// Wrap up the base bytes in a reader.
	baseReader := bytes.NewReader(base)

	// Create an output buffer.
	output := bytes.NewBuffer(nil)

	// Perform application.
	for _, o := range delta {
		if err := e.Patch(output, baseReader, signature, o); err != nil {
			return nil, err
		}
	}

	// Success.
	return output.Bytes(), nil
}