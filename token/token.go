package token

const (
	DOCTYPE        = "DOCTYPE"
	STARTTAG       = "STARTTAG"
	ENDTAG         = "ENDTAG"
	COMMENT        = "COMMENT"
	CHARACTER      = "CHARACTER"
	EOF            = "EOF"
	NOTIMPLEMENTED = "NOTIMPLEMENTED"
)

type TokenType string

type Token struct {
	Type    TokenType
	Literal []byte
}

//DOCTYPE, start tag, end tag, comment, character, end-of-file
