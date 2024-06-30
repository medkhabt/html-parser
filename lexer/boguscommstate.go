package lexer

import "github/medkhabt/prs/token"

type BogusCommentState struct {
	lexer *Lexer
}

func (s BogusCommentState) nextToken() *token.Token {
	t := &token.Token{token.COMMENT, []byte{}}
	for (s.lexer.ch != 0) && (s.lexer.ch != '>') {
		t.Literal = append(t.Literal, s.lexer.ch)
		s.lexer.readChar()
	}
	if s.lexer.ch == 0 {
		s.lexer.unreadChar()
	}
	s.lexer.state = DataState{s.lexer}
	return t
}
