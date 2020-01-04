package parser

import (
	"fmt"

	"github.com/tshinag/monkey/ast"
	"github.com/tshinag/monkey/lexer"
	"github.com/tshinag/monkey/token"
)

// Parser is the implementation of parser
type Parser struct {
	l         *lexer.Lexer
	errors    []error
	curToken  token.Token
	peekToken token.Token
}

// New initializes Parser
func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []error{},
	}

	// 2つトークンを読み込む。curTokenとpeekTokenの両方がセットされる。
	p.nextToken()
	p.nextToken()

	return p
}

// ParseProgram outputs ast.Program
func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for !p.isCurToken(token.EOF) {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}

	return program
}

// Errors returns errors
func (p *Parser) Errors() []error {
	return p.errors
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return nil
	}
}

func (p *Parser) parseLetStatement() ast.Statement {
	stmt := &ast.LetStatement{Token: p.curToken}
	if !p.expectPeek(token.IDENT) {
		return nil
	}
	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	if !p.expectPeek(token.ASSIGN) {
		return nil
	}
	// TODO: セミコロンに遭遇するまで式を読み飛ばしてしまっている
	for !p.isCurToken(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.curToken}
	p.nextToken()
	// TODO: セミコロンに遭遇するまで式を読み飛ばしてしまっている
	for !p.isCurToken(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}

func (p *Parser) expectPeek(t token.Type) bool {
	if p.isPeekToken(t) {
		p.nextToken()
		return true
	}
	p.appendErrorPeek(t)
	return false
}

func (p *Parser) appendErrorPeek(t token.Type) {
	err := fmt.Errorf("expected next token to be %s, got %s instead",
		t, p.peekToken.Type)
	p.errors = append(p.errors, err)
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) isCurToken(t token.Type) bool {
	return p.curToken.IsType(t)
}

func (p *Parser) isPeekToken(t token.Type) bool {
	return p.peekToken.IsType(t)
}
