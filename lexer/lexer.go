package lexer

import (
	token "cantolang/token"
)

type Lexer struct {
	input     []rune
	pos       int
	char      rune
	peek_char rune
}

func new(input string) Lexer {
	l := Lexer{
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
		l.peek_char = l.input[l.pos+1]
	} else {
		l.peek_char = 0
	}
}

func (l *Lexer) read_identifier() string {
	result := ""
	for is_allowed_in_ident(l.char) {
		result += string(l.char)
		l.advance()
	}
	return result
}

func is_allowed_in_ident(char rune) bool {
	restricted := []rune(" 1234567890!@#$%^&*/" +
		token.OPEN_PAREN +
		token.CLOSE_PAREN +
		token.OPEN_BRACE +
		token.CLOSE_BRACE +
		token.FULLSTOP +
		token.COMMA +
		token.ADD +
		token.MINUS +
		token.MULTIPLY +
		token.DIVIDE +
		token.INCREMENT +
		token.DECREMENT +
		token.INITIALIZE +
		token.ASSIGN +
		token.TO +
		token.EQUAL_TO +
		token.LESS_THAN +
		token.GREATER_THAN +
		token.AND +
		token.OR +
		token.NOT +
		token.IF +
		token.GEWA +
		token.THEN +
		token.WHILE +
		token.SI)
	for _, c := range restricted {
		if c == char {
			return false
		}
	}
	return true
}

func (l *Lexer) read_number() string {
	result := ""
	for l.char >= '0' && l.char <= '9' {
		result += string(l.char)
		l.advance()
	}
	return result
}

func (l *Lexer) read_token() token.Token {
	for l.char == ' ' || l.char == '\n' || l.char == '\r' || l.char == '\t' {
		l.advance()
	}
	if l.char == 0 {
		return token.Token{
			TokenType:    token.EOF,
			TokenLiteral: "",
		}
	}
	// double char tokens
	var t token.Token
	switch string(l.char) + string(l.peek_char) {
	case token.INITIALIZE:
		if string(l.char)+string(l.peek_char) == token.INITIALIZE {
			t.TokenType = token.INITIALIZE
			t.TokenLiteral = token.INITIALIZE
			l.advance()
		}
	case token.LESS_THAN:
		if string(l.char)+string(l.peek_char) == token.LESS_THAN {
			t.TokenType = token.LESS_THAN
			t.TokenLiteral = token.LESS_THAN
			l.advance()
		}
	case token.GREATER_THAN:
		if string(l.char)+string(l.peek_char) == token.GREATER_THAN {
			t.TokenType = token.GREATER_THAN
			t.TokenLiteral = token.GREATER_THAN
			l.advance()
		}
	case token.AND:
		if string(l.char)+string(l.peek_char) == token.AND {
			t.TokenType = token.AND
			t.TokenLiteral = token.AND
			l.advance()
		}
	case token.OR:
		if string(l.char)+string(l.peek_char) == token.OR {
			t.TokenType = token.OR
			t.TokenLiteral = token.OR
			l.advance()
		}
	case token.NOT:
		if string(l.char)+string(l.peek_char) == token.NOT {
			t.TokenType = token.NOT
			t.TokenLiteral = token.NOT
			l.advance()
		}
	case token.IF:
		if string(l.char)+string(l.peek_char) == token.IF {
			t.TokenType = token.IF
			t.TokenLiteral = token.IF
			l.advance()
		}
	case token.GEWA:
		if string(l.char)+string(l.peek_char) == token.GEWA {
			t.TokenType = token.GEWA
			t.TokenLiteral = token.GEWA
			l.advance()
		}
	case token.INCREMENT:
		if string(l.char)+string(l.peek_char) == token.INCREMENT {
			t.TokenType = token.INCREMENT
			t.TokenLiteral = token.INCREMENT
			l.advance()
		}
	case token.DECREMENT:
		if string(l.char)+string(l.peek_char) == token.DECREMENT {
			t.TokenType = token.DECREMENT
			t.TokenLiteral = token.DECREMENT
			l.advance()
		}
	case "//":
		for l.char != '\n' {
			l.advance()
		}
		t.TokenType = token.COMMENT
		return t
	}

	if t.TokenType != "" {
		l.advance()
		return t
	}

	// single char tokens
	switch l.char {
	case '/':
		t.TokenType = token.DIVIDE
		t.TokenLiteral = token.DIVIDE
	case '（':
		t.TokenType = token.OPEN_PAREN
		t.TokenLiteral = token.OPEN_PAREN
	case '）':
		t.TokenType = token.CLOSE_PAREN
		t.TokenLiteral = token.CLOSE_PAREN
	case '「':
		t.TokenType = token.OPEN_BRACE
		t.TokenLiteral = token.OPEN_BRACE
	case '」':
		t.TokenType = token.CLOSE_BRACE
		t.TokenLiteral = token.CLOSE_BRACE
	case '。':
		t.TokenType = token.FULLSTOP
		t.TokenLiteral = token.FULLSTOP
	case '，':
		t.TokenType = token.COMMA
		t.TokenLiteral = token.COMMA
	case '+':
		t.TokenType = token.ADD
		t.TokenLiteral = token.ADD
	case '-':
		t.TokenType = token.MINUS
		t.TokenLiteral = token.MINUS
	case '*':
		t.TokenType = token.MULTIPLY
		t.TokenLiteral = token.MULTIPLY
	case '塞':
		t.TokenType = token.ASSIGN
		t.TokenLiteral = token.ASSIGN
	case '入':
		t.TokenType = token.TO
		t.TokenLiteral = token.TO
	case '係':
		t.TokenType = token.EQUAL_TO
		t.TokenLiteral = token.EQUAL_TO
	case '就':
		t.TokenType = token.THEN
		t.TokenLiteral = token.THEN
	case '當':
		t.TokenType = token.WHILE
		t.TokenLiteral = token.WHILE
	case '時':
		t.TokenType = token.SI
		t.TokenLiteral = token.SI
	default:
		if l.char >= '0' && l.char <= '9' {
			t.TokenType = token.NUMBER
			t.TokenLiteral = l.read_number()
			return t
		}

		i := l.read_identifier()
		if i == "" {
			t.TokenType = token.INVALID
			t.TokenLiteral = string(l.char)
		} else {
			t.TokenType = token.IDENTIFIER
			t.TokenLiteral = i
			return t
		}
	}
	l.advance()
	return t
}
