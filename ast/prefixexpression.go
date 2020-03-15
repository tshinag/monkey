package ast

import (
	"bytes"

	"github.com/tshinag/monkey/token"
)

// PrefixExpression implements prefix operator
type PrefixExpression struct {
	Token    token.Token // 前置トークン、例えば「!」
	Operator string
	Right    Expression
}

// TokenLiteral implements Node interface
func (pe *PrefixExpression) TokenLiteral() string {
	return pe.Token.Literal
}

func (pe *PrefixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")

	return out.String()
}
