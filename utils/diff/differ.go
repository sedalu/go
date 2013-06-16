package diff

type Differ struct {
	l, r []interface{}
	f *FilterFunc
}

func (p *Differ) SetLeft(left ...interface{}) {
	p.l = left
}

func (p *Differ) SetRight(right ...interface{}) {
	p.r = right
}

func (p *Differ) Filter(filter *FilterFunc) {
	p.f = filter
}

func (p *Differ) Diff() (diffs []Delta) {
	l, r := p.l, p.r

	if p.f != nil {
		l, r = (*p.f)(p.l...), (*p.f)(p.r...)
	}

	return Diff(l, r)
}