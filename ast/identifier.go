package ast

import "github.com/tshinag/monkey/token"

// Identifier implements identifier
type Identifier struct {
	Token token.Token // token.IDENT トークン
	Value string
}

// TokenLiteral implements Node interface
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

func (i *Identifier) String() string {
	return i.Value
}
