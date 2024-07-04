package lexer

import "github/medkhabt/prs/token"

type AttributeValueDoubleQuotedState struct {
	lexer *Lexer
	token *token.Token
}

func (s AttributeValueDoubleQuotedState) nextToken() *token.Token {
	s.lexer.readChar()
	if s.lexer.ch == byte(0x22) {
		s.lexer.state = AfterAttributeValueQuotedState{s.lexer, s.token}
		return s.lexer.state.nextToken()
	} else if s.lexer.ch == '&' {
		// character reference in attribute value.
	} else if s.lexer.ch == 0 {
		s.lexer.unreadChar()
		s.lexer.state = DataState{s.lexer}
		return s.lexer.state.nextToken()
	} else {
		s.token.Attributes[len(s.token.Attributes)-1].Value += string(s.lexer.ch)
		return s.lexer.state.nextToken()

	}
	return token.NewNotImplemented()
}
