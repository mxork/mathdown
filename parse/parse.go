package parse

import (
	"mathdown/lex"
	re "regexp"
)

var rxs = map[*re.Regexp]func([][])

type Phrase interface {
	String() string
	Rx() string
}

type Expr interface {
	Phrase
}

type Term struct {
	lex.Token
}

func (t Term) Rx() string {
	return t.Ilk.String()
}

type Binop struct {
	Term
	Left, Right Expr
}

func (b Binop) Rx() string {
	return "EXPR"
}

type Parser struct {
	Main, Alt []Phrase
	indices map[int]int // maps string index to real
}

func (p Parser) toreal(x int) {
	return p.indices[i]
}

func New(toks []lex.Token) Parser {
	main := make([]Phrase, 0, len(toks))
	alt := make([]Phrase, 0, len(toks))

	for _, v := range toks {
		main = append(main, Term{v})
	}

	return Parser{main, alt}
}

// Ready for shameless abuse of the regexp package? I am.
func ToRx(ps []Phrase) string {
	s := ""
	for i, v := range ps {
		s += v.Rx()
		if i != len(ps)-1 {
			s += " "
		}
	}
	return s
}
