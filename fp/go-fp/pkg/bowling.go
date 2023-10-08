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
			// case end of early round
			// so we need to add final round
			if i == len(rounds.Early)-1 {
				final := rounds.Final
				switch final.(type) {
				case FinalRoundNotThrow:
					result.Final = FinalRoundFirst{value: throwed}
				case FinalRoundFirst:
					result.Final = final.(FinalRoundFirst).toSecond(throwed)
				case FinalRoundSecond:
					if final.(FinalRoundSecond).canThirdThrow() {
						result.Final = final.(FinalRoundSecond).toThird(throwed)
					}
				}
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
	result += r.Final.score()
	return Score(result)
}

func NewDefaultRounds() Rounds {
	early := make([]Round, 9)
	for i := 0; i < 9; i++ {
		early[i] = newRound()
	}
	return Rounds{
		Early: early,
		Final: FinalRoundNotThrow{},
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

type FinalRound interface {
	score() int
}

type FinalRoundNotThrow struct{}

func (f FinalRoundNotThrow) score() int {
	return 0
}

type FinalRoundFirst struct {
	value ThrowResult
}

func (f FinalRoundFirst) score() int {
	return int(f.value)
}
func (f FinalRoundFirst) toSecond(throwedSecond ThrowResult) FinalRoundSecond {
	return FinalRoundSecond{
		first:  f.value,
		second: throwedSecond,
	}
}

type FinalRoundSecond struct {
	first  ThrowResult
	second ThrowResult
}

func (f FinalRoundSecond) score() int {
	return int(f.first) + int(f.second)
}
func (f FinalRoundSecond) canThirdThrow() bool {
	return f.score() >= 10
}
func (f FinalRoundSecond) toThird(throwedThird ThrowResult) FinalRoundThird {
	return FinalRoundThird{
		first:  f.first,
		second: f.second,
		third:  throwedThird,
	}
}

type FinalRoundThird struct {
	first  ThrowResult
	second ThrowResult
	third  ThrowResult
}

func (f FinalRoundThird) score() int {
	return int(f.first) + int(f.second) + int(f.third)
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
		final := r.Final
		switch final.(type) {
		case FinalRoundNotThrow:
			return NotDeterminableScore{}
		case FinalRoundFirst:
			return NotDeterminableScore{}
		case FinalRoundSecond:
			return DeterminableScore{value: 10 + int(final.(FinalRoundSecond).second)}
		case FinalRoundThird:
			return DeterminableScore{value: 10 + int(final.(FinalRoundThird).first) + int(final.(FinalRoundThird).second)}
		default:
			return NotDeterminableScore{}
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
		if index == len(r.Early)-2 {
			final := r.Final
			switch final.(type) {
			case FinalRoundNotThrow:
				return NotDeterminableScore{}
			case FinalRoundFirst:
				return DeterminableScore{value: 20 + final.(FinalRoundFirst).score()}
			case FinalRoundSecond:
				return DeterminableScore{value: 20 + int(final.(FinalRoundSecond).first)}
			case FinalRoundThird:
				return DeterminableScore{value: 20 + int(final.(FinalRoundThird).first)}
			default:
				return NotDeterminableScore{}
			}
		}
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
		switch rounds.Final.(type) {
		case FinalRoundNotThrow:
			return NotDeterminableScore{}
		default:
			final := rounds.Final
			switch final.(type) {
			case FinalRoundFirst:
				return DeterminableScore{10 + int(final.(FinalRoundFirst).score())}
			case FinalRoundSecond:
				return DeterminableScore{10 + int(final.(FinalRoundSecond).first)}
			case FinalRoundThird:
				return DeterminableScore{10 + int(final.(FinalRoundThird).first)}
			default:
				return NotDeterminableScore{}
			}
		}
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
