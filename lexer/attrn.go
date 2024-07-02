package lexer

import "github/medkhabt/prs/token"

type AttributeNameState struct {
	lexer *Lexer
	token *token.Token
}

func (s AttributeNameState) nextToken() *token.Token {
	return token.NewNotImplemented()
}
