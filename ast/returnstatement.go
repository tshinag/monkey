package ast

import (
	"bytes"

	"github.com/tshinag/monkey/token"
)

// ReturnStatement implements return statement
type ReturnStatement struct {
	Token       token.Token // 'return' トークン
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode() {}

// TokenLiteral implements Node interface
func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}

func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")

	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}

	out.WriteString(";")
	return out.String()
}
