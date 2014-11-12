package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	re "regexp"
)

type token struct {
	label
	src []byte
}

func (t token) String() string {
	return fmt.Sprintf("%v: '%s'\n", t.label, t.src)
}

func (t token) λ() label {
	return t.label
}

type label uint

const (
	word label = iota
	sux
	op
	dec
	space
	tab
	nl
	rpar
	lpar
)

func (l label) String() string {
	return names[l]
}

var names = map[label]string{
	word:  "word",
	sux:   "sux",
	op:    "op",
	dec:   "dec",
	space: "space",
	tab:   "tab",
	nl:    "nl",
	rpar:  "rpar",
	lpar:  "lpar",
}

var rules = map[label]string{
	word:  `[\pN\pL]+`,
	sux:   "[`" + `.]`,
	op:    `[+^/*·-]`,
	dec:   `[=;,:|←⇐⇒→↔∈∀]|<-|->|=>|<=|<=>`,
	space: ` `,
	tab:   "\t",
	nl:    "\n",
	rpar:  `\)`,
	lpar:  `\(`,
}

var res = map[label]*re.Regexp{}

func init() {
	for k, v := range rules {
		res[k] = re.MustCompile(`^` + v) // anchor at the start
	}
}

func toke(c []byte) (token, bool) {
	for label, rex := range res {
		match := rex.Find(c)
		if match != nil {
			return token{label, match}, true
		}
	}
	return token{}, false
}

func tokenize(r io.Reader) []token {
	b, err := ioutil.ReadAll(r)
	ck(err)
	out := make([]token, 0, 50)
	c := b
	for len(c) != 0 {
		t, ok := toke(c)
		if !ok || len(t.src) == 0 {
			log.Fatalf("bad input at char %v, %s\n%v\n%v\n", len(b)-len(c), string(c[0]), string(c), string(b)) // TODO
		}

		out = append(out, t)
		c = c[len(t.src):]
	}

	return out
}

func ck(err error) {
	if err != nil {
		panic(err)
	}
}
