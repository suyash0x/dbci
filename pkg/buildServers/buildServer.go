package buildserver

import (
	"bufio"
	"dbci/helpers"
	"fmt"
	"log"
	"net"
	"time"
)

type BuildServer struct {
	signal chan string
}

func New() *BuildServer {
	return &BuildServer{
		signal: make(chan string),
	}
}

func (b *BuildServer) triggerProcess() {
	for {
		env := <-b.signal
		log.Println("Pulling repo from", env)
		time.Sleep(time.Second * 2)

		log.Println("Building")
		time.Sleep(time.Second * 2)

		log.Println("Deploying")
		time.Sleep(time.Second * 2)

		log.Println("Deployed")
	}
}

func (b *BuildServer) Start(env string) {

	addr, err := net.ResolveTCPAddr("tcp", ":8080")
	helpers.FatalOutError("Starting Build Server: ", err)

	publisher, err := net.DialTCP("tcp", nil, addr)

	helpers.FatalOutError("Connecting to the publisher: ", err)

	log.Println("build server started at", publisher.LocalAddr().String())

	w := bufio.NewWriter(publisher)
	w.WriteString(env + "\n")
	w.Flush()

	go b.triggerProcess()

	for {
		publisherMsg, err := bufio.NewReader(publisher).ReadString('\n')

		msg := fmt.Sprintf("Reading remote address %s", publisher.RemoteAddr().String())
		helpers.LogError(msg, err)

		b.signal <- publisherMsg

	}

}
