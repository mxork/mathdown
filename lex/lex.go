package lex

import (
	"fmt"
	re "regexp"
	"unicode/utf8"
)

type Ilk uint

const (
	WORD Ilk = iota
	SUB
	PUNC
	DEC
	UNOP
	BINOP
	RPAR
	LPAR
	NL
	SPACE
	ERRD
)

var ilkstrings = map[Ilk]string{
	WORD:  "WORD",
	SUB:   "SUB",
	PUNC:  "PUNC",
	DEC:   "DEC",
	UNOP:  "UNOP",
	BINOP: "BINOP",
	RPAR:  "RPAR",
	LPAR:  "LPAR",
	NL:    "NL",
	SPACE: "SPACE",
	ERRD:  "ERRD",
}

func (i Ilk) String() string {
	return ilkstrings[i]
}

type rex re.Regexp

func (r *rex) Is(ch rune) bool {
	return (*re.Regexp)(r).Match([]byte(string(ch)))
}

// handy dandy regexs
var (
	alnum  = (*rex)(re.MustCompile(`[\pL\pN]`))
	punct  = (*rex)(re.MustCompile(`[;:|,]`))
	decl   = (*rex)(re.MustCompile(`[=⇒]`))
	binopr = (*rex)(re.MustCompile(`[+-^/]` + "|`"))
)

const EOF rune = 0xFFD

type Token struct {
	Ilk
	A, Z int
}

func (t Token) String() string {
	return fmt.Sprintf("%v(%v,%v)", t.Ilk, t.A, t.Z)
}

// reduce char count
func sing(i Ilk, pos int) Token {
	return Token{i, pos, pos + 1}
}

type Lexer struct {
	pos int
	ch  rune
	src []byte
}

func (l *Lexer) scan() Token {
	t := l.adv()
	//simple cases
	switch t {
	case '.':
		return sing(SUB, l.pos)
	case ' ':
		return sing(SPACE, l.pos)
	case '\n':
		return sing(NL, l.pos)
	case EOF:
		return Token{ERRD, l.pos, l.pos}
	}

	// regex cases
	switch {
	case alnum.Is(t):
		return l.word()
	case binopr.Is(t):
		return sing(BINOP, l.pos)
	case punct.Is(t):
		return sing(PUNC, l.pos)
	case decl.Is(t):
		return sing(DEC, l.pos)
	}

	return sing(ERRD, l.pos)
}

func (l *Lexer) adv() rune {
	l.pos += len(string(l.ch))

	if l.pos >= len(l.src) {
		l.ch, l.pos = EOF, len(l.src)
		return EOF
	}

	l.ch, _ = utf8.DecodeRune(l.src[l.pos:])

	return l.ch
}

func (l *Lexer) peek() rune {
	ch, _ := utf8.DecodeRune(l.src[l.pos:])
	return ch
}

func (l *Lexer) word() (w Token) {
	w.Ilk, w.A = WORD, l.pos
	for ; alnum.Is(l.peek()); l.adv() {
	}
	w.Z = l.pos
	l.pos--
	return
}

func (l *Lexer) Lex() (toks []Token) {
	toks = make([]Token, 0, len(l.src)) // sensible upper limit
	for t := l.scan(); t.Ilk != ERRD; t = l.scan() {
		toks = append(toks, t)
	}
	return
}

func New(source []byte) *Lexer {
	return &Lexer{-1, 0x00, source}
}
