package vcs

import "log"

type VCS struct {
	environment []string
}

func New() *VCS {
	return &VCS{
		environment: make([]string, 0),
	}
}

func (v *VCS) Start() {
	log.Println("Version control system started")
}
