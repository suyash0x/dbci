package publisher

import (
	"bufio"
	"dbci/helpers"
	"dbci/pkg/vcs"
	"fmt"
	"log"
	"net"
)

type VersionControl interface {
	Start(monitorChan chan string)
	AddEnv(env string)
}

type Publisher struct {
	VersionControl
	monitorChan  chan string
	buildServers map[string][]net.Conn
}

func (p *Publisher) monitor() {
	for {
		environment := <-p.monitorChan
		log.Println(environment)
		for _, buildserver := range p.buildServers[environment] {
			w := bufio.NewWriter(buildserver)
			w.WriteString(environment + "\n")
			w.Flush()
		}
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

		env, err := bufio.NewReader(connection).ReadString('\n')
		helpers.LogError("Reading environment", err)

		p.buildServers[env] = append(p.buildServers[env], connection)
		p.VersionControl.AddEnv(env)

		log.Println("Remote connection accepted ", remoteAddr)
	}

}

func New() *Publisher {
	return &Publisher{
		VersionControl: vcs.New(),
		monitorChan:    make(chan string),
		buildServers:   make(map[string][]net.Conn),
	}
}
