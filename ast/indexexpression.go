package ast

import (
	"bytes"

	"github.com/tshinag/monkey/token"
)

// IndexExpression implements index expression
type IndexExpression struct {
	Token token.Token // '[' トークン
	Left  Expression
	Index Expression
}

// TokenLiteral implements Node interface
func (ie *IndexExpression) TokenLiteral() string {
	return ie.Token.Literal
}

func (ie *IndexExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString("[")
	out.WriteString(ie.Index.String())
	out.WriteString("])")
	return out.String()
}
