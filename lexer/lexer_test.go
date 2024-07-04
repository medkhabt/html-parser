package lexer

import (
	"fmt"
	"github/medkhabt/prs/comparator"
	"github/medkhabt/prs/token"
	"runtime"
	"strings"
	"testing"
)

type TokenTest struct {
	expectedType       token.TokenType
	expectedData       []byte
	expectedName       []byte
	expectedPublicId   []byte
	expectedSystemId   []byte
	expectedAttributes []*token.Attribute

	expectedForceQuirks bool
	expectedSelfClosing bool
}

func TestTagOpen(t *testing.T) {
	inputs := [][]byte{[]byte("<"), []byte(`<<`), []byte(`<#`), []byte(`<--`)}
	tests := [][]*TokenTest{
		[]*TokenTest{
			newEmpty(token.CHARACTER).data([]byte{'<'}),
			newEmpty(token.EOF),
		},
		[]*TokenTest{
			newEmpty(token.CHARACTER).data([]byte{'<'}),
			newEmpty(token.CHARACTER).data([]byte{'<'}),
			newEmpty(token.EOF),
		},
		[]*TokenTest{
			newEmpty(token.CHARACTER).data([]byte{'<'}),
			newEmpty(token.CHARACTER).data([]byte{'#'}),
			newEmpty(token.EOF),
		},
		[]*TokenTest{
			newEmpty(token.CHARACTER).data([]byte{'<'}),
			newEmpty(token.CHARACTER).data([]byte{'-'}),
			newEmpty(token.CHARACTER).data([]byte{'-'}),
			newEmpty(token.EOF),
		},
	}
	nextTokenTestFormat(t, inputs, tests, getFunctionName())
}

func TestEndTagOpen(t *testing.T) {
	inputs := [][]byte{[]byte(`</`), []byte(`</>`)}
	tests := [][]*TokenTest{
		[]*TokenTest{
			newEmpty(token.CHARACTER).data([]byte("<")),
			newEmpty(token.CHARACTER).data([]byte("/")),
			newEmpty(token.EOF),
		},
		[]*TokenTest{
			newEmpty(token.EOF),
		},
	}
	nextTokenTestFormat(t, inputs, tests, getFunctionName())
}

func TestTagName(t *testing.T) {
	inputs := [][]byte{[]byte(`<a></a>`), []byte(`<A></A>`), []byte(`<h2></H2>`), []byte(`<H2></h2>`), []byte(`<quote></quote>`), []byte(`<QUOTE></QUOTE>`), []byte(`<QuoTE></qUOte>`)}
	tests := [][]*TokenTest{
		[]*TokenTest{
			newEmpty(token.STARTTAG).name([]byte("a")),
			newEmpty(token.ENDTAG).name([]byte("a")),
			newEmpty(token.EOF),
		},
		[]*TokenTest{
			newEmpty(token.STARTTAG).name([]byte("a")),
			newEmpty(token.ENDTAG).name([]byte("a")),
			newEmpty(token.EOF),
		},
		[]*TokenTest{
			newEmpty(token.STARTTAG).name([]byte("h2")),
			newEmpty(token.ENDTAG).name([]byte("h2")),
			newEmpty(token.EOF),
		},
		[]*TokenTest{
			newEmpty(token.STARTTAG).name([]byte("h2")),
			newEmpty(token.ENDTAG).name([]byte("h2")),
			newEmpty(token.EOF),
		},
		[]*TokenTest{
			newEmpty(token.STARTTAG).name([]byte("quote")),
			newEmpty(token.ENDTAG).name([]byte("quote")),
			newEmpty(token.EOF),
		},
		[]*TokenTest{
			newEmpty(token.STARTTAG).name([]byte("quote")),
			newEmpty(token.ENDTAG).name([]byte("quote")),
			newEmpty(token.EOF),
		},
		[]*TokenTest{
			newEmpty(token.STARTTAG).name([]byte("quote")),
			newEmpty(token.ENDTAG).name([]byte("quote")),
			newEmpty(token.EOF),
		},
	}
	nextTokenTestFormat(t, inputs, tests, getFunctionName())
}

func TestBogusCommentState(t *testing.T) {
	inputs := [][]byte{
		[]byte("<?random><?random"),
		[]byte("<!DOCTYPR><!doctyp002"),
		[]byte("</723#?></#div"),
	}
	tests := [][]*TokenTest{
		[]*TokenTest{
			newEmpty(token.COMMENT).data([]byte("?random")),
			newEmpty(token.COMMENT).data([]byte("?random")),
			newEmpty(token.EOF),
		},
		[]*TokenTest{
			newEmpty(token.COMMENT).data([]byte("DOCTYPR")),
			newEmpty(token.COMMENT).data([]byte("doctyp002")),
			newEmpty(token.EOF),
		},
		[]*TokenTest{
			newEmpty(token.COMMENT).data([]byte("723#?")),
			newEmpty(token.COMMENT).data([]byte("#div")),
			newEmpty(token.EOF),
		},
	}
	nextTokenTestFormat(t, inputs, tests, getFunctionName())
}

func TestCommentStart(t *testing.T) {
	inputs := [][]byte{
		[]byte("<!--><!--"),
	}
	tests := [][]*TokenTest{
		[]*TokenTest{
			newEmpty(token.COMMENT).data([]byte{}),
			newEmpty(token.COMMENT).data([]byte{}),
			newEmpty(token.EOF),
		},
	}
	nextTokenTestFormat(t, inputs, tests, getFunctionName())
}

func TestDoctype(t *testing.T) {
	inputs := [][]byte{
		[]byte("<!DOCTYPE"),
	}
	tests := [][]*TokenTest{
		[]*TokenTest{
			newEmpty(token.DOCTYPE).setForceQuirksFlag(),
			newEmpty(token.EOF),
		},
	}

	nextTokenTestFormat(t, inputs, tests, getFunctionName())
}

func TestBeforeAttrName(t *testing.T) {
	// EOF OR >
	inputs := [][]byte{
		[]byte("<test ><test "),
		[]byte("</test ></test "),
		[]byte("<test  / >"),
		[]byte("<test name=\"value\" >"), // TODO need After Attribute value quoted
	}
	tests := [][]*TokenTest{
		[]*TokenTest{
			newEmpty(token.STARTTAG).name([]byte("test")),
			newEmpty(token.STARTTAG).name([]byte("test")),
			newEmpty(token.EOF),
		},
		[]*TokenTest{
			newEmpty(token.ENDTAG).name([]byte("test")),
			newEmpty(token.ENDTAG).name([]byte("test")),
			newEmpty(token.EOF),
		},
		[]*TokenTest{
			newEmpty(token.STARTTAG).name([]byte("test")),
			newEmpty(token.EOF),
		},
		[]*TokenTest{
			newEmpty(token.STARTTAG).name([]byte("test")).attribute("name", "value"),
			newEmpty(token.EOF),
		},
	}
	nextTokenTestFormat(t, inputs, tests, getFunctionName())
}

func TestSelfClosingStartTag(t *testing.T) {
	inputs := [][]byte{
		[]byte("<test /><test   /><test /"),
		[]byte("<test/><test/"),
		[]byte("<test//><test / /><test / /"),
		[]byte("<test name/>"),
		[]byte("<test name=\"value\"/>"),
	}
	tests := [][]*TokenTest{
		[]*TokenTest{
			newEmpty(token.STARTTAG).name([]byte("test")).setSelfClosingFlag(),
			newEmpty(token.STARTTAG).name([]byte("test")).setSelfClosingFlag(),
			newEmpty(token.EOF),
		},
		[]*TokenTest{
			newEmpty(token.STARTTAG).name([]byte("test")).setSelfClosingFlag(),
			newEmpty(token.EOF),
		},
		[]*TokenTest{
			newEmpty(token.STARTTAG).name([]byte("test")).setSelfClosingFlag(),
			newEmpty(token.STARTTAG).name([]byte("test")).setSelfClosingFlag(),
			newEmpty(token.EOF),
		},
		[]*TokenTest{
			newEmpty(token.STARTTAG).name([]byte("test")).attribute("name", "").setSelfClosingFlag(),
			newEmpty(token.EOF),
		},
		[]*TokenTest{
			newEmpty(token.STARTTAG).name([]byte("test")).attribute("name", "value").setSelfClosingFlag(),
			newEmpty(token.EOF),
		},
	}
	nextTokenTestFormat(t, inputs, tests, getFunctionName())
}

func TestComment(t *testing.T) {
	//TODO add Tests for flow that comes Comment_end and Comment_end_dash
	inputs := [][]byte{
		[]byte("<!--test"),
		[]byte("<!---test"),
		[]byte("<!--t"),
		[]byte("<!---t"),
		[]byte("<!--t-t"),  // from comment_end_dash
		[]byte("<!--t--t"), // from comment_end passing by enddash
		[]byte("<!----t"),  // from comment_end passing by startdash
	}
	tests := [][]*TokenTest{
		[]*TokenTest{
			newEmpty(token.COMMENT).data([]byte{'t', 'e', 's', 't'}),
			newEmpty(token.EOF),
		},
		[]*TokenTest{
			newEmpty(token.COMMENT).data([]byte{'-', 't', 'e', 's', 't'}),
			newEmpty(token.EOF),
		},
		[]*TokenTest{
			newEmpty(token.COMMENT).data([]byte{'t'}),
			newEmpty(token.EOF),
		},
		[]*TokenTest{
			newEmpty(token.COMMENT).data([]byte{'-', 't'}),
			newEmpty(token.EOF),
		},
		[]*TokenTest{
			newEmpty(token.COMMENT).data([]byte{'t', '-', 't'}),
			newEmpty(token.EOF),
		},
		[]*TokenTest{
			newEmpty(token.COMMENT).data([]byte{'t', '-', '-', 't'}),
			newEmpty(token.EOF),
		},
		[]*TokenTest{
			newEmpty(token.COMMENT).data([]byte{'-', '-', 't'}),
			newEmpty(token.EOF),
		},
	}
	nextTokenTestFormat(t, inputs, tests, getFunctionName())
}
func TestCommentStartDash(t *testing.T) {
	inputs := [][]byte{
		[]byte("<!---><!---"),
	}
	tests := [][]*TokenTest{
		[]*TokenTest{
			newEmpty(token.COMMENT).data([]byte{}),
			newEmpty(token.COMMENT).data([]byte{}),
			newEmpty(token.EOF),
		},
	}
	nextTokenTestFormat(t, inputs, tests, getFunctionName())
}
func TestCommentEndDash(t *testing.T) {
	inputs := [][]byte{
		[]byte("<!--t-"),    // comment -> .
		[]byte("<!--t-t-"),  // comment -> . -> comment -> .
		[]byte("<!--t--t-"), // comment -> . -> comment-end -> comment -> .
		[]byte("<!--t--!-"),
		[]byte("<!--t--!t-"),
		[]byte("<!----!t-"),
	}
	tests := [][]*TokenTest{
		[]*TokenTest{
			newEmpty(token.COMMENT).data([]byte{'t'}),
			newEmpty(token.EOF),
		},
		[]*TokenTest{
			newEmpty(token.COMMENT).data([]byte{'t', '-', 't'}),
			newEmpty(token.EOF),
		},
		[]*TokenTest{
			newEmpty(token.COMMENT).data([]byte{'t', '-', '-', 't'}),
			newEmpty(token.EOF),
		},
		[]*TokenTest{
			newEmpty(token.COMMENT).data([]byte{'t', '-', '-', '!'}),
			newEmpty(token.EOF),
		},
		[]*TokenTest{
			newEmpty(token.COMMENT).data([]byte{'t', '-', '-', '!', 't'}),
			newEmpty(token.EOF),
		},
		[]*TokenTest{
			newEmpty(token.COMMENT).data([]byte{'-', '-', '!', 't'}),
			newEmpty(token.EOF),
		},
	}
	nextTokenTestFormat(t, inputs, tests, getFunctionName())
}
func TestCommentEnd(t *testing.T) {
	inputs := [][]byte{
		[]byte("<!----><!----"),         // comment start dash path
		[]byte("<!--t--><!--t--"),       //comment -> comment-end-dash ->.
		[]byte("<!--t--t--><!--t--t--"), // comment -> comment-end-dash -> commentend -> comment -> comment-end-dash ->.
		[]byte("<!--t--!--><!--t--!--"),
		[]byte("<!--t--!t--><!--t--!t--"),
		[]byte("<!----!--><!----!--"),
		[]byte("<!----!t--><!----!t--"),
	}
	tests := [][]*TokenTest{
		[]*TokenTest{
			newEmpty(token.COMMENT).data([]byte{}),
			newEmpty(token.COMMENT).data([]byte{}),
			newEmpty(token.EOF),
		},
		[]*TokenTest{
			newEmpty(token.COMMENT).data([]byte{'t'}),
			newEmpty(token.COMMENT).data([]byte{'t'}),
			newEmpty(token.EOF),
		},
		[]*TokenTest{
			newEmpty(token.COMMENT).data([]byte{'t', '-', '-', 't'}),
			newEmpty(token.COMMENT).data([]byte{'t', '-', '-', 't'}),
			newEmpty(token.EOF),
		},
		[]*TokenTest{
			newEmpty(token.COMMENT).data([]byte{'t', '-', '-', '!'}),
			newEmpty(token.COMMENT).data([]byte{'t', '-', '-', '!'}),
			newEmpty(token.EOF),
		},
		[]*TokenTest{
			newEmpty(token.COMMENT).data([]byte{'t', '-', '-', '!', 't'}),
			newEmpty(token.COMMENT).data([]byte{'t', '-', '-', '!', 't'}),
			newEmpty(token.EOF),
		},
		[]*TokenTest{
			newEmpty(token.COMMENT).data([]byte{'-', '-', '!'}),
			newEmpty(token.COMMENT).data([]byte{'-', '-', '!'}),
			newEmpty(token.EOF),
		},
		[]*TokenTest{
			newEmpty(token.COMMENT).data([]byte{'-', '-', '!', 't'}),
			newEmpty(token.COMMENT).data([]byte{'-', '-', '!', 't'}),
			newEmpty(token.EOF),
		},
	}
	nextTokenTestFormat(t, inputs, tests, getFunctionName())
}

func TestEndbang(t *testing.T) {
	inputs := [][]byte{
		[]byte("<!--t--!><!--t--!"),
		[]byte("<!--t--!--!><!--t--!t--!>"),
	}
	tests := [][]*TokenTest{
		[]*TokenTest{
			newEmpty(token.COMMENT).data([]byte{'t'}),
			newEmpty(token.COMMENT).data([]byte{'t'}),
			newEmpty(token.EOF),
		},
		[]*TokenTest{
			newEmpty(token.COMMENT).data([]byte{'t', '-', '-', '!'}),
			newEmpty(token.COMMENT).data([]byte{'t', '-', '-', '!', 't'}),
			newEmpty(token.EOF),
		},
	}
	nextTokenTestFormat(t, inputs, tests, getFunctionName())
}

func TestAttributeName(t *testing.T) {
	inputs := [][]byte{
		[]byte("<test name><test name"),
		[]byte("<test name name2>"),
		[]byte("<test name= \"value\"y>"), //TODO requires After attribute value quoted
	}
	tests := [][]*TokenTest{
		[]*TokenTest{
			newEmpty(token.STARTTAG).name([]byte("test")).attribute("name", ""),
			newEmpty(token.EOF),
		},
		[]*TokenTest{
			newEmpty(token.STARTTAG).name([]byte("test")).attribute("name", "").attribute("name2", ""),
			newEmpty(token.EOF),
		},
		[]*TokenTest{
			newEmpty(token.STARTTAG).name([]byte("test")).attribute("name", "value").attribute("y", ""),
			newEmpty(token.EOF),
		},
	}
	nextTokenTestFormat(t, inputs, tests, getFunctionName())
}

func TestAfterAttributeName(t *testing.T) {
	inputs := [][]byte{
		[]byte("<test name  ><test name     "),
	}
	tests := [][]*TokenTest{
		[]*TokenTest{
			newEmpty(token.STARTTAG).name([]byte("test")).attribute("name", ""),
			newEmpty(token.EOF),
		},
	}
	nextTokenTestFormat(t, inputs, tests, getFunctionName())
}

func TestBeforeAttributeValue(t *testing.T) {
	inputs := [][]byte{
		[]byte("<test name=><test name"),
		[]byte("<test name    =       >"), // passing throught AfterAttributeName (spaces after name)
		[]byte("<test name=t >"),          // passing throught AttributeValueUnQuotedState
	}
	tests := [][]*TokenTest{
		[]*TokenTest{
			newEmpty(token.STARTTAG).name([]byte("test")).attribute("name", ""),
			newEmpty(token.EOF),
		},
		[]*TokenTest{
			newEmpty(token.STARTTAG).name([]byte("test")).attribute("name", ""),
			newEmpty(token.EOF),
		},
		[]*TokenTest{
			newEmpty(token.STARTTAG).name([]byte("test")).attribute("name", "t"),
			newEmpty(token.EOF),
		},
	}
	nextTokenTestFormat(t, inputs, tests, getFunctionName())
}

func TestAttributeValueSingleQuoted(t *testing.T) {
	inputs := [][]byte{
		[]byte("<test name='V"),
	}
	tests := [][]*TokenTest{
		[]*TokenTest{
			newEmpty(token.EOF),
		},
	}
	nextTokenTestFormat(t, inputs, tests, getFunctionName())
}

func TestAttributeValueDoubleQuoted(t *testing.T) {
	inputs := [][]byte{
		[]byte("<test name=\"V"),
	}
	tests := [][]*TokenTest{
		[]*TokenTest{
			newEmpty(token.EOF),
		},
	}
	nextTokenTestFormat(t, inputs, tests, getFunctionName())
}

func TestAttributeValueUnquoted(t *testing.T) {
	inputs := [][]byte{
		[]byte("<test name=value><test name=value"),
	}
	tests := [][]*TokenTest{
		[]*TokenTest{
			newEmpty(token.STARTTAG).name([]byte("test")).attribute("name", "value"),
			newEmpty(token.EOF),
		},
	}
	nextTokenTestFormat(t, inputs, tests, getFunctionName())
}

func TestAfterAttributeValueQuoted(t *testing.T) {
	inputs := [][]byte{
		[]byte("<test name=\"value\"><test name=\"value\""),
		[]byte("<test name='value'>"),
	}
	tests := [][]*TokenTest{
		[]*TokenTest{
			newEmpty(token.STARTTAG).name([]byte("test")).attribute("name", "value"),
			newEmpty(token.EOF),
		},
		[]*TokenTest{
			newEmpty(token.STARTTAG).name([]byte("test")).attribute("name", "value"),
			newEmpty(token.EOF),
		},
	}
	nextTokenTestFormat(t, inputs, tests, getFunctionName())
}

// GENERAL
func TestGeneral(t *testing.T) {
	inputs := [][]byte{[]byte("<!-- This is a comment -->")}
	tests := [][]*TokenTest{
		[]*TokenTest{
			newEmpty(token.COMMENT).data([]byte{' ', 'T', 'h', 'i', 's', ' ', 'i', 's', ' ', 'a', ' ', 'c', 'o', 'm', 'm', 'e', 'n', 't', ' '}),
		},
	}
	nextTokenTestFormat(t, inputs, tests, getFunctionName())
	//t.Logf("%v %v %v %v %v", l.NextToken(), l.NextToken(), l.NextToken(), l.NextToken(), l.NextToken())

}

func getFunctionName() string {
	pc, _, _, _ := runtime.Caller(1)
	return strings.Split(runtime.FuncForPC(pc).Name(), ".")[1]
}
func nextTokenTestFormat(t *testing.T, inputs [][]byte, tests [][]*TokenTest, testName string) {
	for j, inp := range inputs {
		l := New(inp)
		for i, tt := range tests[j] {
			tok := l.NextToken()
			assertEqual(t, tok, tt, j, i, testName)
		}
	}

}
func assertEqual(t *testing.T, tok *token.Token, tt *TokenTest, testIndex int, tokenIndex int, testName string) {
	if tok.Type != tt.expectedType {
		t.Fatalf("[%s] ::: tests[%d] : token[%d] - token type wrong. expecte=%q, got=%q", testName, testIndex, tokenIndex, tt.expectedType, tok.Type)
	}
	if !comparator.CmpSlice(tok.Data, tt.expectedData) {
		t.Fatalf("tests[%d] : token[%d] - token data wrong. expecte=%q, got=%q", testIndex, tokenIndex, tt.expectedData, tok.Data)
	}
	if !comparator.CmpSlice(tok.Name, tt.expectedName) {
		t.Fatalf("tests[%d] : token[%d] - token name wrong. expecte=%q, got=%q", testIndex, tokenIndex, tt.expectedName, tok.Name)
	}
	if !comparator.CmpSlice(tok.PublicId, tt.expectedPublicId) {
		t.Fatalf("tests[%d] : token[%d] - token public id wrong. expecte=%q, got=%q", testIndex, tokenIndex, tt.expectedPublicId, tok.PublicId)
	}
	if !comparator.CmpSlice(tok.SystemId, tt.expectedSystemId) {
		t.Fatalf("tests[%d] : token[%d] - token system id wrong. expecte=%q, got=%q", testIndex, tokenIndex, tt.expectedSystemId, tok.SystemId)
	}
	if !comparator.CmpSlicePointers(tok.Attributes, tt.expectedAttributes) {
		attributesFormat := "["
		if len(tok.Attributes) != len(tt.expectedAttributes) {
			attributesFormat = fmt.Sprintf("len not equal got=%d, expecte=%d", len(tok.Attributes), len(tt.expectedAttributes))
		} else {
			for i, v := range tok.Attributes {
				attributesFormat += fmt.Sprintf("[%d] got(%s,%s) :: expected(%s,%s), ", i, v.Name, v.Value, tt.expectedAttributes[i].Name, tt.expectedAttributes[i].Value)
			}
		}
		attributesFormat += "]"
		t.Fatalf("tests[%d] : token[%d] - token attributes wrong. %s", testIndex, tokenIndex, attributesFormat)
	}
	// Just wanted to make it a little more fun, got bored refactoring, it's just a XOR.
	if (tok.ForceQuirks && !tt.expectedForceQuirks) || (!tok.ForceQuirks && tt.expectedForceQuirks) {
		t.Fatalf("tests[%d] : token[%d] - token forceQuirk flag wrong. expecte=%t, got=%t", testIndex, tokenIndex, tt.expectedForceQuirks, tok.ForceQuirks)
	}
	if (tok.SelfClosing && !tt.expectedSelfClosing) || (!tok.SelfClosing && tt.expectedSelfClosing) {
		t.Fatalf("tests[%d] : token[%d] - token selfClosing flag  wrong. expecte=%t, got=%t", testIndex, tokenIndex, tt.expectedSelfClosing, tok.SelfClosing)
	}

}

func newEmpty(t token.TokenType) *TokenTest {
	var attrs []*token.Attribute = nil
	if t == token.STARTTAG || t == token.ENDTAG {
		attrs = []*token.Attribute{}
	}
	return &TokenTest{t, nil, nil, nil, nil, attrs, false, false}
}
func (tt *TokenTest) setSelfClosingFlag() *TokenTest {
	tt.expectedSelfClosing = true
	return tt
}
func (tt *TokenTest) unsetSelfClosingFlag() *TokenTest {
	tt.expectedSelfClosing = false
	return tt
}
func (tt *TokenTest) setForceQuirksFlag() *TokenTest {
	tt.expectedForceQuirks = true
	return tt
}
func (tt *TokenTest) unsetForceQuirksFlag() *TokenTest {
	tt.expectedForceQuirks = false
	return tt
}
func (tt *TokenTest) attribute(name, value string) *TokenTest {
	tt.expectedAttributes = append(tt.expectedAttributes, &token.Attribute{name, value})
	return tt
}
func (tt *TokenTest) systemId(x []byte) *TokenTest {
	tt.expectedSystemId = x
	return tt
}
func (tt *TokenTest) name(x []byte) *TokenTest {
	tt.expectedName = x
	return tt
}
func (tt *TokenTest) data(x []byte) *TokenTest {
	tt.expectedData = x
	return tt
}
func (tt *TokenTest) publicId(x []byte) *TokenTest {
	tt.expectedPublicId = x
	return tt
}
