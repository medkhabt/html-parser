package lexer

import "github/medkhabt/prs/token"

type TagOpenState struct {
	lexer *Lexer
}

func (s TagOpenState) nextToken() *token.Token {
	s.lexer.readChar()
	if s.lexer.ch == '/' {
		s.lexer.state = EndTagOpenState{s.lexer}
		return s.lexer.state.nextToken()
	} else if s.lexer.ch >= 'a' && s.lexer.ch <= 'z' {
		s.lexer.state = TagNameState{s.lexer, &token.Token{token.STARTTAG, []byte{s.lexer.ch}}}
		return s.lexer.state.nextToken()
	} else if s.lexer.ch >= 'A' && s.lexer.ch <= 'Z' {
		s.lexer.state = TagNameState{s.lexer, &token.Token{token.STARTTAG, []byte{s.lexer.ch + byte(0x20)}}}
		return s.lexer.state.nextToken()
	} else if s.lexer.ch == '?' {
		s.lexer.state = BogusCommentState{s.lexer}
		return s.lexer.state.nextToken()
	} else if s.lexer.ch == '!' {
		s.lexer.state = MarkupDeclarationOpen{s.lexer}
		return s.lexer.state.nextToken()
	} else {
		s.lexer.unreadChar()
		s.lexer.state = DataState{s.lexer}
		return &(token.Token{token.CHARACTER, []byte{'<'}})
	}
}
