package object

import (
	"bytes"
	"fmt"
	"strings"
)

// Hash is the implementation of hash
type Hash struct {
	Pairs map[HashKey]HashPair
}

// Type returns the type of object
func (h *Hash) Type() Type {
	return HashType
}

// Inspect returns the string expression of object
func (h *Hash) Inspect() string {
	var out bytes.Buffer
	pairs := []string{}
	for _, pair := range h.Pairs {
		pairs = append(pairs, fmt.Sprintf("%s: %s",
			pair.Key.Inspect(), pair.Value.Inspect()))
	}
	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")
	return out.String()
}
