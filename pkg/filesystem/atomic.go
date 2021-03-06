package filesystem

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
)

const (
	// atomicWriteTemporaryNamePrefix is the file name prefix to use for
	// intermediate temporary files used in atomic writes.
	atomicWriteTemporaryNamePrefix = ".mutagen-atomic-write"
	// atomicRenameTemporaryNamePrefix is the file name prefix to use for
	// intermediate temporary files used in atomic rename emulation.
	atomicRenameTemporaryNamePrefix = ".mutagen-atomic-rename"
)

// IsAtomicOperationFileName determines whether or not a file name (not a file
// path) is the name of an intermediate temporary file used in an atomic
// operation.
func IsAtomicOperationFileName(name string) bool {
	return strings.HasPrefix(name, atomicWriteTemporaryNamePrefix) ||
		strings.HasPrefix(name, atomicRenameTemporaryNamePrefix)
}

// WriteFileAtomic writes a file to disk in an atomic fashion by using an
// intermediate temporary file that is swapped in place using a rename
// operation.
func WriteFileAtomic(path string, data []byte, permissions os.FileMode) error {
	// Create a temporary file. The ioutil module already uses secure
	// permissions for creating the temporary file, so we don't need to specify
	// any.
	temporary, err := ioutil.TempFile(filepath.Dir(path), atomicWriteTemporaryNamePrefix)
	if err != nil {
		return errors.Wrap(err, "unable to create temporary file")
	}

	// Write data.
	if _, err = temporary.Write(data); err != nil {
		temporary.Close()
		os.Remove(temporary.Name())
		return errors.Wrap(err, "unable to write data to temporary file")
	}

	// Close out the file.
	if err = temporary.Close(); err != nil {
		os.Remove(temporary.Name())
		return errors.Wrap(err, "unable to close temporary file")
	}

	// Set the file's permissions.
	if err = os.Chmod(temporary.Name(), permissions); err != nil {
		os.Remove(temporary.Name())
		return errors.Wrap(err, "unable to change file permissions")
	}

	// Rename the file.
	if err = os.Rename(temporary.Name(), path); err != nil {
		os.Remove(temporary.Name())
		return errors.Wrap(err, "unable to rename file")
	}

	// Success.
	return nil
}

// RenameFileAtomic performs an atomic file rename. In the simplest case, it's a
// simple alias for os.Rename. However, if moving a file across filesystems, it
// will fall back to a copy/rename combination that should still approximate
// atomicity (at least in terms of swapping the destination file contents). It
// does NOT support renaming directories, only files. It takes inspiration from:
// https://github.com/golang/dep/blob/4ad9f4ec24012607dc247ca24528e3224d61519a/fs.go#L80
func RenameFileAtomic(oldPath, newPath string) error {
	// Try to peform an atomic rename. If we succeed, or encounter an error that
	// isn't a cross-device error, then we're done.
	if err := os.Rename(oldPath, newPath); err == nil {
		return nil
	} else if !isCrossDeviceError(err) {
		return err
	}

	// Open the source file.
	source, err := os.Open(oldPath)
	if err != nil {
		return errors.Wrap(err, "unable to open source file")
	}

	// Grab the source file's mode and ownership information.
	metadata, err := source.Stat()
	if err != nil {
		source.Close()
		return errors.Wrap(err, "unable to grab source file metadata")
	}
	mode := metadata.Mode()
	uid, gid, err := GetOwnership(metadata)
	if err != nil {
		source.Close()
		return errors.Wrap(err, "unable to grab source file ownership")
	}

	// Create a temporary file next to the destination.
	temporary, err := ioutil.TempFile(filepath.Dir(newPath), atomicRenameTemporaryNamePrefix)
	if err != nil {
		source.Close()
		return errors.Wrap(err, "unable to create temporary file")
	}
	temporaryPath := temporary.Name()

	// Copy the file contents. We'll handle errors below.
	_, err = io.Copy(temporary, source)

	// Close out files.
	source.Close()
	temporary.Close()

	// If there was a copy error, then remove the temporary and abort.
	if err != nil {
		os.Remove(temporaryPath)
		return errors.Wrap(err, "unable to copy file contents")
	}

	// Set the file mode on the destination. Note that we can't do this using
	// os.File.Chmod because that's not supported on Windows - only path-based
	// Chmod is supported.
	if err := os.Chmod(temporaryPath, mode); err != nil {
		os.Remove(temporaryPath)
		return errors.Wrap(err, "unable to set file mode")
	}

	// Set the file ownership on the destination.
	if err := SetOwnership(temporaryPath, uid, gid); err != nil {
		os.Remove(temporaryPath)
		return errors.Wrap(err, "unable to set file ownership")
	}

	// Move the temporary file into place.
	if err := os.Rename(temporaryPath, newPath); err != nil {
		os.Remove(temporaryPath)
		return errors.Wrap(err, "unable to rename temporary file")
	}

	// The file is in place, so remove the source file. We don't check for
	// errors on this removal since there's not much point in trying to do
	// anything about them.
	os.Remove(oldPath)

	// Success.
	return nil
}
