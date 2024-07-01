package lexer

import "github/medkhabt/prs/token"

type BogusCommentState struct {
	lexer *Lexer
}

func (s BogusCommentState) nextToken() *token.Token {

	t := token.NewComment([]byte{})
	for (s.lexer.ch != 0) && (s.lexer.ch != '>') {
		t.Data = append(t.Data, s.lexer.ch)
		s.lexer.readChar()
	}
	if s.lexer.ch == 0 {
		s.lexer.unreadChar()
	}
	s.lexer.state = DataState{s.lexer}
	return t
}
