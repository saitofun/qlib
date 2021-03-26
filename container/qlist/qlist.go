package qlist

type Element struct {
	n, p *Element
	v    interface{}
}

type List struct {
}
