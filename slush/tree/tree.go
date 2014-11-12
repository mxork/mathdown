package main

/*
func fromParseTree(n *pnode) ast {
}
*/

type kind uint

const (
	expr kind = iota
)

type ast *SyntaxNode
type SyntaxNode interface {
	Kind() kind
	Children() []*SyntaxNode
	Plain() string
}

type Expr interface {
	SyntaxNode
}

type BinExpr struct {
	op          string
	left, right Expr
}

type Var struct {
	sigil string
}
func BinOp(n *pnode) (*BinExpr, bool) {
	if n.a == nil || n.b == nil {
		return nil, false
	}
	check := func(a, b) (*BinExpr, bool) {
		if a.label == word &&
			b.a != nil && b.b != nil &&
			b.a.label == op && b.b.label == word {
			return &BinExpr{string(b.a.token.src),
			Var{a, b.b}, true
		}
		return nil, false
	}

	if α, a := check(n.a, n.b); a {
		return α, a
	}
	if β, b := check(n.a, n.b); b {
		return β, b
	}

	return nil, false
}

func condense(n *pnode) *pnode {
	if n.a != nil {
		n.a = condenseop(n.a)
		condense(n.a)
	}
	if n.b != nil {
		n.b = condenseop(n.b)
		condense(n.b)
	}
	return condenseop(n)
}

func condenseop(n *pnode) *pnode {
	if n.a == nil || n.b == nil {
		return n
	}

	check := func(a, b *pnode) *pnode {
		if a.label == word &&
			b.a != nil && b.b != nil &&
			b.a.label == op && b.b.label == word {
			return &pnode{op, a, b.b, b.a.token}
		}
		return nil
	}

	α, β := check(n.a, n.b), check(n.b, n.a)
	if α != nil {
		return α
	}
	if β != nil {
		return β
	}

	return n
}
