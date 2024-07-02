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
		// TODO should i add also token?
		s.lexer.state = SelfClosingStartTag{s.lexer, s.token}
		return s.lexer.state.nextToken()
	} else if s.lexer.ch == '>' {
		// switch to the data state
		s.lexer.state = DataState{s.lexer}
		return s.token
	} else if s.lexer.ch == 0 {
		s.lexer.unreadChar()
		s.lexer.state = DataState{s.lexer}
		return s.token
	} else {
		// Pares error
		// 	if s.lexer.ch == byte(0x27) || s.lexer.ch == byte(0x22) || s.lexer.ch == byte(0x3C) || s.lexer.ch == byte(0x3D) {
		if s.lexer.ch >= 'A' && s.lexer.ch <= 'Z' {
			s.lexer.ch = s.lexer.ch + byte(0x20)
		}
		s.token.Attributes = append(s.token.Attributes, &token.Attribute{string(s.lexer.ch), ""})
		s.lexer.state = AttributeNameState{s.lexer, s.token}
		return s.lexer.state.nextToken()
	}
}
