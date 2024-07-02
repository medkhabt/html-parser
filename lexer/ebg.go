package lexer

import "github/medkhabt/prs/token"

type EndBangState struct {
	lexer *Lexer
	token *token.Token
}

func (s EndBangState) nextToken() *token.Token {
	s.lexer.readChar()
	if s.lexer.ch == '>' || s.lexer.ch == 0 {
		if s.lexer.ch == 0 {
			s.lexer.unreadChar()
		}
		s.lexer.state = DataState{s.lexer}
		return s.token
	} else if s.lexer.ch == '-' {
		s.token.Data = append(s.token.Data, '-', '-', '!')
		s.lexer.state = CommentEndDashState{s.lexer, s.token}
		return s.lexer.state.nextToken()
	} else {
		s.token.Data = append(s.token.Data, '-', '-', '!', s.lexer.ch)
		s.lexer.state = CommentState{s.lexer, s.token}
		return s.lexer.state.nextToken()
	}
}
