package ast

import (
	"bytes"
)

// Node is a node of AST tree
type Node interface {
	TokenLiteral() string
	String() string
}

// Statement is the expression of statement
type Statement interface {
	Node
	StatementNode()
}

// Expression is the expression of expression
type Expression interface {
	Node
	ExpressionNode()
}

// Program is a set of Statement
type Program struct {
	Statements []Statement
}

// TokenLiteral returns token literal for debug
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}
