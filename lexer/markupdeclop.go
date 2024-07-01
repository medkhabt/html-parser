package lexer

import (
	"github/medkhabt/prs/comparator"
	"github/medkhabt/prs/token"
)

type MarkupDeclarationOpen struct {
	lexer *Lexer
}

func (s MarkupDeclarationOpen) nextToken() *token.Token {
	inputByte := []byte(s.lexer.input)
	s.lexer.readChar()
	if comparator.CmpSlice(inputByte[s.lexer.position:s.lexer.position+2], []byte{'-', '-'}) {
		// 1 character jump
		s.lexer.readChar()
		t := token.NewComment([]byte{})
		s.lexer.state = CommentStartState{s.lexer, t}
		return s.lexer.state.nextToken()
		// comment start state .
	} else if comparator.CmpInsensitiveByteSlice(inputByte[s.lexer.position:s.lexer.position+7], []byte{'D', 'O', 'C', 'T', 'Y', 'P', 'E'}) {
		// 6 charachters jump
		s.lexer.readChar()
		s.lexer.readChar()
		s.lexer.readChar()
		s.lexer.readChar()
		s.lexer.readChar()
		s.lexer.readChar()
		s.lexer.state = DoctypeState{s.lexer}
		return s.lexer.state.nextToken()
	} else {
		s.lexer.state = BogusCommentState{s.lexer}
		return s.lexer.state.nextToken()
		//bogus comment state
	}
	return token.NewNotImplemented()
}
