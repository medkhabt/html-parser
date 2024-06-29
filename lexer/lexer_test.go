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

func TestEndTagOpen(t *testing.T) {
	inputs := []string{`</`, `</>`}
	tests := [][]TokenTest{
		[]TokenTest{
			{token.CHARACTER, "<"},
			{token.CHARACTER, "/"},
			{token.EOF, "EOF"},
		},
		[]TokenTest{
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

func TestTagName(t *testing.T) {
	inputs := []string{`<a></a>`, `<A></A>`, `<h2></H2>`, `<H2></h2>`, `<quote></quote>`, `<QUOTE></QUOTE>`, `<QuoTE></qUOte>`}
	tests := [][]TokenTest{
		[]TokenTest{
			{token.STARTTAG, "a"},
			{token.ENDTAG, "a"},
		},
		[]TokenTest{
			{token.STARTTAG, "a"},
			{token.ENDTAG, "a"},
		},
		[]TokenTest{
			{token.STARTTAG, "h2"},
			{token.ENDTAG, "h2"},
		},
		[]TokenTest{
			{token.STARTTAG, "h2"},
			{token.ENDTAG, "h2"},
		},
		[]TokenTest{
			{token.STARTTAG, "quote"},
			{token.ENDTAG, "quote"},
		},
		[]TokenTest{
			{token.STARTTAG, "quote"},
			{token.ENDTAG, "quote"},
		},
		[]TokenTest{
			{token.STARTTAG, "quote"},
			{token.ENDTAG, "quote"},
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
