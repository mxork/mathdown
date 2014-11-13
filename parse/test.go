package parse

import (
	"fmt"
	"mathdown/lex"
)

func Test(toks []lex.Token) {
	p := New(toks)
	for i := 0; i < len(rules); i++ {
		if p.apply(rules[i]) {
			i = -1
		}
	}
	fmt.Println(p.main)
}
