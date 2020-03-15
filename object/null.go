package object

// Null is the implementation of null
type Null struct{}

// Type returns the type of object
func (n *Null) Type() Type {
	return NullType
}

// Inspect returns the string expression of object
func (n *Null) Inspect() string {
	return "null"
}
