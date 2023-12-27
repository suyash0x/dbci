package main

import buildserver "dbci/pkg/buildServers"

func main() {
	buildSever := buildserver.New()
	buildSever.Start()
}
