package lexer

import (
	"github/medkhabt/prs/token"
)

// TODO change input to ioRead
type Lexer struct {
	input        string
	state        State
	position     int
	readPosition int
	ch           byte // For UTF-8 support use rune.
}

// ***** STATES *****
type State interface {
	nextToken() *token.Token
}
type DataState struct {
	lexer *Lexer
}
type TagOpenState struct {
	lexer *Lexer
}
type EndTagOpenState struct {
	lexer *Lexer
}
type TagNameState struct {
	lexer *Lexer
	token *token.Token
}

// ********* TAGNAMESTATE ******

func (s TagNameState) nextToken() *token.Token {
	s.lexer.readChar()

	if s.lexer.ch == '>' {
		s.lexer.state = DataState{s.lexer}
		return s.token
	} else {
		if s.lexer.ch >= 'A' && s.lexer.ch <= 'Z' {
			// lowercase characters
			s.lexer.ch += byte(0x20)
		}
		s.token.Literal += string(s.lexer.ch)
		return s.nextToken()
	}
}

// ********* ENDTAGOPENSTATE ******

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
		return &(token.Token{token.CHARACTER, string('<')})
	} else if s.lexer.ch >= 'a' && s.lexer.ch <= 'z' {
		s.lexer.state = TagNameState{s.lexer, &token.Token{token.ENDTAG, string(s.lexer.ch)}}
		return s.lexer.state.nextToken()
	} else if s.lexer.ch >= 'A' && s.lexer.ch <= 'Z' {
		s.lexer.state = TagNameState{s.lexer, &token.Token{token.ENDTAG, string(s.lexer.ch + byte(0x20))}}
		return s.lexer.state.nextToken()
	} else {
		return &(token.Token{token.NOTIMPLEMENTED, ""})
	}
}

// ********* TAGOPENSTATE ******
func (s TagOpenState) nextToken() *token.Token {
	s.lexer.readChar()
	if s.lexer.ch == '/' {
		s.lexer.state = EndTagOpenState{s.lexer}
		return s.lexer.state.nextToken()
	} else if s.lexer.ch >= 'a' && s.lexer.ch <= 'z' {
		s.lexer.state = TagNameState{s.lexer, &token.Token{token.STARTTAG, string(s.lexer.ch)}}
		return s.lexer.state.nextToken()
	} else if s.lexer.ch >= 'A' && s.lexer.ch <= 'Z' {
		s.lexer.state = TagNameState{s.lexer, &token.Token{token.STARTTAG, string(s.lexer.ch + byte(0x20))}}
		return s.lexer.state.nextToken()
	} else {
		s.lexer.unreadChar()
		s.lexer.state = DataState{s.lexer}
		return &(token.Token{token.CHARACTER, string('<')})
	}
}

// ********* DATASTATE ******
func (s DataState) nextToken() *token.Token {
	s.lexer.readChar()
	switch s.lexer.ch {
	case '<':
		s.lexer.state = TagOpenState{s.lexer}
		return s.lexer.state.nextToken()
	case 0:
		// TODO change the detection of the eof when implementing a io.Read
		return &token.Token{token.EOF, "EOF"}
	}
	t := &token.Token{token.CHARACTER, string(s.lexer.ch)}
	return t
}
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.state = DataState{l}
	return l
}

func (l *Lexer) NextToken() *token.Token {
	return l.state.nextToken()
}

// TODO Support UTF-8
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}
func (l *Lexer) unreadChar() {
	l.readPosition -= 1
	l.position -= 1
	l.ch = l.input[l.position]
}
