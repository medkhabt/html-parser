package lexer

import "github/medkhabt/prs/token"

type AfterAttributeName struct {
	lexer *Lexer
	token *token.Token
}

func (s AfterAttributeName) nextToken() *token.Token {
	s.lexer.readChar()
	if s.lexer.ch == byte(0x09) || s.lexer.ch == byte(0x0A) || s.lexer.ch == byte(0x0C) || s.lexer.ch == byte(0x20) {
		return s.lexer.state.nextToken()
	} else if s.lexer.ch == '>' {
		//TODO check if duplicated attr names exists, before emitting
		s.lexer.state = DataState{s.lexer}
		return s.token
	} else if s.lexer.ch == 0 {
		s.lexer.unreadChar()
		s.lexer.state = DataState{s.lexer}
		return s.lexer.state.nextToken()
	} else {
		ch := s.lexer.ch
		if s.lexer.ch >= 'A' && s.lexer.ch <= 'Z' {
			ch = s.lexer.ch + byte(0x20)
		}
		s.token.Attributes = append(s.token.Attributes, &token.Attribute{string(ch), ""})
		s.lexer.state = AttributeNameState{s.lexer, s.token}
		return s.lexer.state.nextToken()
	}
	return token.NewNotImplemented()
}
