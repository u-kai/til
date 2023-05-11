package main

import (
	"fmt"
	"math/rand"
	"time"
)

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

type Actor struct {
	repo *DetectedPublicRepository
}

// この条件が増えていくと大変
// Actionのテストも大変になってくる
// if文の条件を状態として定義するとことで，状態に応じた処理を行うことができる
// 条件の判断と，それに応じた処理を分離することで，処理の追加が容易になる

func (a *Actor) Action() {
	if a.isRegularTime() {
		return
	}
	if a.repo.detectedCount == 0 {
		Alert("admin", "public repository detected")
		Alert(a.repo.owner, "public repository detected")
		a.repo.detectedCount++
		return
	}
	if a.repo.detectedCount%20 == 0 {
		Alert("admin", "public repository detected")
		Alert(a.repo.owner, "forbidden public repository")
		Alert(a.repo.owner, "your account will be suspended!")
	}
	a.repo.detectedCount++
	SaveDetectedInfo(a.repo)
}

func (a *Actor) isRegularTime() bool {
	return time.Now().Hour() <= 17
}

func main() {
	ini := DetectedPublicRepository{name: "init", owner: "init", detectedCount: 0}
	initActor := Actor{repo: &ini}
	initActor.Action()
	for i := 0; i < 100; i++ {
		r := CreateRandomDetectedPublicRepository()
		actor := Actor{repo: r}
		actor.Action()
	}
}
