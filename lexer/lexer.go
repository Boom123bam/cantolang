package lexer

import (
	token "cantolang/token"
)

type Lexer struct {
	input    []rune
	pos      int
	char     rune
	peekChar rune
}

func New(input string) *Lexer {
	l := &Lexer{
		input: []rune(input),
	}
	l.pos = -1
	l.advance()
	return l
}

func (l *Lexer) advance() {
	l.pos++
	if l.pos < len(l.input) {
		l.char = l.input[l.pos]
	} else {
		l.char = 0
	}
	if l.pos+1 < len(l.input) {
		l.peekChar = l.input[l.pos+1]
	} else {
		l.peekChar = 0
	}
}

func (l *Lexer) readIdentifier() string {
	result := ""
	for isAllowedInIdent(l.char) {
		result += string(l.char)
		l.advance()
	}
	return result
}

func isAllowedInIdent(char rune) bool {
	if char == 0 {
		return false
	}
	restricted := []rune(" \n1234567890!@#$%^&" +
		token.OPEN_PAREN +
		token.CLOSE_PAREN +
		token.OPEN_BRACE +
		token.CLOSE_BRACE +
		token.FULLSTOP +
		token.COMMA +
		token.ADD +
		token.MINUS +
		token.MULTIPLY +
		token.DIVIDE)
	for _, c := range restricted {
		if c == char {
			return false
		}
	}
	return true
}

func (l *Lexer) readNumber() string {
	result := ""
	for l.char >= '0' && l.char <= '9' {
		result += string(l.char)
		l.advance()
	}
	return result
}

func (l *Lexer) ReadToken() token.Token {
	for l.char == ' ' || l.char == '\n' || l.char == '\r' || l.char == '\t' {
		l.advance()
	}
	if l.char == 0 {
		return token.Token{
			TokenType:    token.EOF,
			TokenLiteral: "",
		}
	}
	var t token.Token
	// single char tokens
	switch l.char {
	case '/':
		// check for comment
		if l.peekChar == '/' {
			for l.char != '\n' {
				l.advance()
			}
			t.TokenType = token.COMMENT
			return t
		}
		t.TokenType = token.DIVIDE
		t.TokenLiteral = string(l.char)
	case '（':
		t.TokenType = token.OPEN_PAREN
		t.TokenLiteral = string(l.char)
	case '）':
		t.TokenType = token.CLOSE_PAREN
		t.TokenLiteral = string(l.char)
	case '「':
		t.TokenType = token.OPEN_BRACE
		t.TokenLiteral = string(l.char)
	case '」':
		t.TokenType = token.CLOSE_BRACE
		t.TokenLiteral = string(l.char)
	case '。':
		t.TokenType = token.FULLSTOP
		t.TokenLiteral = string(l.char)
	case '，':
		t.TokenType = token.COMMA
		t.TokenLiteral = string(l.char)
	case '+':
		t.TokenType = token.ADD
		t.TokenLiteral = string(l.char)
	case '-':
		t.TokenType = token.MINUS
		t.TokenLiteral = string(l.char)
	case '*':
		t.TokenType = token.MULTIPLY
		t.TokenLiteral = string(l.char)
	case '塞':
		t.TokenType = token.ASSIGN
		t.TokenLiteral = string(l.char)
	case '入':
		t.TokenType = token.TO
		t.TokenLiteral = string(l.char)
	case '係':
		t.TokenType = token.EQUAL_TO
		t.TokenLiteral = string(l.char)
	case '就':
		t.TokenType = token.THEN
		t.TokenLiteral = string(l.char)
	case '當':
		t.TokenType = token.WHILE
		t.TokenLiteral = string(l.char)
	case '時':
		t.TokenType = token.SI
		t.TokenLiteral = string(l.char)
	default:
		if l.char >= '0' && l.char <= '9' {
			t.TokenType = token.NUMBER
			t.TokenLiteral = l.readNumber()
			return t
		}

		i := l.readIdentifier()
		if i == "" {
			t.TokenType = token.INVALID
			t.TokenLiteral = string(l.char)
		} else {
			t.TokenType = token.LookUpIdent(i)
			t.TokenLiteral = i
			return t
		}
	}
	l.advance()
	return t
}
