package object

import (
	"bytes"
	"strings"
)

// Array is the implementation of array
type Array struct {
	Elements []Object
}

// Type returns the type of object
func (a *Array) Type() Type {
	return ArrayObj
}

// Inspect returns the string expression of object
func (a *Array) Inspect() string {
	var out bytes.Buffer

	elements := []string{}
	for _, e := range a.Elements {
		elements = append(elements, e.Inspect())
	}

	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")

	return out.String()
}
