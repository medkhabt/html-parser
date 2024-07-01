package lexer

import "github/medkhabt/prs/token"

type BeforeAttributeNameState struct {
	lexer *Lexer
	token *token.Token
}

func (s BeforeAttributeNameState) nextToken() *token.Token {
	s.lexer.readChar()
	if s.lexer.ch == byte(0x09) || s.lexer.ch == byte(0x0A) || s.lexer.ch == byte(0x0C) || s.lexer.ch == byte(0x20) {
		// ignore character
		return s.lexer.state.nextToken()
	} else if s.lexer.ch == '/' {
		// switch to self-closing start tag state
	} else if s.lexer.ch == '>' {
		// switch to the data state
		s.lexer.state = DataState{s.lexer}
		return s.token
	}
	return token.NewNotImplemented()
}
