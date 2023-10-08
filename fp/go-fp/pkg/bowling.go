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
	for i, round := range r.Early {
		switch round.throw.(type) {
		case Strike:
			score := round.throw.(Strike).score(r, i)
			switch score.(type) {
			case DeterminableScore:
				result += score.(DeterminableScore).score()
			case NotDeterminableScore:
				break loop
			}
		case Spare:
			score := round.throw.(Spare).score(r, i)
			switch score.(type) {
			case DeterminableScore:
				result += score.(DeterminableScore).score()
			case NotDeterminableScore:
				break loop
			}
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
type RoundScore interface {
	score() int
}
type DeterminableScore struct {
	value int
}

func triple() DeterminableScore {
	return DeterminableScore{value: 30}
}
func double(nextThrowState ThrowState) RoundScore {
	switch nextThrowState.(type) {
	case ThrowFirst:
		return DeterminableScore{value: 20 + nextThrowState.(ThrowFirst).score}
	case ThrowSecond:
		return DeterminableScore{value: 20 + nextThrowState.(ThrowSecond).first}
	case Spare:
		return DeterminableScore{value: 20 + nextThrowState.(Spare).first}
	case Strike:
		return triple()
	default:
		return NotDeterminableScore{}
	}
}

func (d DeterminableScore) score() int {
	return d.value
}

type NotDeterminableScore struct{}

func (n NotDeterminableScore) score() int {
	return 0
}

type ThrowState interface {
	ToString() string
}

type Strike struct {
}

func (s Strike) score(r Rounds, index int) RoundScore {
	if index == len(r.Early)-1 {
		return DeterminableScore{
			value: 10 + int(r.Final.FirstThrow) + int(r.Final.SecondThrow),
		}
	}
	nextThrow := r.Early[index+1].throw
	switch nextThrow.(type) {
	case NotThrow:
		// can not determined
		return NotDeterminableScore{}
	case ThrowFirst:
		// can not determined
		return NotDeterminableScore{}
		// case Double or Triple
	case ThrowSecond:
		return DeterminableScore{value: 10 + nextThrow.(ThrowSecond).score()}
	case Strike:
		// TODO index
		return double(r.Early[index+2].throw)
	case Spare:
		return DeterminableScore{value: 20}
	default:
		panic(fmt.Sprintf("unknown type %T", nextThrow))
	}
}

type ThrowFirst struct {
	score int
}

func (t ThrowFirst) secondThrow(throwResult ThrowResult) ThrowState {
	if t.score+int(throwResult) == 10 {
		return Spare{
			first:  t.score,
			second: int(throwResult),
		}
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

type Spare struct {
	first  int
	second int
}

func (s Spare) score(rounds Rounds, i int) RoundScore {
	if i == len(rounds.Early)-1 {
		return DeterminableScore{10 + int(rounds.Final.FirstThrow)}
	}
	switch rounds.Early[i+1].throw.(type) {
	case Strike:
		return DeterminableScore{value: 20}
	case ThrowFirst:
		return DeterminableScore{value: 10 + rounds.Early[i+1].throw.(ThrowFirst).score}
	case ThrowSecond:
		return DeterminableScore{value: 10 + rounds.Early[i+1].throw.(ThrowSecond).first}
	case Spare:
		return DeterminableScore{value: 10 + rounds.Early[i+1].throw.(Spare).first}
	case NotThrow:
		return NotDeterminableScore{}
	default:
		return NotDeterminableScore{}
	}
}

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
