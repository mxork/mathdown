package parse

import (
	"mathdown/lex"
)

type Phrase interface {
	String() string
	Rx() string
	Children() []Phrase
}

type Expr struct {
	Phrase
}

func (e Expr) Rx() string {
	return `EXPR`
}

type Term struct {
	lex.Token
}

func (t Term) Rx() string {
	return t.Ilk.String()
}

func (t Term) String() string {
	return t.Rx()
}

func (t Term) Children() []Phrase {
	return nil
}

type Word struct {
	*Term
}

func (w Word) String() string {
	return w.Rx()
}

func (w Word) Rx() string {
	return `EXPR`
}

func (w Word) Children() []Phrase {
	return []Phrase{w.Term}
}

type BinExpr struct {
	*Term
	Left, Right Expr
}

func NewBinExpr(src []Phrase) *BinExpr {
	return &BinExpr{src[1].(*Term), Expr{src[0]}, Expr{src[2]}}
}

func (b BinExpr) Rx() string {
	return `EXPR`
}

func (b BinExpr) String() string {
	return `EXPR`
}

func (b BinExpr) Children() []Phrase {
	return []Phrase{b.Left, b.Term, b.Right}
}

type TagWord struct {
	BinExpr
}

type Statement struct {
	BinExpr
}

func (s Statement) Rx() string {
	return `STMT`
}

func (s Statement) String() string {
	return s.Rx()
}
