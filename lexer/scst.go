package lexer

import "github/medkhabt/prs/token"

// TODO add state to the type name !
type SelfClosingStartTag struct {
	lexer *Lexer
	token *token.Token
}

func (s SelfClosingStartTag) nextToken() *token.Token {
	s.lexer.readChar()
	if s.lexer.ch == '>' {
		s.token.SelfClosing = true
		s.lexer.state = DataState{s.lexer}
		return s.token
	} else if s.lexer.ch == 0 {
		s.lexer.unreadChar()
		s.lexer.state = DataState{s.lexer}
		return s.lexer.state.nextToken()
	} else {
		s.lexer.unreadChar()
		s.lexer.state = BeforeAttributeNameState{s.lexer, s.token}
		return s.lexer.state.nextToken()
	}
	return token.NewNotImplemented()
}
