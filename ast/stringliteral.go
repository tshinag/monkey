package ast

import "github.com/tshinag/monkey/token"

// StringLiteral implements string literal
type StringLiteral struct {
	Token token.Token
	Value string
}

// TokenLiteral implements Node interface
func (il *StringLiteral) TokenLiteral() string {
	return il.Token.Literal
}

func (il *StringLiteral) String() string {
	return il.Token.Literal
}
