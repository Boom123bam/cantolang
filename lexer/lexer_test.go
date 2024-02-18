package lexer

import (
	"cantolang/token"
	"testing"
)

func TestLexer(t *testing.T) {
	input := `
	// initialize
    叫佢 i。

    當 （i 細過 8） 時，就「
        講（i）。
        塞 i + 1 入 i。
    」

    a 大D。
`

	expectedTokens := []struct {
		Type    string
		Literal string
	}{
		{token.COMMENT, ""},
		{token.INITIALIZE, "叫佢"},
		{token.IDENTIFIER, "i"},
		{token.FULLSTOP, "。"},
		{token.WHILE, "當"},
		{token.OPEN_PAREN, "（"},
		{token.IDENTIFIER, "i"},
		{token.LESS_THAN, "細過"},
		{token.NUMBER, "8"},
		{token.CLOSE_PAREN, "）"},
		{token.SI, "時"},
		{token.COMMA, "，"},
		{token.THEN, "就"},
		{token.OPEN_BRACE, "「"},
		{token.IDENTIFIER, "講"},
		{token.OPEN_PAREN, "（"},
		{token.IDENTIFIER, "i"},
		{token.CLOSE_PAREN, "）"},
		{token.FULLSTOP, "。"},
		{token.ASSIGN, "塞"},
		{token.IDENTIFIER, "i"},
		{token.ADD, "+"},
		{token.NUMBER, "1"},
		{token.TO, "入"},
		{token.IDENTIFIER, "i"},
		{token.FULLSTOP, "。"},
		{token.CLOSE_BRACE, "」"},
		{token.IDENTIFIER, "a"},
		{token.INCREMENT, "大D"},
		{token.FULLSTOP, "。"},
		{token.EOF, ""},
	}
	l := New(input)
	for i, exp := range expectedTokens {
		got := l.ReadToken()
		if got.TokenLiteral != exp.Literal {
			t.Errorf("tests[%d] Expected literal '%s' got '%s'", i, exp.Literal, got.TokenLiteral)
		}
		if got.TokenType != exp.Type {
			t.Errorf("tests[%d] Expected type '%s' got '%s'", i, exp.Type, got.TokenType)
		}

	}
}
