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

// この条件が増えていくと大変
// Actionのテストも大変になってくる
// if文の条件を状態として定義するとことで，状態に応じた処理を行うことができる
// 条件の判断と，それに応じた処理を分離することで，処理の追加が容易になる
func (r *DetectedPublicRepository) Action() {
	if r.isRegularTime() {
		return
	}
	if r.detectedCount == 0 {
		Alert("admin", "public repository detected")
		Alert(r.owner, "public repository detected")
		r.detectedCount++
		return
	}
	if r.detectedCount%20 == 0 {
		Alert("admin", "public repository detected")
		Alert(r.owner, "forbidden public repository")
		Alert(r.owner, "your account will be suspended!")
	}
	r.detectedCount++
	SaveDetectedInfo(r)
}

func (r *DetectedPublicRepository) isRegularTime() bool {
	return time.Now().Hour() <= 17
}

func CreateRandomDetectedPublicRepository() *DetectedPublicRepository {
	randInt := rand.Intn(100)
	return &DetectedPublicRepository{
		name:          "test" + fmt.Sprint(randInt),
		owner:         "test" + fmt.Sprint(randInt),
		detectedCount: randInt,
	}
}

func main() {
	ini := DetectedPublicRepository{name: "init", owner: "init", detectedCount: 0}
	ini.Action()
	for i := 0; i < 100; i++ {
		r := CreateRandomDetectedPublicRepository()
		r.Action()
	}
}
