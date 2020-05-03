package lexer

import (
	"unicode/utf8"
)

const (
	EOF rune = -1
)

// A Lexer is the structure to iterate through the input and emit the items
// into the channel
type Lexer struct {
	Input    string
	Start    int
	Position int
	Items    chan Item
}

// emit allows to add the current token to the channel
func (lexer *Lexer) emit(itemType ItemType) {
	lexer.Items <- Item{
		Type: itemType,
		Value: lexer.Input[lexer.Start:lexer.Position],
	}

	lexer.Start = lexer.Position
}

// next moves the position to the next rune and returns it
func (lexer *Lexer) next() rune {
	// Returns EOF if the position if over the length of the input
	if lexer.Position >= len(lexer.Input) {
		return EOF
	}

	// Decodes the first rune in the given input, gets it and its width
	rune, width := utf8.DecodeRuneInString(lexer.Input[lexer.Position:])
	lexer.Position += width

	return rune
}

// run iterate through the runes of the lexer inputs and lex them
func (lexer *Lexer) run() {
	for lexer.Position < len(lexer.Input) {
		lexer.next()
	}
}

// Lex creates a Lexer with the given input, runs it in a go routine and returns the lexer and
// its channel for items
func Lex(input string) (*Lexer, chan Item) {
	lexer := &Lexer{
		Input: input,
		Items: make(chan Item),
	}

	// Go routine for later
	lexer.run()
	return lexer, lexer.Items
}
