package lexer

import "github/medkhabt/prs/token"

type DoctypeState struct {
	lexer *Lexer
}

func (s DoctypeState) nextToken() *token.Token {
	s.lexer.readChar()
	if s.lexer.ch == byte(0x09) || s.lexer.ch == byte(0x0A) || s.lexer.ch == byte(0x0C) || s.lexer.ch == byte(0x20) {
		// Switch to Before DOCTYPE name state
	} else if s.lexer.ch == 0 {
		t := &token.Token{token.DOCTYPE, []byte(""), true}
		s.lexer.unreadChar()
		s.lexer.state = DataState{s.lexer}
		return t
	} else {
		s.lexer.unreadChar()
		// switch to before doctype name state
	}
	return &token.Token{token.NOTIMPLEMENTED, []byte(""), false}
}
