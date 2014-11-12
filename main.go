package main

import (
	"fmt"
	"mathdown/lex"
	"mathdown/parse"
)

func main() {
	l := lex.New([]byte("happy.once+joos`2 = great / 2 ; tt"))
	ts := l.Lex()
	fmt.Println(ts)
	p := parse.New(ts)
	fmt.Println(parse.ToRx(p.Main))
}
