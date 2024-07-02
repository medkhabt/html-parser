package lexer

import "github/medkhabt/prs/token"

type CommentEndDashState struct {
	lexer *Lexer
	token *token.Token
}

func (s CommentEndDashState) nextToken() *token.Token {
	s.lexer.readChar()
	if s.lexer.ch == '-' {
		s.lexer.state = CommentEndState{s.lexer, s.token}
		return s.lexer.state.nextToken()
	} else if s.lexer.ch == 0 {
		s.lexer.state = DataState{s.lexer}
		return s.token
	} else {
		s.token.Data = append(s.token.Data, '-', s.lexer.ch)
		s.lexer.state = CommentState{s.lexer, s.token}
		return s.lexer.state.nextToken()
	}
}
