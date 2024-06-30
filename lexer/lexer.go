package lexer

import (
	"github/medkhabt/prs/token"
)

// TODO change input to ioRead
type Lexer struct {
	input        []byte
	state        State
	position     int
	readPosition int
	ch           byte // For UTF-8 support use rune.
}

type State interface {
	nextToken() *token.Token
}

func New(input []byte) *Lexer {
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
