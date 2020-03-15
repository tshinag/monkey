package ast

import (
	"bytes"

	"github.com/tshinag/monkey/token"
)

// InfixExpression implements infix operator
type InfixExpression struct {
	Token    token.Token // 演算子トークン、例えば「+」
	Left     Expression
	Operator string
	Right    Expression
}

// TokenLiteral implements Node interface
func (oe *InfixExpression) TokenLiteral() string {
	return oe.Token.Literal
}

func (oe *InfixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(oe.Left.String())
	out.WriteString(" " + oe.Operator + " ")
	out.WriteString(oe.Right.String())
	out.WriteString(")")

	return out.String()
}
