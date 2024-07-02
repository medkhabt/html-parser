package lexer

import "github/medkhabt/prs/token"

type TagNameState struct {
	lexer *Lexer
	token *token.Token
}

func (s TagNameState) nextToken() *token.Token {
	s.lexer.readChar()

	if s.lexer.ch == '>' {
		s.lexer.state = DataState{s.lexer}
		return s.token
	} else if s.lexer.ch == byte(0x09) || s.lexer.ch == byte(0x0A) || s.lexer.ch == byte(0x0C) || s.lexer.ch == byte(0x20) {
		// switch to before attribute name state, also passs the token
		s.lexer.state = BeforeAttributeNameState{s.lexer, s.token}
		return s.lexer.state.nextToken()
	} else if s.lexer.ch == '/' {
		// self closing start tag state
		s.lexer.state = SelfClosingStartTag{s.lexer, s.token}
		return s.lexer.state.nextToken()
	} else if s.lexer.ch == 0 {
		// reconsume eof in data state
	} else {
		if s.lexer.ch >= 'A' && s.lexer.ch <= 'Z' {
			// lowercase characters
			s.lexer.ch += byte(0x20)
		}
		s.token.Name = append(s.token.Name, s.lexer.ch)
		return s.nextToken()
	}
	return token.NewNotImplemented()
}
