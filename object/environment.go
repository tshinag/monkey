package object

// Environment is the map of variables
type Environment struct {
	store map[string]Object
}

// NewEnvironment initializes and returns Environment
func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s}
}

// Get returns the value bound to variable
func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	return obj, ok
}

// Set binds the value to variable
func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}
