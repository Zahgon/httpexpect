package httpexpect

import (
	"bytes"
	"context"
	"io"
	"sync"
)

// Wrapper for request or response body reader.
//
// Allows to read body multiple times using two approaches:
//   - use Read to read body contents and Rewind to restart reading from beginning
//   - use GetBody to get new reader for body contents
//
// When bodyWrapper is created, it does not read anything. Also, until anything is
// read, rewind operations are no-op.
//
// When the user starts reading body, bodyWrapper automatically copies retrieved
// content in memory. Then, when the body is fully read and Rewind is requested,
// it will close original body and switch to reading body from memory.
//
// If Rewind, GetBody, or Close is invoked before the body is fully read first time,
// bodyWrapper automatically performs full read.
//
// At any moment, the user can call DisableRewinds. In this case, Rewind and GetBody
// functionality is disabled, memory cache is cleared, and bodyWrapper switches to
// reading original body (if it's not fully read yet).
//
// bodyWrapper automatically creates finalizer that will close original body if the
// user never reads it fully or calls Closes.
type bodyWrapper struct {
	// Protects all operations.
	mu sync.Mutex

	// Original reader of HTTP response body.
	httpReader io.ReadCloser

	// Cancellation function for original HTTP response.
	// If set, called after HTTP response is fully read into memory.
	httpCancelFunc context.CancelFunc

	// Reader for HTTP response body stored in memory.
	// Rewind() resets this reader to start from the beginning.
	memReader *bytes.Reader

	// HTTP response body stored in memory.
	memBytes []byte

	// Cached read and close errors.
	readErr  error
	closeErr error

	// If true, Read will not store bytes in memory, and memBytes and memReader
	// won't be used.
	isRewindDisabled bool

	// True means that HTTP response was fully read into memory already.
	isFullyRead bool

	// True means that a read operation of any type was called at least once.
	isReadBefore bool
}

func newBodyWrapper(reader io.ReadCloser, cancelFunc context.CancelFunc) *bodyWrapper {
	_ = "STUB: not implemented"
	return nil
}

// Finalizer will close body if closeAndCancel was never called.

// Read body contents.
func (bw *bodyWrapper) Read(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// Regular read from original HTTP response.

// Read from original HTTP response + store into memory.

// Read from memory.

// Close body.
func (bw *bodyWrapper) Close() error { _ = "STUB: not implemented"; return nil }

// Preserve original reader error.

// Rewind or GetBody may be called later, so be sure to
// read body into memory before closing.

// Close original reader.

// Reset memory reader.

// Free memory when rewind is disabled.

// Rewind reading to the beginning.
func (bw *bodyWrapper) Rewind() { _ = "STUB: not implemented"; return }

// Rewind is no-op if disabled.

// Rewind is no-op until first read operation.

// If HTTP response is not fully read yet, do it now.
// If error occurs, it will be reported next read operation.

// Reset memory reader.

// Create new reader to retrieve body contents.
// New reader always reads body from the beginning.
// Does not affected by Rewind().
func (bw *bodyWrapper) GetBody() (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

// Preserve original reader error.

// GetBody() requires rewinds to be enabled.

// If HTTP response is not fully read yet, do it now.

// Return fresh reader for memory chunk.

// Disables storing body contents in memory and clears the cache.
func (bw *bodyWrapper) DisableRewinds() { _ = "STUB: not implemented"; return }

// Free memory if reading from original HTTP response, or reading from memory
// and memory reader has nothing left to read.
// Otherwise, i.e. when we're reading from memory, and there is more to read,
// memReadNext() will free memory later when it hits EOF.

func (bw *bodyWrapper) memReadNext(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// Free memory after we hit EOF when reading from memory,
// if rewinds were disabled while we were reading from it.

func (bw *bodyWrapper) httpReadNext(p []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Switch to reading from memory.

func (bw *bodyWrapper) httpReadFull() error { _ = "STUB: not implemented"; return nil }

// Switch to reading from memory.

func (bw *bodyWrapper) closeAndCancel() error { _ = "STUB: not implemented"; return nil }

// Finalizer is not needed anymore.
