package token

const (
	DOCTYPE   = "DOCTYPE"
	STARTTAG  = "STARTTAG"
	ENDTAG    = "ENDTAG"
	COMMENT   = "COMMENT"
	CHARACTER = "CHARACTER"
	EOF       = "EOF"
)

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

//DOCTYPE, start tag, end tag, comment, character, end-of-file
