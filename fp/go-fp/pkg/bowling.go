package pkg

import "fmt"

type HitPin int

func (p HitPin) add(l HitPin) HitPin {
	return p + l
}
func (p HitPin) add10() HitPin {
	return p + 10
}
func (p HitPin) addDouble() HitPin {
	return p + 20
}
func (p HitPin) toScore() Score {
	return Score(p)
}

type FrameState interface {
	toString() string
}

// impl FrameState
// first state of a frame
type NotThrow struct{}

func (n NotThrow) toString() string {
	return ""
}

func firstThrow(init NotThrow, first HitPin) FrameState {
	if first == 10 {
		return Strike{}
	}
	return ThrowedFirst{first: first}
}

// impl FrameState
type ThrowedFirst struct {
	first HitPin
}

func (t ThrowedFirst) toString() string {
	return fmt.Sprintf("%d", t.first)
}
func throwSecond(first ThrowedFirst, second HitPin) FrameState {
	if first.first.add(second) == 10 {
		return Spare{
			first:  first.first,
			second: second,
		}
	}
	return ThrowedSecond{
		first:  first.first,
		second: second,
	}
}

type FinalFrameState interface {
	score() Score
}
type FinalRoundNotThrow struct{}

func (n FinalRoundNotThrow) score() Score {
	return 0
}

type FinalRoundFirst struct {
	first HitPin
}

func (f FinalRoundFirst) score() Score {
	return f.first.toScore()
}

type FinalRoundSecond struct {
	first  HitPin
	second HitPin
}

func (f FinalRoundSecond) score() Score {
	return f.first.add(f.second).toScore()
}

func canThirdThrow(f FinalRoundSecond) bool {
	return f.first.add(f.second) >= 10
}
func toFinalRoundThird(f FinalRoundSecond, third HitPin) FinalFrameState {
	return FinalRoundThird{
		first:  f.first,
		second: f.second,
		third:  third,
	}
}

type FinalRoundThird struct {
	first  HitPin
	second HitPin
	third  HitPin
}

func (f FinalRoundThird) score() Score {
	return f.first.add(f.second).add(f.third).toScore()
}

// impl FrameState
type ThrowedSecond struct {
	first  HitPin
	second HitPin
}

func (t ThrowedSecond) toString() string {
	return fmt.Sprintf("%d %d", t.first, t.second)
}
func (t ThrowedSecond) score() Score {
	return t.first.add(t.second).toScore()
}

// impl FrameState
type Spare struct {
	first  HitPin
	second HitPin
}

func (s Spare) toString() string {
	return fmt.Sprintf("%d /", s.first)
}

type Strike struct{}

func (s Strike) toString() string {
	return "X"
}

type Frame struct {
	State FrameState
}
type Frames struct {
	normals []Frame
	final   FinalFrameState
}

func newDefaultGameFrame() Frames {
	frames := make([]Frame, 9)
	for i := 0; i < 9; i++ {
		frames[i] = Frame{State: NotThrow{}}
	}
	return Frames{
		normals: frames,
		final:   FinalRoundNotThrow{},
	}
}

type Score int

func (s Score) add(l Score) Score {
	return s + l
}

func (p HitPin) score() Score {
	return Score(p)
}

// FrameScore is a score of a frame
// frame score is determined or not determined
type FrameScore interface {
	score() Score
}

// impl FrameScore
type DeterminableScore struct {
	value Score
}

func (d DeterminableScore) score() Score {
	return d.value
}

// impl FrameScore
type NotDeterminableScore struct{}

func (n NotDeterminableScore) score() Score {
	return Score(0)
}

func spareScore(frames Frames, frameIndex int) FrameScore {
	if frameIndex == len(frames.normals)-1 {
		switch frames.final.(type) {
		case FinalRoundNotThrow:
			return NotDeterminableScore{}
		default:
			final := frames.final
			switch final.(type) {
			case FinalRoundFirst:
				return DeterminableScore{
					value: final.(FinalRoundFirst).first.add10().toScore(),
				}
			case FinalRoundSecond:
				return DeterminableScore{
					value: final.(FinalRoundSecond).first.add10().toScore(),
				}
			case FinalRoundThird:
				return DeterminableScore{
					value: final.(FinalRoundThird).first.add10().toScore(),
				}
			default:
				return NotDeterminableScore{}
			}
		}
	}
	frameState := frames.normals[frameIndex+1].State
	switch frameState.(type) {
	case Strike:
		return DeterminableScore{value: 20}
	case ThrowedFirst:
		return DeterminableScore{value: frameState.(ThrowedFirst).first.add10().toScore()}
	case ThrowedSecond:
		return DeterminableScore{value: frameState.(ThrowedSecond).first.add10().toScore()}
	case Spare:
		return DeterminableScore{value: frameState.(Spare).first.add10().toScore()}
	case NotThrow:
		return NotDeterminableScore{}
	default:
		return NotDeterminableScore{}
	}
}

func double(nextState FrameState) FrameScore {
	triple := func() DeterminableScore {
		return DeterminableScore{value: 30}
	}
	switch nextState.(type) {
	case ThrowedFirst:
		return DeterminableScore{value: nextState.(ThrowedFirst).first.addDouble().toScore()}
	case ThrowedSecond:
		return DeterminableScore{value: nextState.(ThrowedSecond).first.addDouble().toScore()}
	case Spare:
		return DeterminableScore{value: nextState.(Spare).first.addDouble().toScore()}
	case Strike:
		return triple()
	default:
		return NotDeterminableScore{}
	}
}

func strikeScore(f Frames, index int) FrameScore {
	if index == len(f.normals)-1 {
		final := f.final
		switch final.(type) {
		case FinalRoundNotThrow:
			return NotDeterminableScore{}
		case FinalRoundFirst:
			return NotDeterminableScore{}
		case FinalRoundSecond:
			return DeterminableScore{value: final.(FinalRoundSecond).second.add10().toScore()}
		case FinalRoundThird:
			return DeterminableScore{value: final.(FinalRoundThird).first.add(final.(FinalRoundThird).second).add10().toScore()}
		default:
			return NotDeterminableScore{}
		}
	}
	nextThrow := f.normals[index+1].State
	switch nextThrow.(type) {
	case NotThrow:
		// can not determined
		return NotDeterminableScore{}
	case ThrowedFirst:
		// can not determined
		return NotDeterminableScore{}
		// case Double or Triple
	case ThrowedSecond:
		return DeterminableScore{value: nextThrow.(ThrowedSecond).first.add10().toScore()}
	case Strike:
		if index == len(f.normals)-2 {
			final := f.final
			switch final.(type) {
			case FinalRoundNotThrow:
				return NotDeterminableScore{}
			case FinalRoundFirst:
				return DeterminableScore{value: final.(FinalRoundFirst).first.addDouble().toScore()}
			case FinalRoundSecond:
				return DeterminableScore{value: final.(FinalRoundSecond).first.add10().addDouble().toScore()}
			case FinalRoundThird:
				return DeterminableScore{value: final.(FinalRoundThird).first.addDouble().toScore()}
			default:
				return NotDeterminableScore{}
			}
		}
		return double(f.normals[index+2].State)
	case Spare:
		return DeterminableScore{value: 20}
	default:
		panic(fmt.Sprintf("unknown type %T", nextThrow))
	}
}

type PlayerName string
type Player struct {
	Name PlayerName
	Frames
}

func NewPlayer(name PlayerName) Player {
	return Player{
		Name:   name,
		Frames: newDefaultGameFrame(),
	}
}

func throw(p Player, fn ThrowBall) Player {
	result := NewPlayer(p.Name)
loop:
	for i, frame := range p.Frames.normals {
		switch frame.State.(type) {
		case NotThrow:
			result.normals[i].State = firstThrow(frame.State.(NotThrow), fn())
			break loop
		case ThrowedFirst:
			result.normals[i].State = throwSecond(frame.State.(ThrowedFirst), fn())
			break loop

		default:
			result.normals[i].State = frame.State
		}
	}
	final := p.final
	switch final.(type) {
	case FinalRoundNotThrow:
		switch p.normals[len(p.normals)-1].State.(type) {
		case NotThrow:
			return result
		case ThrowedFirst:
			return result
		default:
			result.final = FinalRoundFirst{first: fn()}
			return result
		}
	case FinalRoundFirst:
		result.final = FinalRoundSecond{
			first:  final.(FinalRoundFirst).first,
			second: fn(),
		}
	case FinalRoundSecond:
		if canThirdThrow(final.(FinalRoundSecond)) {
			result.final = toFinalRoundThird(final.(FinalRoundSecond), fn())
		}
	}
	return result

}

type Bowling struct {
	Players         []Player
	throwPlayerTurn int
}

func Play(b Bowling, throwFn ThrowBall) Bowling {
	throwedPlayer := throw(b.throwPlayer(), throwFn)
	// update
	players := b.Players
	players[b.throwPlayerTurn] = throwedPlayer

	return Bowling{
		Players:         players,
		throwPlayerTurn: updatePlayerIndex(b),
	}
}
func updatePlayerIndex(b Bowling) int {
	nextPlayerIndex := (b.throwPlayerTurn + 1) % len(b.Players)
	return nextPlayerIndex
}

// has side effect function
// because it changes the player skill
type ThrowBall func() HitPin

func AllScore(b Bowling) map[PlayerName]Score {
	result := make(map[PlayerName]Score, len(b.Players))
	for _, player := range b.Players {
		result[player.Name] = player.Score()
	}
	return result
}

func (b Bowling) throwPlayer() Player {
	return b.Players[b.throwPlayerTurn]
}

func NewGame(players ...Player) Bowling {
	return Bowling{
		Players:         players,
		throwPlayerTurn: 0,
	}
}

func (f Frames) Score() Score {
	result := Score(0)
loop:
	for i, frame := range f.normals {
		switch frame.State.(type) {
		case Strike:
			score := strikeScore(f, i)
			switch score.(type) {
			case DeterminableScore:
				result = score.(DeterminableScore).score().add(result)
			case NotDeterminableScore:
				break loop
			}
		case Spare:
			score := spareScore(f, i)
			switch score.(type) {
			case DeterminableScore:
				result = score.(DeterminableScore).score().add(result)
			case NotDeterminableScore:
				break loop
			}
		case ThrowedFirst:
			result = frame.State.(ThrowedFirst).first.toScore().add(result)
		case ThrowedSecond:
			result = frame.State.(ThrowedSecond).score().add(result)
		case NotThrow:
			break loop
		}
	}
	result = f.final.score().add(result)
	return Score(result)
}
