package ast

import (
	"bytes"

	"github.com/tshinag/monkey/token"
)

// IfExpression implements if expression
type IfExpression struct {
	Token       token.Token // 'if' トークン
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

// TokenLiteral implements Node interface
func (ie *IfExpression) TokenLiteral() string {
	return ie.Token.Literal
}

func (ie *IfExpression) String() string {
	var out bytes.Buffer

	out.WriteString("if")
	out.WriteString(ie.Condition.String())
	out.WriteString(" ")
	out.WriteString(ie.Consequence.String())

	if ie.Alternative != nil {
		out.WriteString("else ")
		out.WriteString(ie.Alternative.String())
	}

	return out.String()
}
