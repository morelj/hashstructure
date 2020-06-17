package hashstructure

import "io"

// Hashable is an interface which can be implemented on any type which is able to compute its own hash.
// Some type cannot be hashed by this library (e.g. struct with no exported fields, like time.Time). In such case
// a custom Hashable implementation should be used.
type Hashable interface {
	// Hash computes the hash of the value it is called on, by sending data to the io.Writer.
	Hash(w io.Writer) error
}
