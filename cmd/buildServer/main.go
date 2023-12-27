package main

import (
	buildserver "dbci/pkg/buildServers"
	"sync"
)

func main() {
	buildSever := buildserver.New()
	var wg sync.WaitGroup

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		buildSever.Start("development")
	}(&wg)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		buildSever.Start("testing")
	}(&wg)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		buildSever.Start("uat")
	}(&wg)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		buildSever.Start("production")
	}(&wg)

	wg.Wait()

}
