package ast

import (
	"bytes"
	"strings"

	"github.com/tshinag/monkey/token"
)

// HashLiteral implements hash literal
type HashLiteral struct {
	Token token.Token
	Pairs map[Expression]Expression
}

// TokenLiteral implements Node interface
func (hl *HashLiteral) TokenLiteral() string {
	return hl.Token.Literal
}

func (hl *HashLiteral) String() string {
	var out bytes.Buffer
	pairs := []string{}
	for key, value := range hl.Pairs {
		pairs = append(pairs, key.String()+":"+value.String())
	}
	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")
	return out.String()
}
