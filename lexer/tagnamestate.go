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
	} else {
		if s.lexer.ch >= 'A' && s.lexer.ch <= 'Z' {
			// lowercase characters
			s.lexer.ch += byte(0x20)
		}
		s.token.Literal = append(s.token.Literal, s.lexer.ch)
		return s.nextToken()
	}
}
