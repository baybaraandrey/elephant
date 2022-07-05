package ast

import (
	"testing"

	"github.com/baybaraandrey/elephant/token"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "myVar"},
					Value: "myVar",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "anotherVar"},
					Value: "anotherVar",
				},
			},
		},
	}

	if program.String() != "let myVar = anotherVar;" {
		t.Errorf("program.String() wrong. got=%q", program.String())
	}
}

func TestAssignString(t *testing.T) {
	exp := &AssignStatement{
		Token: token.Token{Type: token.IDENT, Literal: "myX"},
		Name: &Identifier{
			Token: token.Token{Type: token.IDENT, Literal: "myX"},
			Value: "myVar",
		},
		Value: &Identifier{
			Token: token.Token{Type: token.IDENT, Literal: "anotherVar"},
			Value: "anotherVar",
		},
	}

	if exp.String() != "myVar = anotherVar;" {
		t.Errorf("exp.String() wrong. got=%q", exp.String())
	}
}
