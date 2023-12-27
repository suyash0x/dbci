package vcs

import (
	"log"
	"time"
)

type VCS struct {
	environment []string
}

func New() *VCS {
	return &VCS{
		environment: make([]string, 0),
	}
}

func (v *VCS) Start(monitorChan chan string) {
	log.Println("Version control system started")
	ticker := time.NewTicker(time.Second * 5)

	go func(ticker *time.Ticker) {
		defer ticker.Stop()
		for range ticker.C {
			monitorChan <- "Env updated"
		}
	}(ticker)

}
