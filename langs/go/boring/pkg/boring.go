package pkg

import "fmt"

type Boring struct {
	frames     []Frame
	frameIndex int
}

func NewBoring() *Boring {
	return &Boring{
		frames: []Frame{NewFrame()},
	}
}

func (b *Boring) Throw(i int) error {
	if len(b.frames) >= 10 && !b.frames[9].Throwable() {
		return fmt.Errorf("frame is %d", len(b.frames))
	}
	frame := b.frames[b.frameIndex]

	// second throw
	if frame.Throwable() {
		err := frame.Throw(i)
		if err != nil {
			return err
		}
		return nil
	}

	// first throw
	var newFrame Frame
	if len(b.frames) == 9 {
		newFrame = NewLastFrame()
	} else {
		newFrame = NewFrame()
	}
	err := newFrame.Throw(i)
	if err != nil {
		return err
	}
	b.frames = append(b.frames, newFrame)
	b.frameIndex++
	return nil
}

func (b *Boring) Score() int {
	result := 0
	for i, frame := range b.frames {
		if isLast(i) {
			result += frame.Score()
			break
		}
		if frame.IsStrike() && i+1 < len(b.frames) && i+1 < 9 {
			next := b.frames[i+1]
			if next.IsStrike() && i+2 < len(b.frames) {
				nextNext := b.frames[i+2]
				result += frame.Score() + next.Score() + nextNext.Score()
				continue
			}
			result += frame.Score() + next.Score()
			continue
		}
		result += frame.Score()
	}
	return result
}

func isLast(i int) bool {
	return i == 9
}

type Frame interface {
	Score() int
	Throw(i int) error
	Throwable() bool
	IsSpare() bool
	IsStrike() bool
}
type LastFrame struct {
	firstThrow  *int
	secondThrow *int
	thirdThrow  *int
}

func (f *LastFrame) Score() int {
	result := 0
	if f.firstThrow != nil {
		result += *f.firstThrow
	}
	if f.secondThrow != nil {
		result += *f.secondThrow
	}
	if f.thirdThrow != nil {
		result += *f.thirdThrow
	}
	return result
}
func (f *LastFrame) Throwable() bool {
	return f.firstThrow == nil || f.secondThrow == nil || f.IsStrike() || f.IsSpare()
}

func (f *LastFrame) Throw(i int) error {
	if i > 10 {
		return fmt.Errorf("throw is %d", i)
	}
	if !f.Throwable() {
		return fmt.Errorf("can not throw %v", *f)
	}
	if f.firstThrow == nil {
		f.firstThrow = &i
		return nil
	}
	if f.secondThrow == nil {
		f.secondThrow = &i
		return nil
	}
	f.thirdThrow = &i
	return nil
}

func (f *LastFrame) IsSpare() bool {
	return f.firstThrow != nil && f.secondThrow != nil && *f.firstThrow+*f.secondThrow == 10 && !f.IsStrike()
}
func (f *LastFrame) IsStrike() bool {
	if f.thirdThrow != nil {
		return *f.thirdThrow == 10
	}
	if f.secondThrow != nil {
		return *f.secondThrow == 10
	}
	return f.firstThrow != nil && *f.firstThrow == 10
}

func NewLastFrame() *LastFrame {
	return &LastFrame{}
}

type NormalFrame struct {
	firstThrow  *int
	secondThrow *int
}

func NewFrame() *NormalFrame {
	return &NormalFrame{}
}

func (f *NormalFrame) Throwable() bool {
	return f.secondThrow == nil && !f.IsStrike()
}
func (f *NormalFrame) Throw(i int) error {
	if i > 10 {
		return fmt.Errorf("throw is %d", i)
	}
	if f.firstThrow == nil {
		f.firstThrow = &i
		return nil
	}
	if f.Throwable() {
		f.secondThrow = &i
		return nil
	}
	return fmt.Errorf("throw is %d", i)
}

func (f *NormalFrame) IsSpare() bool {
	return f.firstThrow != nil && f.secondThrow != nil && *f.firstThrow+*f.secondThrow == 10 && !f.IsStrike()
}
func (f *NormalFrame) isSecondThrow() bool {
	return f.secondThrow == nil && !f.IsStrike()
}

func (f *NormalFrame) IsStrike() bool {
	return f.firstThrow != nil && *f.firstThrow == 10
}

func (f *NormalFrame) Score() int {
	if f.firstThrow == nil {
		return 0
	}
	if f.secondThrow == nil {
		return *f.firstThrow
	}
	return *f.firstThrow + *f.secondThrow
}
