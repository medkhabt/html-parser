package lexer

import (
	"github/medkhabt/prs/comparator"
	"github/medkhabt/prs/token"
	"testing"
)

type TokenTest struct {
	expectedType        token.TokenType
	expectedLiteral     []byte
	expectedForceQuirks bool
}

func TestGeneral(t *testing.T) {
	input := []byte("<test>ra1</test>")
	l := New(input)
	t.Logf("%v %v %v %v %v", l.NextToken(), l.NextToken(), l.NextToken(), l.NextToken(), l.NextToken())
}
func TestTagOpen(t *testing.T) {
	inputs := [][]byte{[]byte("<"), []byte(`<<`), []byte(`<#`), []byte(`<--`)}
	tests := [][]TokenTest{
		[]TokenTest{
			{token.CHARACTER, []byte{'<'}, false},
			{token.EOF, []byte{'E', 'O', 'F'}, false},
		},
		[]TokenTest{
			{token.CHARACTER, []byte{'<'}, false},
			{token.CHARACTER, []byte{'<'}, false},
			{token.EOF, []byte{'E', 'O', 'F'}, false},
		},
		[]TokenTest{
			{token.CHARACTER, []byte{'<'}, false},
			{token.CHARACTER, []byte{'#'}, false},
			{token.EOF, []byte{'E', 'O', 'F'}, false},
		},
		[]TokenTest{
			{token.CHARACTER, []byte{'<'}, false},
			{token.CHARACTER, []byte{'-'}, false},
			{token.CHARACTER, []byte{'-'}, false},
			{token.EOF, []byte{'E', 'O', 'F'}, false},
		},
	}
	nextTokenTestFormat(t, inputs, tests)
}

func TestEndTagOpen(t *testing.T) {
	inputs := [][]byte{[]byte(`</`), []byte(`</>`)}
	tests := [][]TokenTest{
		[]TokenTest{
			{token.CHARACTER, []byte("<"), false},
			{token.CHARACTER, []byte("/"), false},
			{token.EOF, []byte("EOF"), false},
		},
		[]TokenTest{
			{token.EOF, []byte("EOF"), false},
		},
	}
	nextTokenTestFormat(t, inputs, tests)
}

func TestTagName(t *testing.T) {
	inputs := [][]byte{[]byte(`<a></a>`), []byte(`<A></A>`), []byte(`<h2></H2>`), []byte(`<H2></h2>`), []byte(`<quote></quote>`), []byte(`<QUOTE></QUOTE>`), []byte(`<QuoTE></qUOte>`)}
	tests := [][]TokenTest{
		[]TokenTest{
			{token.STARTTAG, []byte("a"), false},
			{token.ENDTAG, []byte("a"), false},
			{token.EOF, []byte{'E', 'O', 'F'}, false},
		},
		[]TokenTest{
			{token.STARTTAG, []byte("a"), false},
			{token.ENDTAG, []byte("a"), false},
			{token.EOF, []byte{'E', 'O', 'F'}, false},
		},
		[]TokenTest{
			{token.STARTTAG, []byte("h2"), false},
			{token.ENDTAG, []byte("h2"), false},
			{token.EOF, []byte{'E', 'O', 'F'}, false},
		},
		[]TokenTest{
			{token.STARTTAG, []byte("h2"), false},
			{token.ENDTAG, []byte("h2"), false},
			{token.EOF, []byte{'E', 'O', 'F'}, false},
		},
		[]TokenTest{
			{token.STARTTAG, []byte("quote"), false},
			{token.ENDTAG, []byte("quote"), false},
			{token.EOF, []byte{'E', 'O', 'F'}, false},
		},
		[]TokenTest{
			{token.STARTTAG, []byte("quote"), false},
			{token.ENDTAG, []byte("quote"), false},
			{token.EOF, []byte{'E', 'O', 'F'}, false},
		},
		[]TokenTest{
			{token.STARTTAG, []byte("quote"), false},
			{token.ENDTAG, []byte("quote"), false},
			{token.EOF, []byte{'E', 'O', 'F'}, false},
		},
	}
	nextTokenTestFormat(t, inputs, tests)
}

func TestBogusCommentState(t *testing.T) {
	inputs := [][]byte{
		[]byte("<?random><?random"),
		[]byte("<!DOCTYPR><!doctyp002"),
		[]byte("</723#?></#div"),
	}
	tests := [][]TokenTest{
		[]TokenTest{
			{token.COMMENT, []byte("?random"), false},
			{token.COMMENT, []byte("?random"), false},
			{token.EOF, []byte("EOF"), false},
		},
		[]TokenTest{
			{token.COMMENT, []byte("DOCTYPR"), false},
			{token.COMMENT, []byte("doctyp002"), false},
			{token.EOF, []byte("EOF"), false},
		},
		[]TokenTest{
			{token.COMMENT, []byte("723#?"), false},
			{token.COMMENT, []byte("#div"), false},
			{token.EOF, []byte("EOF"), false},
		},
	}
	nextTokenTestFormat(t, inputs, tests)
}

func TestCommentStart(t *testing.T) {
	inputs := [][]byte{
		[]byte("<!--><!--"),
	}
	tests := [][]TokenTest{
		[]TokenTest{
			{token.COMMENT, []byte(""), false},
			{token.COMMENT, []byte(""), false},
			{token.EOF, []byte("EOF"), false},
		},
	}
	nextTokenTestFormat(t, inputs, tests)
}

func TestDoctype(t *testing.T) {
	inputs := [][]byte{
		[]byte("<!DOCTYPE"),
	}
	tests := [][]TokenTest{
		[]TokenTest{
			{token.DOCTYPE, []byte(""), true},
		},
	}

	nextTokenTestFormat(t, inputs, tests)
}
func nextTokenTestFormat(t *testing.T, inputs [][]byte, tests [][]TokenTest) {
	for j, inp := range inputs {
		l := New(inp)
		for i, tt := range tests[j] {
			tok := l.NextToken()
			if tok.Type != tt.expectedType {
				t.Fatalf("tests[%d] : token[%d] - tokentype wrong. expecte=%q, got=%q", j, i, tt.expectedType, tok.Type)
			}
			if !comparator.CmpSlice(tok.Literal, tt.expectedLiteral) {
				t.Fatalf("tests[%d] : token[%d] - tokenLiteral wrong. expecte=%q, got=%q", j, i, tt.expectedLiteral, tok.Literal)
			}
			if tok.ForceQuirks != tt.expectedForceQuirks {
				t.Fatalf("tests[%d] : token[%d] - token flag ForceQuirks wrong. expecte=%t, got=%t", j, i, tt.expectedForceQuirks, tok.ForceQuirks)
			}

		}
	}

}
