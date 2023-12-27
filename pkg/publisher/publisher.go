package publisher

import "log"

type Publisher struct{}

func New() *Publisher {
	return &Publisher{}
}

func (p *Publisher) Start() {
	log.Println("Publisher started")
}
