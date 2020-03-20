package object

import "fmt"

// Boolean is the implementation of boolean
type Boolean struct {
	Value bool
}

// Type returns the type of object
func (b *Boolean) Type() Type {
	return BooleanType
}

// Inspect returns the string expression of object
func (b *Boolean) Inspect() string {
	return fmt.Sprintf("%t", b.Value)
}

// HashKey returns the hash key for hash map
func (b *Boolean) HashKey() HashKey {
	var value uint64
	if b.Value {
		value = 1
	} else {
		value = 0
	}
	return HashKey{Type: b.Type(), Value: value}
}
