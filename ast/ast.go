package ast

import "github.com/tshinag/monkey/token"

// Node is a node of AST tree
type Node interface {
	TokenLiteral() string
}

// Statement is the expression of statement
type Statement interface {
	Node
	statementNode()
}

// Expression is the expression of expression
type Expression interface {
	Node
	expressionNode()
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

// LetStatement implements let statement
type LetStatement struct {
	Token token.Token // token.LET トークン
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}

// TokenLiteral implements Node interface
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

// Identifier implements identifier
type Identifier struct {
	Token token.Token // token.IDENT トークン
	Value string
}

func (i *Identifier) expressionNode() {}

// TokenLiteral implements Node interface
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}
