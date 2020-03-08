package object

// Environment is the map of variables
type Environment struct {
	store map[string]Object
	outer *Environment
}

// NewEnvironment initializes and returns Environment
func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s, outer: nil}
}

// NewEnclosedEnvironment initializes and returns Environment for closure
func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer
	return env
}

// Get returns the value bound to variable
func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	if ok || e.outer == nil {
		return obj, ok
	}
	return e.outer.Get(name)
}

// Set binds the value to variable
func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}
