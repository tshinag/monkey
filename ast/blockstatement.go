package ast

import (
	"bytes"

	"github.com/tshinag/monkey/token"
)

// BlockStatement implements block statement
type BlockStatement struct {
	Token      token.Token // { トークン
	Statements []Statement
}

// TokenLiteral implements Node interface
func (bs *BlockStatement) TokenLiteral() string {
	return bs.Token.Literal
}

func (bs *BlockStatement) String() string {
	var out bytes.Buffer

	for _, s := range bs.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}
