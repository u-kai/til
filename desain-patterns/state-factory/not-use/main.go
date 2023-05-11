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

func fetchPublicOkRepo(r *DetectedPublicRepository) bool {
	return r.detectedCount == -1
}
func detectRepoChange(r *DetectedPublicRepository) *string {
	s := "change"
	return &s
}

type Actor struct {
	repo *DetectedPublicRepository
}

// この条件が増えていくと大変
// 条件が変わったり，追加された場合は変更が大変
// また，テストも大変
// だんだんファットになっていく
func (a *Actor) Action() {
	if time.Now().Weekday() == time.Saturday || time.Now().Weekday() == time.Sunday {
		if a.repo.detectedCount == 0 {
			Alert("admin", "public repository detected")
			Alert(a.repo.owner, "you are not allowed to create public repository")
			Alert(a.repo.owner, "and you are not allowed to work on weekend")
			a.repo.detectedCount++
			return
		}
		if fetchPublicOkRepo(a.repo) {
			change := detectRepoChange(a.repo)
			if change != nil {
				Alert("admin", "public repository has change")
				Alert("admin", "change: "+*change)
				a.repo.detectedCount++
				return
			}
			return
		}

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
