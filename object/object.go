package object

// Type expresses the type of object
type Type string

const (
	// NullObj is the type of null
	NullObj = "NULL"
	// IntegerObj is the type of integer
	IntegerObj = "INTEGER"
	// StringObj is the type of string
	StringObj = "STRING"
	// BooleanObj is the type of boolean
	BooleanObj = "BOOLEAN"
	// ReturnValueObj is the type of return value
	ReturnValueObj = "RETURN_VALUE"
	// ErrorObj is the type of evalutation error
	ErrorObj = "ERROR"
	// FunctionObj is the type of function
	FunctionObj = "FUNCTION"
)

// Object is the expression of object
type Object interface {
	Type() Type
	Inspect() string
}
