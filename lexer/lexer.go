package lexer

import (
	token "cantolang/token"
)

type Lexer struct {
	input      []rune
	pos        int
	char       rune
	peekChar   rune
	quotePairs map[rune]rune
}

func New(input string) *Lexer {
	l := &Lexer{
		input: []rune(input),
	}
	l.pos = -1
	l.quotePairs = map[rune]rune{
		'"': '"',
		'“': '”',
		'”': 0,
	}
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

func (l *Lexer) readString(endChar rune) string {
	l.advance()
	result := ""
	for l.char != endChar && l.char != 0 {
		result += string(l.char)
		l.advance()
	}
	if l.char == endChar {
		l.advance()
	}
	return result
}

func isAllowedInIdent(char rune) bool {
	if char == 0 {
		return false
	}
	restricted := []rune(" \n1234567890!@#$%^&")
	for _, c := range restricted {
		if c == char {
			return false
		}
	}
	if token.LookUpSymbol(char) == token.TEMP_NOT_SYMBOL {
		return true
	}
	return false
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

	// check for comment
	if l.char == '/' && l.peekChar == '/' {
		for l.char != '\n' {
			l.advance()
		}
		t.TokenType = token.COMMENT
		return t
	}
	// check for symbol
	symbol := token.LookUpSymbol(l.char)
	if symbol != token.TEMP_NOT_SYMBOL {
		t.TokenType = symbol
		t.TokenLiteral = string(l.char)
		l.advance()
		return t
	}
	// check for number
	if l.char >= '0' && l.char <= '9' {
		t.TokenType = token.NUMBER
		t.TokenLiteral = l.readNumber()
		return t
	}
	// check for string
	matchingQuote, ok := l.quotePairs[l.char]
	if ok {
		if matchingQuote == 0 {
			t.TokenType = token.INVALID
			t.TokenLiteral = string(l.char)
			return t
		}
		t.TokenType = token.STRING
		t.TokenLiteral = l.readString(matchingQuote)
		return t
	}

	// identifier
	i := l.readIdentifier()
	if i == "" {
		t.TokenType = token.INVALID
		t.TokenLiteral = string(l.char)
	} else {
		t.TokenType = token.LookUpIdent(i)
		t.TokenLiteral = i
		return t
	}

	l.advance()
	return t
}
