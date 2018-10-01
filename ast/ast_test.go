package ast

import (
	"github.com/tamago0224/monkey/token"
	"testing"
)

func TestString(t *testing.T) {
	program := &Program{
		Statement: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "left"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "myVar"},
					Value: "myValue",
				},
				Value: &Identifier{
					Token: token.Token{Token: token.IDENT, Literal: "anotherVar"},
					Value: "anotherVar",
				},
			},
		},
	}

	if program.String() != "let myVar= anotherVar;" {
		t.Errorf("program.String() wrong. got=%q", program.String())
	}
}
