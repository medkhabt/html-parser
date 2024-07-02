package lexer

import "github/medkhabt/prs/token"

type CommentStartState struct {
	lexer *Lexer
	token *token.Token
}

// TODO differenciate bewteen null and EOF
func (s CommentStartState) nextToken() *token.Token {
	s.lexer.readChar()
	if s.lexer.ch == '-' {
		s.lexer.state = CommentStartDashState{s.lexer, s.token}
		return s.lexer.state.nextToken()
	} else if s.lexer.ch == 0 || s.lexer.ch == '>' {
		if s.lexer.ch == 0 {
			s.lexer.unreadChar()
		}
		s.lexer.state = DataState{s.lexer}
		return s.token
	} else {
		s.token.Data = append(s.token.Data, s.lexer.ch)
		s.lexer.state = CommentState{s.lexer, s.token}
		return s.lexer.state.nextToken()
	}
	return token.NewNotImplemented()
}
