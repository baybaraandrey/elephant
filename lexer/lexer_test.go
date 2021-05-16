package lexer

import (
	"testing"

	"github.com/baybaraandrey/elephant/token"
)

func TestNextToken(t *testing.T) {
	input := `
	var number = 10
	var fu = def(a, b) { a + b }
	var sum = fu(1, 2)
	var add = def(x, y) {
		x + y
	}
	def nameOfFunction() {
		if 5 < 10 {
			return true
		} else {
			return false
		}
	}
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
		{token.VAR, "var"},
		{token.INDENT, "add"},
		{token.ASSIGN, "="},
		{token.DEF, "def"},
		{token.LPAR, "("},
		{token.INDENT, "x"},
		{token.COMMA, ","},
		{token.INDENT, "y"},
		{token.RPAR, ")"},
		{token.LBRACE, "{"},
		{token.INDENT, "x"},
		{token.PLUS, "+"},
		{token.INDENT, "y"},
		{token.RBRACE, "}"},
		{token.DEF, "def"},
		{token.INDENT, "nameOfFunction"},
		{token.LPAR, "("},
		{token.RPAR, ")"},
		{token.LBRACE, "{"},
		{token.IF, "if"},
		{token.INT, "5"},
		{token.LESS, "<"},
		{token.INT, "10"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.RBRACE, "}"},
		{token.RBRACE, "}"},
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
