package main

import (
	"strings"
)

func (p pnode) String() string {
	return p.label.String()
}

func (p pnode) Pretty() string {
	return p.shelp(0, []int{})
}

func (p pnode) shelp(depth int, split []int) string {
	me := p.label.String()
	mxt := me + `--`
	cdepth := depth + len(mxt)

	switch {
	case p.a != nil && p.b != nil:
		return mxt + p.a.shelp(cdepth, append(split, cdepth)) +
			"\n" + strings.Repeat(" ", cdepth-1) + `\` +
			p.b.shelp(cdepth, split)
	case p.a != nil:
		return mxt + p.a.shelp(cdepth, split)
	case p.b != nil:
		return mxt + p.b.shelp(cdepth, split)
	default:
		return me
	}
}
