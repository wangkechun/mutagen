package session

import (
	"context"

	"github.com/havoc-io/mutagen/pkg/rsync"
	"github.com/havoc-io/mutagen/pkg/sync"
)

// endpoint provides an interface to synchronization endpoints. It provides all
// primitives necessary to support synchronization. None of its methods are safe
// for concurrent invocation except shutdown. If any method returns an error,
// the endpoint should be considered failed and no more of its methods (other
// than shutdown) should be invoked.
type endpoint interface {
	// poll performs a one-shot poll for filesystem modifications in the
	// endpoint's root. It blocks until an event occurs, the provided context is
	// cancelled, or an error occurs. In the first two cases it returns nil.
	poll(context context.Context) error

	// scan performs a scan of the endpoint's synchronization root. It requires
	// the ancestor to be passed in for optimized snapshot transfers if the
	// endpoint is remote. The ancestor may be nil, in which transfers from
	// remote endpoints may be less than optimal. It returns the scan result, a
	// bool indicating whether or not to re-try the scan, and any error that
	// occurred while trying to create the scan. Only one of these values will
	// be non-nil/false. If all are nil, it indicates that the synchronization
	// root doesn't exist on the endpoint, but that the scan otherwise completed
	// successfully.
	scan(ancestor *sync.Entry) (*sync.Entry, bool, error, bool)

	// stage performs staging on the endpoint. It accepts a list of file paths
	// and digests for those paths. It will filter the list based on what it
	// already has staged from previously interrupted stagings and what can be
	// staged from local contents (e.g. in cases of renames and copies), and
	// then return a list of paths, their signatures, and a receiver to receive
	// them. If the list of paths is empty, then all paths were either already
	// staged or able to be staged from local data, and the receiver will be
	// nil. Otherwise, the receiver will be non-nil and must be finalized (i.e.
	// transmitted to) before subsequent methods can be invoked on the endpoint.
	// If the receiver fails, the endpoint should be considered contaminated and
	// not used (though shutdown can and should still be invoked).
	stage(entries map[string][]byte) ([]string, []rsync.Signature, rsync.Receiver, error)

	// supply transmits files in a streaming fashion using the rsync algorithm
	// to the specified receiver.
	supply(paths []string, signatures []rsync.Signature, receiver rsync.Receiver) error

	// transition performs the specified transitions on the endpoint. It returns
	// a list of successfully applied changes and a list of problems that
	// occurred while applying transitions.
	// TODO: Should we consider pre-emptability for transition? It could
	// probably be done by just checking for cancellation during each transition
	// path and reporting "cancelled" for problems arising after that, but
	// usually the long-blocking transitions are going to be the ones where
	// we're creating the root with a huge number of files and wouldn't catch
	// cancellation until they're all done anyway.
	transition(transitions []*sync.Change) ([]*sync.Entry, []*sync.Problem, error)

	// shutdown terminates any resources associated with the endpoint. For local
	// endpoints, shutdown will not preempt calls, but for remote endpoints it
	// will because it closes the underlying connection to the endpoint
	// (actually, it terminates that connection). shutdown can safely be called
	// concurrently with other methods, though it's only recommended when you
	// don't want the possibility of preempting the method (e.g. in transition)
	// or you know that the operation can continue and terminate on its own
	// (e.g. in scan). shutdown should only be invoked once.
	shutdown() error
}
