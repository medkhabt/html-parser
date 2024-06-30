package lexer

import (
	"github/medkhabt/prs/comparator"
	"github/medkhabt/prs/token"
	"testing"
)

type TokenTest struct {
	expectedType    token.TokenType
	expectedLiteral []byte
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
			{token.CHARACTER, []byte{'<'}},
			{token.EOF, []byte{'E', 'O', 'F'}},
		},
		[]TokenTest{
			{token.CHARACTER, []byte{'<'}},
			{token.CHARACTER, []byte{'<'}},
			{token.EOF, []byte{'E', 'O', 'F'}},
		},
		[]TokenTest{
			{token.CHARACTER, []byte{'<'}},
			{token.CHARACTER, []byte{'#'}},
			{token.EOF, []byte{'E', 'O', 'F'}},
		},
		[]TokenTest{
			{token.CHARACTER, []byte{'<'}},
			{token.CHARACTER, []byte{'-'}},
			{token.CHARACTER, []byte{'-'}},
			{token.EOF, []byte{'E', 'O', 'F'}},
		},
	}
	nextTokenTestFormat(t, inputs, tests)
}

func TestEndTagOpen(t *testing.T) {
	inputs := [][]byte{[]byte(`</`), []byte(`</>`)}
	tests := [][]TokenTest{
		[]TokenTest{
			{token.CHARACTER, []byte("<")},
			{token.CHARACTER, []byte("/")},
			{token.EOF, []byte("EOF")},
		},
		[]TokenTest{
			{token.EOF, []byte("EOF")},
		},
	}
	nextTokenTestFormat(t, inputs, tests)
}

func TestTagName(t *testing.T) {
	inputs := [][]byte{[]byte(`<a></a>`), []byte(`<A></A>`), []byte(`<h2></H2>`), []byte(`<H2></h2>`), []byte(`<quote></quote>`), []byte(`<QUOTE></QUOTE>`), []byte(`<QuoTE></qUOte>`)}
	tests := [][]TokenTest{
		[]TokenTest{
			{token.STARTTAG, []byte("a")},
			{token.ENDTAG, []byte("a")},
			{token.EOF, []byte{'E', 'O', 'F'}},
		},
		[]TokenTest{
			{token.STARTTAG, []byte("a")},
			{token.ENDTAG, []byte("a")},
			{token.EOF, []byte{'E', 'O', 'F'}},
		},
		[]TokenTest{
			{token.STARTTAG, []byte("h2")},
			{token.ENDTAG, []byte("h2")},
			{token.EOF, []byte{'E', 'O', 'F'}},
		},
		[]TokenTest{
			{token.STARTTAG, []byte("h2")},
			{token.ENDTAG, []byte("h2")},
			{token.EOF, []byte{'E', 'O', 'F'}},
		},
		[]TokenTest{
			{token.STARTTAG, []byte("quote")},
			{token.ENDTAG, []byte("quote")},
			{token.EOF, []byte{'E', 'O', 'F'}},
		},
		[]TokenTest{
			{token.STARTTAG, []byte("quote")},
			{token.ENDTAG, []byte("quote")},
			{token.EOF, []byte{'E', 'O', 'F'}},
		},
		[]TokenTest{
			{token.STARTTAG, []byte("quote")},
			{token.ENDTAG, []byte("quote")},
			{token.EOF, []byte{'E', 'O', 'F'}},
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
			{token.COMMENT, []byte("?random")},
			{token.COMMENT, []byte("?random")},
			{token.EOF, []byte("EOF")},
		},
		[]TokenTest{
			{token.COMMENT, []byte("DOCTYPR")},
			{token.COMMENT, []byte("doctyp002")},
			{token.EOF, []byte("EOF")},
		},
		[]TokenTest{
			{token.COMMENT, []byte("723#?")},
			{token.COMMENT, []byte("#div")},
			{token.EOF, []byte("EOF")},
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

		}
	}

}
