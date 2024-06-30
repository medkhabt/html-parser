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
	if comparator.CmpSlice(inputByte[s.lexer.position:s.lexer.position+3], []byte{'!', '-', '-'}) {
		// 2 characters jump
		s.lexer.readChar()
		s.lexer.readChar()
		// comment start state .
	} else if comparator.CmpInsensitiveByteSlice(inputByte[s.lexer.position:s.lexer.position+8], []byte{'!', 'D', 'O', 'C', 'T', 'Y', 'P', 'E'}) {
		// 7 charachters jump
		s.lexer.readChar()
		s.lexer.readChar()
		s.lexer.readChar()
		s.lexer.readChar()
		s.lexer.readChar()
		s.lexer.readChar()
		s.lexer.readChar()
		//DOCTYPE State
	} else {
		s.lexer.readChar()
		s.lexer.state = BogusCommentState{s.lexer}
		return s.lexer.state.nextToken()
		//bogus comment state
	}
	return &token.Token{token.NOTIMPLEMENTED, []byte{}}
}
