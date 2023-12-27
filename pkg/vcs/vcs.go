package vcs

import (
	"log"
	"math/rand"
	"time"
)

type VCS struct {
	environments []string
}

func New() *VCS {
	return &VCS{
		environments: make([]string, 0),
	}
}

func (v *VCS) Start(monitorChan chan string) {
	log.Println("Version control system started")
	ticker := time.NewTicker(time.Second * 5)

	go func(ticker *time.Ticker) {
		defer ticker.Stop()

		for range ticker.C {
			if len(v.environments) > 0 {
				randSource := rand.NewSource(time.Now().UnixNano())
				random := rand.New(randSource)
				randomIndex := random.Intn(len(v.environments))
				randomEnvironment := v.environments[randomIndex]
				monitorChan <- randomEnvironment
			}
		}
	}(ticker)

}

func (v *VCS) AddEnv(env string) {
	if found := v.containsEnv(env); !found {
		v.environments = append(v.environments, env)
	}
}

func (v *VCS) containsEnv(env string) bool {
	for _, val := range v.environments {
		if val == env {
			return true
		}
	}
	return false
}
