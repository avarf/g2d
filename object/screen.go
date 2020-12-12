package object

import (
	"github.com/lucasepe/g2d/gg"
)

type Screen struct {
	Value *gg.Context
}

func (sc *Screen) String() string { return sc.Inspect() }

func (sc *Screen) Bool() bool { return sc.Value != nil }

// Type returns the type of the object
func (sc *Screen) Type() Type { return CANVAS }

// Inspect returns a stringified version of the object for debugging
func (sc *Screen) Inspect() string { return "<CANVAS>" }

// ToInterface converts this object to a go-interface, which will allow
// it to be used naturally in our sprintf/printf primitives.
//
// It might also be helpful for embedded users.
func (sc *Screen) ToInterface() interface{} { return "<CANVAS>" }

// Reset resets screen using the specified options.
func (sc *Screen) Reset(width, height int) {
	sc.Value = gg.NewContext(width, height)
}
