package lexer

import "github/medkhabt/prs/token"

type CommentState struct {
	lexer *Lexer
	token *token.Token
}

func (s CommentState) nextToken() *token.Token {
	s.lexer.readChar()
	if s.lexer.ch == '-' {
		// comment end dash
	} else if s.lexer.ch == 0 {
		s.lexer.unreadChar()
		s.lexer.state = DataState{s.lexer}
		return s.token
	} else {
		s.token.Data = append(s.token.Data, s.lexer.ch)
		return s.lexer.state.nextToken()
	}
	return token.NewNotImplemented()
}
