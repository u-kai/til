package pkg

import "fmt"

type Turn func(Throw, Bowling) Bowling

type Throw func(Rounds) Rounds

func ThrowByAnyLogic(logic func() ThrowResult) Throw {
	return func(rounds Rounds) Rounds {
		throwed := logic()
		result := NewDefaultRounds()
	loop:
		for i, round := range rounds.Early {
			switch round.throw.(type) {
			case NotThrow:
				result.Early[i].throw = round.throw.(NotThrow).firstThrow(throwed)
				break loop
			case ThrowFirst:
				result.Early[i].throw = round.throw.(ThrowFirst).secondThrow(throwed)
				break loop

			default:
				result.Early[i].throw = round.throw

			}
		}
		return result
	}
}

type ThrowResult int

type Bowling struct {
	Players         []Player
	throwPlayerTurn int
}

type Rounds struct {
	Early []Round
	Final FinalRound
}

func (r Rounds) Score() Score {
	result := 0
loop:
	for _, round := range r.Early {
		switch round.throw.(type) {
		case Strike:
			result += 10
		case ThrowFirst:
			result += round.throw.(ThrowFirst).score
		case ThrowSecond:
			result += round.throw.(ThrowSecond).score()
		case NotThrow:
			break loop
		}
	}
	return Score(result)
}

func NewDefaultRounds() Rounds {
	early := make([]Round, 9)
	for i := 0; i < 9; i++ {
		early[i] = newRound()
	}
	return Rounds{
		Early: early,
		Final: newFinalRound(),
	}
}

type Round struct {
	throw ThrowState
}

func newRound() Round {
	return Round{
		throw: NotThrow{},
	}
}

type FinalRound struct {
	FirstThrow  ThrowResult
	SecondThrow ThrowResult
	ThirdThrow  ThrowResult
}

func newFinalRound() FinalRound {
	return FinalRound{
		FirstThrow:  -1,
		SecondThrow: -1,
		ThirdThrow:  -1,
	}
}
func (r FinalRound) throwedFirst() bool {
	return r.FirstThrow != -1
}
func (r FinalRound) throwedSecond() bool {
	return r.SecondThrow != -1
}
func (r FinalRound) throwedThird() bool {
	return r.ThirdThrow != -1
}

type Player struct {
	Name PlayerName
	Rounds
}

type PlayerName string

type Score int
type RoundScore int

const NotFilled RoundScore = -1

type ThrowState interface {
	ToString() string
}

type Strike struct {
}
type ThrowFirst struct {
	score int
}

func (t ThrowFirst) secondThrow(throwResult ThrowResult) ThrowState {
	if t.score+int(throwResult) == 10 {
		return Spare{}
	}
	return ThrowSecond{t.score, int(throwResult)}
}

type ThrowSecond struct {
	first  int
	second int
}

func (t ThrowSecond) score() int {
	return t.first + t.second
}

type Spare struct{}

type NotThrow struct{}

func (n NotThrow) firstThrow(throwResult ThrowResult) ThrowState {
	if throwResult == 10 {
		return Strike{}
	}
	return ThrowFirst{int(throwResult)}
}

func (s Strike) ToString() string {
	return "X"
}
func (s NotThrow) ToString() string {
	return "-"
}
func (s ThrowFirst) ToString() string {
	return fmt.Sprint(s.score)
}
func (s ThrowSecond) ToString() string {
	return fmt.Sprintf("%d %d", s.first, s.second)
}
func (s Spare) ToString() string {
	return "/"
}
