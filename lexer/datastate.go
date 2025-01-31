package lexer

import "github/medkhabt/prs/token"

type DataState struct {
	lexer *Lexer
}

func (s DataState) nextToken() *token.Token {
	s.lexer.readChar()
	switch s.lexer.ch {
	case '<':
		s.lexer.state = TagOpenState{s.lexer}
		return s.lexer.state.nextToken()
	case 0:
		// TODO change the detection of the eof when implementing a io.Read
		return token.NewEOF()
	}
	t := token.NewCharacter([]byte{s.lexer.ch})
	return t
}
