package main

// parse-node
type pnode struct {
	label
	a, b   *pnode
	*token // if terminal
}

var seqs = []struct{ a, b, x label }{
	{word, word, word},
	{sux, word, word},
	{op, word, word},

	{space, word, word},
	{word, space, word},

	//{tab, space, tab},
	{word, rpar, rpar},
	{lpar, rpar, word},
}

func parse(ts []token) []*pnode {
	// current, output
	sa := make([]*pnode, len(ts))
	sb := make([]*pnode, 0, len(ts))

	for i, t := range ts {
		sa[i] = &pnode{t.label, nil, nil, &t}
	}

	for i := 0; i < len(seqs); i++ {
		s, changed := seqs[i], false

		for len(sa) != 0 {
			if len(sa) == 1 {
				sb = append(sb, sa[0])
				sa = sa[0:0]
				break
			}

			a, b := sa[0], sa[1]
			if a.label == s.a && b.label == s.b {
				changed = true
				sb = append(sb, &pnode{s.x, a, b, nil})
				sa = sa[2:]
			} else {
				sb = append(sb, a)
				sa = sa[1:]
			}

		}

		// go back to hi-priority rules
		if changed {
			i = -1
		}
		sa, sb = sb, sa[0:0]
	}

	for i := range sa {
		sa[i] = removespace(sa[i])
	}

	return sa
}

func onchild(n *pnode, f func(*pnode)) {
	if n.a != nil {
		f(n.a)
	}
	if n.b != nil {
		f(n.b)
	}
}

func dft(n *pnode, f func(*pnode)) {
	onchild(n, func(m *pnode) {
		dft(m, f)
	})

	f(n)
}

func removespace(r *pnode) *pnode {
	if r.a != nil {
		removespace(r.a)
		r.a = trunc(r.a)
	}
	if r.b != nil {
		removespace(r.b)
		r.b = trunc(r.b)
	}

	return trunc(r)
}

func trunc(r *pnode) *pnode {
	if r.a != nil && r.a.label == space {
		return r.b
	}
	if r.b != nil && r.b.label == space {
		return r.a
	}
	return r
}
