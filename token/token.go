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

// TODO implement slice recursive comparaison before changing Attributes fields to []byte
type Attribute struct {
	Name  string
	Value string
}
type Token struct {
	Type       TokenType
	Data       []byte
	Name       []byte
	PublicId   []byte
	SystemId   []byte
	Attributes []*Attribute

	ForceQuirks bool
	SelfClosing bool
}

func NewStartTag(name []byte) *Token {
	return &Token{STARTTAG, nil, name, nil, nil, []*Attribute{}, false, false}
}
func NewEndTag(name []byte) *Token {
	return &Token{ENDTAG, nil, name, nil, nil, []*Attribute{}, false, false}
}
func NewCharacter(data []byte) *Token {
	return &Token{CHARACTER, data, nil, nil, nil, nil, false, false}
}
func NewEOF() *Token {
	return &Token{EOF, nil, nil, nil, nil, nil, false, false}
}
func NewComment(data []byte) *Token {
	return &Token{COMMENT, data, nil, nil, nil, nil, false, false}
}
func NewNotImplemented() *Token {
	return &Token{NOTIMPLEMENTED, nil, nil, nil, nil, nil, false, false}
}
func NewDoctype() *Token {
	return &Token{DOCTYPE, nil, nil, nil, nil, nil, false, false}
}
