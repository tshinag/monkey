package object

import "fmt"

// Boolean is the implementation of boolean
type Boolean struct {
	Value bool
}

// Type returns the type of object
func (b *Boolean) Type() Type {
	return BooleanObj
}

// Inspect returns the string expression of object
func (b *Boolean) Inspect() string {
	return fmt.Sprintf("%t", b.Value)
}
