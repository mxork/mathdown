package parse

import (
	re "regexp"
)

var rules = []rule{
	R(`WORD`, func(src []Phrase) Phrase {
		return Word{src[0].(*Term)}
	}),
	R(`EXPR SUB EXPR`, func(src []Phrase) Phrase {
		return TagWord{*NewBinExpr(src)}
	}),
	R(`EXPR BINOP EXPR`, func(src []Phrase) Phrase {
		return NewBinExpr(src)
	}),

	// whitepace...
	R(`SPACE EXPR SPACE`, func(src []Phrase) Phrase {
		return src[1]
	}),
	R(`EXPR SPACE`, func(src []Phrase) Phrase {
		return src[0]
	}),
	R(`SPACE EXPR`, func(src []Phrase) Phrase {
		return src[1]
	}),
	R(`EXPR DEC EXPR`, func(src []Phrase) Phrase {
		return Statement{*NewBinExpr(src)}
	}),
}

type rule struct {
	*re.Regexp
	reduce (func([]Phrase) Phrase)
}

func R(s string, f func([]Phrase) Phrase) rule {
	return rule{re.MustCompile(s), f}
}
