package lexer

import (
	"github.com/tshinag/monkey/token"
)

// Lexer is the experssion of lexer
type Lexer struct {
	input        string
	position     int  // 入力における現在の位置（現在の文字を指し示す）
	readPosition int  // これから読み込む位置（現在の文字の次）
	char         byte // 現在検査中の文字
}

// New initializes Lexer with input string
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// NextToken tokenize current charactor, then reads next
func (l *Lexer) NextToken() token.Token {
	l.skipWhitespace()
	switch l.char {
	case '=':
		if l.peekChar() == '=' {
			literal := l.readTwoChar()
			defer l.readChar()
			return token.New(token.EQ, literal)
		}
		defer l.readChar()
		return token.NewChar(token.ASSIGN, l.char)
	case '+':
		defer l.readChar()
		return token.NewChar(token.PLUS, l.char)
	case '-':
		defer l.readChar()
		return token.NewChar(token.MINUS, l.char)
	case '!':
		if l.peekChar() == '=' {
			literal := l.readTwoChar()
			defer l.readChar()
			return token.New(token.NOTEQ, literal)
		}
		defer l.readChar()
		return token.NewChar(token.BANG, l.char)
	case '/':
		defer l.readChar()
		return token.NewChar(token.SLASH, l.char)
	case '*':
		defer l.readChar()
		return token.NewChar(token.ASTERISK, l.char)
	case '<':
		defer l.readChar()
		return token.NewChar(token.LT, l.char)
	case '>':
		defer l.readChar()
		return token.NewChar(token.GT, l.char)
	case ';':
		defer l.readChar()
		return token.NewChar(token.SEMICOLON, l.char)
	case '(':
		defer l.readChar()
		return token.NewChar(token.LPAREN, l.char)
	case ')':
		defer l.readChar()
		return token.NewChar(token.RPAREN, l.char)
	case ',':
		defer l.readChar()
		return token.NewChar(token.COMMA, l.char)
	case '{':
		defer l.readChar()
		return token.NewChar(token.LBRACE, l.char)
	case '}':
		defer l.readChar()
		return token.NewChar(token.RBRACE, l.char)
	case '[':
		defer l.readChar()
		return token.NewChar(token.LBRACKET, l.char)
	case ']':
		defer l.readChar()
		return token.NewChar(token.RBRACKET, l.char)
	case 0:
		defer l.readChar()
		return token.New(token.EOF, "")
	case '"':
		str := l.readString()
		defer l.readChar()
		return token.New(token.STRING, str)
	default:
		if isLetter(l.char) {
			ident := l.readIdentifier()
			return token.NewIdent(ident)
		} else if isDigit(l.char) {
			num := l.readNumber()
			return token.New(token.INT, num)
		}
		defer l.readChar()
		return token.NewChar(token.ILLEGAL, l.char)
	}
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.char) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.char) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readString() string {
	position := l.position + 1
	l.readChar()
	for !isEndOfString(l.char) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readTwoChar() string {
	current := l.char
	l.readChar()
	next := l.char
	return string(current) + string(next)
}

func (l *Lexer) skipWhitespace() {
	for isWhitespace(l.char) {
		l.readChar()
	}
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.char = 0
	} else {
		l.char = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}

func isWhitespace(ch byte) bool {
	return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r'
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func isEndOfString(ch byte) bool {
	return ch == '"' || ch == 0
}
