package object

import "hash/fnv"

// String is the implementation of string
type String struct {
	Value string
}

// Type returns the type of object
func (s *String) Type() Type {
	return StringType
}

// Inspect returns the string expression of object
func (s *String) Inspect() string {
	return s.Value
}

// HashKey returns the hash key for hash map
func (s *String) HashKey() HashKey {
	h := fnv.New64a()
	h.Write([]byte(s.Value))
	return HashKey{Type: s.Type(), Value: h.Sum64()}
}
