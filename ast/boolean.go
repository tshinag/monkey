package ast

import "github.com/tshinag/monkey/token"

// Boolean implements boolean literal
type Boolean struct {
	Token token.Token
	Value bool
}

// TokenLiteral implements Node interface
func (b *Boolean) TokenLiteral() string {
	return b.Token.Literal
}

func (b *Boolean) String() string {
	return b.Token.Literal
}
