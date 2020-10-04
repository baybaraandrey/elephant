package lexer

import (
	"github.com/baybaraandrey/elephant/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `
	var number = 10
	var fu = def(a, b) { a + b }
	var sum = fu(1, 2)
	`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.VAR, "var"},
		{token.INDENT, "number"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.VAR, "var"},
		{token.INDENT, "fu"},
		{token.ASSIGN, "="},
		{token.DEF, "def"},
		{token.LPAR, "("},
		{token.INDENT, "a"},
		{token.COMMA, ","},
		{token.INDENT, "b"},
		{token.RPAR, ")"},
		{token.LBRACE, "{"},
		{token.INDENT, "a"},
		{token.PLUS, "+"},
		{token.INDENT, "b"},
		{token.RBRACE, "}"},
		{token.VAR, "var"},
		{token.INDENT, "sum"},
		{token.ASSIGN, "="},
		{token.INDENT, "fu"},
		{token.LPAR, "("},
		{token.INT, "1"},
		{token.COMMA, ","},
		{token.INT, "2"},
		{token.RPAR, ")"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q, ch=%q",
				i, tt.expectedType, tok.Type, tok.Literal)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
