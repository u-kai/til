package pkg

type Boring struct {
	score int
}

func NewBoring() *Boring {
	return &Boring{}
}

func (b *Boring) Slow(i int) {
	b.score += i
}

func (b *Boring) Score() int {
	return b.score
}
