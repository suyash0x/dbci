package publisher

import (
	"dbci/helpers"
	"dbci/pkg/vcs"
	"fmt"
	"log"
	"net"
)

type VersionControl interface {
	Start()
}

type Publisher struct {
	VersionControl
}

func New() *Publisher {
	return &Publisher{
		VersionControl: vcs.New(),
	}
}

func (p *Publisher) Start() {

	addr, err := net.ResolveTCPAddr("tcp", ":8080")
	helpers.FatalOutError("Starting Tcp Server", err)

	listener, err := net.ListenTCP("tcp", addr)
	helpers.FatalOutError("Listening Tcp Server", err)

	log.Println("Publisher Listening on :8080")
	p.VersionControl.Start()

	for {
		connection, err := listener.Accept()
		remoteAddr := connection.RemoteAddr().String()
		msg := fmt.Sprintf("Accepting remote connection %s", remoteAddr)
		helpers.LogError(msg, err)
		log.Println("Remote connection accepted ", remoteAddr)
	}

}
