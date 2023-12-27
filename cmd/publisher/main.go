package main

import "dbci/pkg/publisher"

func main() {
	publisher := publisher.New()
	publisher.Start()
}
