package main

import (
	"fmt"
	"io/ioutil"
	"mathdown/lex"
	"mathdown/parse"
	"os"
)

func main() {
	b, _ := ioutil.ReadAll(os.Stdin)
	fmt.Println(string(b))
	l := lex.New(b)
	ts := l.Lex()
	fmt.Println(ts)
	parse.Test(ts)
}
