package publisher

import (
	"dbci/helpers"
	"dbci/pkg/vcs"
	"fmt"
	"log"
	"net"
)

type VersionControl interface {
	Start(monitorChan chan string)
}

type Publisher struct {
	VersionControl
	monitorChan chan string
}

func (p *Publisher) monitor() {
	for {
		environment := <-p.monitorChan
		log.Println(environment)
	}
}

func (p *Publisher) Start() {

	addr, err := net.ResolveTCPAddr("tcp", ":8080")
	helpers.FatalOutError("Starting Tcp Server", err)

	listener, err := net.ListenTCP("tcp", addr)
	helpers.FatalOutError("Listening Tcp Server", err)

	log.Println("Publisher Listening on :8080")
	go p.VersionControl.Start(p.monitorChan)
	go p.monitor()

	for {
		connection, err := listener.Accept()
		remoteAddr := connection.RemoteAddr().String()
		msg := fmt.Sprintf("Accepting remote connection %s", remoteAddr)
		helpers.LogError(msg, err)
		log.Println("Remote connection accepted ", remoteAddr)
	}

}

func New() *Publisher {
	return &Publisher{
		VersionControl: vcs.New(),
		monitorChan:    make(chan string),
	}
}
