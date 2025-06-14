package ast

import (
	"bytes"
	"strings"

	"github.com/tshinag/monkey/token"
)

// ArrayLiteral implements array literal
type ArrayLiteral struct {
	Token    token.Token
	Elements []Expression
}

// TokenLiteral implements Node interface
func (al *ArrayLiteral) TokenLiteral() string {
	return al.Token.Literal
}

func (al *ArrayLiteral) String() string {
	var out bytes.Buffer
	elements := []string{}
	for _, el := range al.Elements {
		elements = append(elements, el.String())
	}
	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")
	return out.String()
}
