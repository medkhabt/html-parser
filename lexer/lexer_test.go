package lexer

import (
	"github/medkhabt/prs/token"
	"testing"
)

type TokenTest struct {
	expectedType    token.TokenType
	expectedLiteral string
}

func TestTagOpen(t *testing.T) {
	inputs := []string{"<", `<<`, `<#`, `<--`}
	tests := [][]TokenTest{
		[]TokenTest{
			{token.CHARACTER, "<"},
			{token.EOF, "EOF"},
		},
		[]TokenTest{
			{token.CHARACTER, "<"},
			{token.CHARACTER, "<"},
			{token.EOF, "EOF"},
		},
		[]TokenTest{
			{token.CHARACTER, "<"},
			{token.CHARACTER, "#"},
			{token.EOF, "EOF"},
		},
		[]TokenTest{
			{token.CHARACTER, "<"},
			{token.CHARACTER, "-"},
			{token.CHARACTER, "-"},
			{token.EOF, "EOF"},
		},
	}
	for j, inp := range inputs {
		l := New(inp)
		for i, tt := range tests[j] {
			tok := l.NextToken()
			if tok.Type != tt.expectedType {
				t.Fatalf("tests[%d] : token[%d] - tokentype wrong. expecte=%q, got=%q", j, i, tt.expectedType, tok.Type)
			}
			if tok.Literal != tt.expectedLiteral {
				t.Fatalf("tests[%d] : token[%d] - tokenLiteral wrong. expecte=%q, got=%q", j, i, tt.expectedLiteral, tok.Literal)
			}

		}
	}

}
