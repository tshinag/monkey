package object

// ReturnValue is the implementation of return value
type ReturnValue struct {
	Value Object
}

// Type returns the type of object
func (rv *ReturnValue) Type() Type {
	return ReturnValueType
}

// Inspect returns the string expression of object
func (rv *ReturnValue) Inspect() string {
	return rv.Value.Inspect()
}
