package object

import "fmt"

// Integer is the implementation of integer
type Integer struct {
	Value int64
}

// Type returns the type of object
func (i *Integer) Type() Type {
	return IntegerType
}

// Inspect returns the string expression of object
func (i *Integer) Inspect() string {
	return fmt.Sprintf("%d", i.Value)
}

// HashKey returns the hash key for hash map
func (i *Integer) HashKey() HashKey {
	return HashKey{Type: i.Type(), Value: uint64(i.Value)}
}
