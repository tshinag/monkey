package object

// Error is the evalutation error
type Error struct {
	Message string
}

// Type returns the type of object
func (e *Error) Type() Type {
	return ErrorType
}

// Inspect returns the string expression of object
func (e *Error) Inspect() string {
	return "ERROR: " + e.Message
}
