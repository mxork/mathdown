package main

import (
	"fmt"
	"mathdown/lex"
)

func main() {
	l := lex.New([]byte("happy.once+joos`2 = great / 2 ; tt"))
	ts := l.Lex()
	fmt.Println(ts)
}
