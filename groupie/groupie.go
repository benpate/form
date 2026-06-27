// Package groupie tracks group boundaries while iterating a sorted list,
// reporting when a value differs from the one seen on the previous call.
package groupie

// Groupie remembers the last value it was shown so that callers can detect
// when a new group begins in an already-sorted sequence.
type Groupie struct {
	lastValue any
}

// New returns an empty Groupie, ready to track group boundaries.
func New() *Groupie {
	return &Groupie{}
}

// Header returns TRUE when value differs from the value seen on the previous
// call, signaling that a new group header should be drawn.
func (g *Groupie) Header(value any) bool {

	if g.lastValue == value {
		return false
	}

	g.lastValue = value
	return true
}
