package lexer

import "github/medkhabt/prs/token"

type EndTagOpenState struct {
	lexer *Lexer
}

func (s EndTagOpenState) nextToken() *token.Token {
	s.lexer.readChar()
	if s.lexer.ch == '>' {
		// ..</> => ignore it.
		s.lexer.state = DataState{s.lexer}
		return s.lexer.state.nextToken()
	} else if s.lexer.ch == 0 {
		// case ...</EOF => return Token('char: <'), and move pointer back to '<', so DataState takes '/'
		//TODO specify how many steps in unread ?
		s.lexer.unreadChar()
		s.lexer.unreadChar()
		s.lexer.state = DataState{s.lexer}
		return token.NewCharacter([]byte{'<'})
	} else if s.lexer.ch >= 'a' && s.lexer.ch <= 'z' {
		s.lexer.state = TagNameState{s.lexer, token.NewEndTag([]byte{s.lexer.ch})}
		return s.lexer.state.nextToken()
	} else if s.lexer.ch >= 'A' && s.lexer.ch <= 'Z' {
		s.lexer.state = TagNameState{s.lexer, token.NewEndTag([]byte{s.lexer.ch + byte(0x20)})}
		return s.lexer.state.nextToken()
	} else {
		s.lexer.state = BogusCommentState{s.lexer}
		return s.lexer.state.nextToken()
	}
}
