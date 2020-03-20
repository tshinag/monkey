package object

// Type expresses the type of object
type Type string

const (
	// NullType is the type of null
	NullType = "NULL"
	// IntegerType is the type of integer
	IntegerType = "INTEGER"
	// StringType is the type of string
	StringType = "STRING"
	// BooleanType is the type of boolean
	BooleanType = "BOOLEAN"
	// ReturnValueType is the type of return value
	ReturnValueType = "RETURN_VALUE"
	// ErrorType is the type of evalutation error
	ErrorType = "ERROR"
	// FunctionType is the type of function
	FunctionType = "FUNCTION"
	// ArrayType is the type of function
	ArrayType = "ARRAY"
	// HashType is the type of hash
	HashType = "HASH"
	// BuiltinType is the type of built-in object
	BuiltinType = "BUILTIN"
)

// Object is the expression of object
type Object interface {
	Type() Type
	Inspect() string
}
