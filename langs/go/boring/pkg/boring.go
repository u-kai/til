package pkg

import "fmt"

type Boring struct {
	flame int
	score int
}

func NewBoring() *Boring {
	return &Boring{}
}

func (b *Boring) Slow(i int) error {
	if b.flame >= 20 {
		return fmt.Errorf("flame is %d", b.flame)
	}
	b.flame++
	b.score += i
	return nil
}

func (b *Boring) Score() int {
	return b.score
}

type Frame interface {
	Score() int
}

type NormalFrame struct {
	firstThrow  int
	secondThrow int
	throw       bool
}

func NewFlame() *NormalFrame {
	return &NormalFrame{}
}

func (f *NormalFrame) Slow(i int) error {
	if i > 10 {
		return fmt.Errorf("throw is %d", i)
	}
	if !f.throw {
		f.firstThrow = i
		f.throw = true
		return nil
	}
	if f.isSecondThrow() {
		f.secondThrow = i
	}
	return fmt.Errorf("throw is %d", i)
}
func (f *NormalFrame) isSecondThrow() bool {
	return f.throw && !f.isStrike()
}

func (f *NormalFrame) isStrike() bool {
	return f.firstThrow == 10
}

func (n *NormalFrame) Score() int {
	return n.firstThrow + n.secondThrow
}
