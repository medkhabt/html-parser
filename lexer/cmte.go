package lexer

import "github/medkhabt/prs/token"

type CommentEndState struct {
	lexer *Lexer
	token *token.Token
}

func (s CommentEndState) nextToken() *token.Token {
	s.lexer.readChar()
	if s.lexer.ch == '!' {
		// switch end bang state
	} else if s.lexer.ch == '>' {
		s.lexer.state = DataState{s.lexer}
		return s.token
	} else if s.lexer.ch == 0 {
		s.lexer.unreadChar()
		s.lexer.state = DataState{s.lexer}
		return s.token
	} else {
		s.token.Data = append(s.token.Data, '-', '-', s.lexer.ch)
		s.lexer.state = CommentState{s.lexer, s.token}
		return s.lexer.state.nextToken()
	}
	return token.NewNotImplemented()
}
