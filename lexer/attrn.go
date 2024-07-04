package lexer

import "github/medkhabt/prs/token"

type AttributeNameState struct {
	lexer *Lexer
	token *token.Token
}

func (s AttributeNameState) nextToken() *token.Token {
	s.lexer.readChar()
	if s.lexer.ch == byte(0x09) || s.lexer.ch == byte(0x0A) || s.lexer.ch == byte(0x0C) || s.lexer.ch == byte(0x20) {
		s.lexer.state = AfterAttributeName{s.lexer, s.token}
		return s.lexer.state.nextToken()
	} else if s.lexer.ch == '=' {
		// switch to before attribute value
		s.lexer.state = BeforeAttributeValueState{s.lexer, s.token}
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
		ch := s.lexer.ch
		if s.lexer.ch >= 'A' && s.lexer.ch <= 'Z' {
			ch = s.lexer.ch + byte(0x20)
		}
		s.token.Attributes[len(s.token.Attributes)-1].Name += string(ch)
		return s.lexer.state.nextToken()
	}
	return token.NewNotImplemented()
}
