package tool

import "fmt"

// ErrorKind is the fixed, closed set of failure categories a tool classifies
// any API error into. This is the anti-corruption boundary: the messy variety
// of real-world errors collapses into these kinds, and nothing past this
// package has to know what the underlying API client returned.
//
// Unknown is the zero value, so an unclassified or unset kind is never
// mistaken for a specific, confident classification.
type ErrorKind int

const (
	Unknown      ErrorKind = iota // 0 — couldn't classify; safe default
	Timeout                       // the call exceeded its time budget
	RateLimited                   // the API throttled us (e.g. HTTP 429)
	NotFound                      // ticker/endpoint doesn't exist (e.g. HTTP 404)
	Unauthorized                  // bad/missing API key (e.g. HTTP 401/403)
	Malformed                     // response wasn't the shape we expected
	Unavailable                   // service is down briefly (e.g. HTTP 503)
	Network                       // connection-level failure (DNS, reset, refused)
)

// String makes ErrorKind print as its name in error strings and logs instead
// of a bare integer.
func (k ErrorKind) String() string {
	switch k {
	case Unknown:
		return "Unknown"
	case Timeout:
		return "Timeout"
	case RateLimited:
		return "RateLimited"
	case NotFound:
		return "NotFound"
	case Unauthorized:
		return "Unauthorized"
	case Malformed:
		return "Malformed"
	case Unavailable:
		return "Unavailable"
	case Network:
		return "Network"
	default:
		return fmt.Sprintf("ErrorKind(%d)", int(k))
	}
}

// Retryable reports whether this kind of failure is EVER worth retrying. It
// answers only the universal question; the worker owns the how-much (backoff,
// attempt count, remaining time budget).
//
// The default branch makes Unknown — and any future kind we add — non-retryable
// until deliberately classified, so a new failure mode never causes accidental
// retry-hammering on a broken service.
func (k ErrorKind) Retryable() bool {
	switch k {
	case Timeout, RateLimited, Network, Unavailable:
		return true
	default:
		return false
	}
}

// ToolError is the single error type tools return. It carries the classified
// Kind, a human-readable Message, and the original Wrapped error (if any).
type ToolError struct {
	Kind    ErrorKind
	Message string
	Wrapped error // the original error from the API client, if any
}

// Error satisfies the built-in error interface. The pointer receiver means
// *ToolError is the error value (a nil *ToolError is not a usable error).
func (e *ToolError) Error() string {
	if e.Wrapped != nil {
		return fmt.Sprintf("%s: %s: %v", e.Kind, e.Message, e.Wrapped)
	}
	return fmt.Sprintf("%s: %s", e.Kind, e.Message)
}

// Unwrap exposes the wrapped error so errors.Is and errors.As can walk through
// a *ToolError to inspect the original cause.
func (e *ToolError) Unwrap() error { return e.Wrapped }
