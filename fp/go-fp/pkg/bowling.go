package pkg

type Turn func(Throw, Bowling) Bowling

type Throw func(Rounds) Rounds

func ThrowByAnyLogic(logic func() ThrowResult) Throw {
	return func(rounds Rounds) Rounds {
		throwed := logic()
		result := NewDefaultRounds()
		for i, round := range rounds.Early {
			if !round.throwedFirst() {
				result.Early[i].FirstThrow = throwed
				break
			}
			result.Early[i].FirstThrow = round.FirstThrow
			if !round.throwedSecond() {
				result.Early[i].SecondThrow = throwed
				break
			}
			result.Early[i].SecondThrow = round.SecondThrow
			// case of last round
			if i == len(rounds.Early)-1 {
				if !rounds.Final.throwedFirst() {
					result.Final.FirstThrow = throwed
					break
				}
				if !rounds.Final.throwedSecond() {
					result.Final.SecondThrow = throwed
					break
				}
				if !rounds.Final.throwedThird() {
					result.Final.ThirdThrow = throwed
					break
				}
			}
		}
		return result
	}
}

type ThrowResult int

const (
	Strike ThrowResult = 10
	Gutter ThrowResult = 0
)

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
	beforeSpare := false
	for _, round := range r.Early {
		if !round.throwedFirst() {
			break
		}
		if beforeSpare {
			result += int(round.FirstThrow)
		}
		result += int(round.FirstThrow)
		if !round.throwedSecond() {
			break
		}
		result += int(round.SecondThrow)
		beforeSpare = round.isSpare()
	}

	if r.Final.throwedFirst() {
		result += int(r.Final.FirstThrow)
	}
	if r.Final.throwedSecond() {
		result += int(r.Final.SecondThrow)
	}
	if r.Final.throwedThird() {
		result += int(r.Final.ThirdThrow)
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
	FirstThrow  ThrowResult
	SecondThrow ThrowResult
}

func (r Round) isStrike() bool {
	return r.FirstThrow == Strike
}
func (r Round) isSpare() bool {
	return r.FirstThrow != Strike && r.FirstThrow+r.SecondThrow == 10
}

func (r Round) throwedFirst() bool {
	return r.FirstThrow != -1
}

func (r Round) throwedSecond() bool {
	return r.SecondThrow != -1
}

func newRound() Round {
	return Round{
		FirstThrow:  -1,
		SecondThrow: -1,
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

type CalcScore func(Rounds) Score
type Score int
