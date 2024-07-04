package lexer

import "github/medkhabt/prs/token"

type BeforeAttributeValueState struct {
	lexer *Lexer
	token *token.Token
}

func (s BeforeAttributeValueState) nextToken() *token.Token {
	s.lexer.readChar()
	if s.lexer.ch == byte(0x09) || s.lexer.ch == byte(0x0A) || s.lexer.ch == byte(0x0C) || s.lexer.ch == byte(0x20) {
		return s.lexer.state.nextToken()
	} else if s.lexer.ch == byte(0x27) {
		s.lexer.state = AttributeValueSingleQuotedState{s.lexer, s.token}
		return s.lexer.state.nextToken()
	} else if s.lexer.ch == byte(0x22) {
		s.lexer.state = AttributeValueDoubleQuotedState{s.lexer, s.token}
		return s.lexer.state.nextToken()
	} else if s.lexer.ch == 0 {
		s.lexer.unreadChar()
		s.lexer.state = DataState{s.lexer}
		return s.lexer.state.nextToken()
	} else if s.lexer.ch == '>' {
		s.lexer.state = DataState{s.lexer}
		return s.token
	} else if s.lexer.ch == '&' {
		s.lexer.unreadChar()
		s.lexer.state = AttributeValueUnquotedState{s.lexer, s.token}
		return s.lexer.state.nextToken()
	} else {
		s.token.Attributes[len(s.token.Attributes)-1].Value += string(s.lexer.ch)
		s.lexer.state = AttributeValueUnquotedState{s.lexer, s.token}
		return s.lexer.state.nextToken()
	}
	return token.NewNotImplemented()
}
