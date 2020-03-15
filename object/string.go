package object

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
