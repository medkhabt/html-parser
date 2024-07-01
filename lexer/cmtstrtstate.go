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
		// comment start dash state
	} else if s.lexer.ch == 0 || s.lexer.ch == '>' {
		if s.lexer.ch == 0 {
			s.lexer.unreadChar()
		}
		s.lexer.state = DataState{s.lexer}
		return s.token
	} else {
		s.token.Data = append(s.token.Data, s.lexer.ch)
		// change the state to CommentState and pass the token to then new state i would assume.
	}
	return token.NewNotImplemented()
}
