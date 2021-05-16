package token

import (
	"strconv"
)

type TokenType int

// List of token types
const (
	ILLEGAL TokenType = iota
	EOF               // End of file
	COMMENT           // Comment or comments

	literals_b
	// literals and identifiers
	INDENT // A
	INT    // 45
	FLOAT  // 3.1415
	STRING // "string"
	literals_e

	operators_b
	// operators
	PERCENT // %
	AMPER   // &
	LPAR    // (
	RPAR    // )
	STAR    // *
	PLUS    // +
	COMMA   // ,
	MINUS   // -
	DOT     // .
	SLASH   // /
	COLON   // :
	SEMI    // ;
	LESS    // <
	ASSIGN  // =
	EQUAL   // ==
	DEFINE  // :=

	GREATER    // >
	AT         // @
	LSQB       // [
	RSQB       // ]
	CIRCUMFLEX // ^
	LBRACE     // {
	VBAR       // |
	RBRACE     // }
	TILDE      // ~
	operators_e

	keywords_b
	// Keywords
	DEF    // def
	VAR    // var
	IF     // if
	ELSE   // else
	RETURN // return
	TRUE   // true
	FALSE  // false

	keywords_e
)

var tokens = [...]string{
	ILLEGAL: "ILLEGAL",

	EOF:     "EOF",
	COMMENT: "COMMENT",

	INDENT: "INDENT",
	INT:    "INT",
	FLOAT:  "FLOAT",
	STRING: "STRING",

	PERCENT: "%",
	AMPER:   "&",
	LPAR:    "(",
	RPAR:    ")",
	STAR:    "*",
	PLUS:    "+",
	COMMA:   ",",
	MINUS:   "-",
	DOT:     ".",
	SLASH:   "/",
	COLON:   ":",
	SEMI:    ";",
	LESS:    "<",
	ASSIGN:  "=",
	EQUAL:   "==",
	DEFINE:  ":=",

	GREATER:    ">",
	AT:         "@",
	LSQB:       "[",
	RSQB:       "]",
	CIRCUMFLEX: "^",
	LBRACE:     "{",
	VBAR:       "|",
	RBRACE:     "}",
	TILDE:      "~",

	DEF:    "def",
	VAR:    "var",
	IF:     "if",
	ELSE:   "else",
	RETURN: "return",
	TRUE:   "true",
	FALSE:  "false",
}

func (tok TokenType) String() string {
	s := ""
	if 0 <= tok && tok < TokenType(len(tokens)) {
		s = tokens[tok]
	}
	if s == "" {
		s = "token(" + strconv.Itoa(int(tok)) + ")"
	}
	return s
}

type Token struct {
	Type    TokenType
	Literal string
}
