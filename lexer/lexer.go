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

// ********* TAGOPENSTATE ******
func (s TagOpenState) nextToken() *token.Token {
	s.lexer.state = DataState{s.lexer}
	return &(token.Token{token.CHARACTER, string('<')})
}

// ********* DATASTATE ******
func (s DataState) nextToken() *token.Token {
	switch s.lexer.ch {
	case '<':
		s.lexer.readChar()
		s.lexer.state = TagOpenState{s.lexer}
		return s.lexer.state.nextToken()
	case 0:
		// TODO change the detection of the eof when implementing a io.Read
		return &token.Token{token.EOF, "EOF"}
	}
	t := &token.Token{token.CHARACTER, string(s.lexer.ch)}
	s.lexer.readChar()
	return t
}
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.state = DataState{l}
	l.readChar()
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
