package lexer

import "github/medkhabt/prs/token"

type AfterAttributeValueQuotedState struct {
	lexer *Lexer
	token *token.Token
}

func (s AfterAttributeValueQuotedState) nextToken() *token.Token {
	s.lexer.readChar()
	if s.lexer.ch == byte(0x09) || s.lexer.ch == byte(0x0A) || s.lexer.ch == byte(0x0C) || s.lexer.ch == byte(0x20) {
		s.lexer.state = BeforeAttributeNameState{s.lexer, s.token}
		return s.lexer.state.nextToken()
	} else if s.lexer.ch == '/' {
		s.lexer.state = SelfClosingStartTag{s.lexer, s.token}
		return s.lexer.state.nextToken()
	} else if s.lexer.ch == '>' {
		s.lexer.state = DataState{s.lexer}
		return s.token
	} else if s.lexer.ch == 0 {
		s.lexer.unreadChar()
		s.lexer.state = DataState{s.lexer}
		return s.lexer.state.nextToken()
	} else {
		s.lexer.unreadChar()
		s.lexer.state = BeforeAttributeNameState{s.lexer, s.token}
		return s.lexer.state.nextToken()
	}
	return token.NewNotImplemented()
}
