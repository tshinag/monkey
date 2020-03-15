package ast

import "github.com/tshinag/monkey/token"

// IntegerLiteral implements integer literal
type IntegerLiteral struct {
	Token token.Token
	Value int64
}

// TokenLiteral implements Node interface
func (il *IntegerLiteral) TokenLiteral() string {
	return il.Token.Literal
}

func (il *IntegerLiteral) String() string {
	return il.Token.Literal
}
