package object

// BuiltinFunction is the implementation of built-in function
type BuiltinFunction func(args ...Object) Object

// Builtin is the implementation of built-in object
type Builtin struct {
	Fn BuiltinFunction
}

// Type returns the type of object
func (b *Builtin) Type() Type {
	return BuiltinObj
}

// Inspect returns the string expression of object
func (b *Builtin) Inspect() string {
	return "builtin function"
}
