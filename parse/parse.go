package parse

import (
	"mathdown/lex"
)

type Parser struct {
	main, alt []Phrase
	text      string
	is        map[int]int
}

const NOLIMIT = -1

func (p *Parser) apply(r rule) (changed bool) {
	s := p.textify()
	locs := r.FindAllStringIndex(s, NOLIMIT)
	if locs == nil {
		return
	}

	last := 0
	for _, loc := range locs {
		// find index in string, cvt to index in []Phrase,
		start, end := p.is[loc[0]], p.is[loc[1]+1] // off by one

		// take off chaff at the top, and reduce
		p.alt = append(p.alt, p.main[last:start]...)
		p.alt = append(p.alt, r.reduce(p.main[start:end]))
		last = end
	}

	//clean-up
	p.alt = append(p.alt, p.main[last:]...)
	changed = true

	// swap
	p.main, p.alt = p.alt, p.main[0:0]
	return
}

// encodes the current phrase array as string so
// I can leverage the regex package
// Pretty much just strings.Join(Stringer...) if that function
// existed
func (p *Parser) textify() string {
	// clear it out
	s := ""
	for k := range p.is {
		delete(p.is, k)
	}

	// build a text, keep track of token starts
	for i, v := range p.main {
		p.is[len(s)] = i
		s += v.Rx()
		if i != len(p.main)-1 {
			s += " "
		}
	}
	// and a final post
	p.is[len(s)+1] = len(p.main)

	p.text = s
	return s
}
func New(toks []lex.Token) *Parser {
	main := make([]Phrase, 0, len(toks))
	alt := make([]Phrase, 0, len(toks))

	for _, v := range toks {
		main = append(main, &Term{v})
	}

	return &Parser{main, alt, "", map[int]int{}}
}
