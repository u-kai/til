package main

import (
	"fmt"
	"math/rand"
	"time"
)

////////////////////////////////////////////////////////////////////////////////////////////////////
// パターンを利用しないケースと同じ

func Alert(to string, msg string) {
	println("alert to", to, ":", msg)
}
func SaveDetectedInfo(r *DetectedPublicRepository) {
	println("save detected info:", r.name, r.owner, r.detectedCount)
}

type DetectedPublicRepository struct {
	name          string
	owner         string
	detectedCount int
}

func CreateRandomDetectedPublicRepository() *DetectedPublicRepository {
	randInt := rand.Intn(100)
	return &DetectedPublicRepository{
		name:          "test" + fmt.Sprint(randInt),
		owner:         "test" + fmt.Sprint(randInt),
		detectedCount: randInt,
	}
}

// //////////////////////////////////////////////////////////////////////////////////////////////////
// パターンを利用したケース
// 具体的な構造体から抽象的なインターフェースを定義を変更する
type Actor interface {
	Action()
}

// 以下のxxxActorはActorインターフェースを実装している
// 具体的な構造体は状態を加味した構造体のため，その状態の時にやるべきことだけを定義すれば良い
// Actionだけに集中できる
// その状態かどうかの判断は，別に任せる(今回だとfactory)
type RegalTimeActor struct {
}
type InitDetectedActor struct {
	repo *DetectedPublicRepository
}
type OnlyIncrementActor struct {
	repo *DetectedPublicRepository
}
type SuspendedActor struct {
	repo *DetectedPublicRepository
}

func (a *RegalTimeActor) Action() {
}
func (a *OnlyIncrementActor) Action() {
	a.repo.detectedCount++
	SaveDetectedInfo(a.repo)
}
func (a *InitDetectedActor) Action() {
	Alert("admin", "public repository detected")
	Alert(a.repo.owner, "public repository detected")
	a.repo.detectedCount++
	SaveDetectedInfo(a.repo)
}
func (a *SuspendedActor) Action() {
	Alert("admin", "public repository detected")
	Alert(a.repo.owner, "forbidden public repository")
	Alert(a.repo.owner, "your account will be suspended!")
	a.repo.detectedCount++
	SaveDetectedInfo(a.repo)
}

// factory pattern
// 生成という責務を全うする(Actor自体は自身の生成方法を気にせずにActionだけに集中できる)
// actorを使う側は何も考えずにこの関数を呼び出せば良い
// actorが増えた時に変更が必要だが，actorを使う側は変更の必要がない
// また，変更がないactorには影響がない
// また，actorのテストも容易になる
// さらにactorの生成方法を変更したい場合も，この関数のみを変更すれば良い
func createActor(r *DetectedPublicRepository) Actor {
	if time.Now().Hour() <= 17 {
		return &RegalTimeActor{}
	}
	if r.detectedCount == 0 {
		return &InitDetectedActor{repo: r}
	}
	if r.detectedCount%20 == 0 {
		return &SuspendedActor{repo: r}
	}
	return &OnlyIncrementActor{repo: r}
}

func main() {
	ini := DetectedPublicRepository{name: "init", owner: "init", detectedCount: 0}
	initActor := createActor(&ini)
	initActor.Action()
	for i := 0; i < 100; i++ {
		r := CreateRandomDetectedPublicRepository()
		actor := createActor(r)
		actor.Action()
	}
}
