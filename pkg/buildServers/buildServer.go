package buildserver

import (
	"bufio"
	"dbci/helpers"
	"fmt"
	"log"
	"net"
)

type BuildServer struct{}

func New() *BuildServer {
	return &BuildServer{}
}

func (b *BuildServer) Start() {

	addr, err := net.ResolveTCPAddr("tcp", ":8080")
	helpers.FatalOutError("Starting Build Server: ", err)

	publisher, err := net.DialTCP("tcp", nil, addr)

	helpers.FatalOutError("Connecting to the publisher: ", err)

	log.Println("build server started at", publisher.LocalAddr().String())

	for {
		publisherMsg, err := bufio.NewReader(publisher).ReadString('\n')

		msg := fmt.Sprintf("Reading remote address %s", publisher.RemoteAddr().String())
		helpers.LogError(msg, err)

		log.Println(publisherMsg)
	}

}
