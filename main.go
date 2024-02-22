package main

import (
	"fmt"
	"os"
	"sync"

	"github.com/devsaud247/utterancesproject/server"
	"github.com/devsaud247/utterancesproject/utils"
)

const DirToRead string = "./utterances"

func main() {

	uChan := make(chan utils.Utterance)
	var wg sync.WaitGroup

	fileNames, err := os.ReadDir(DirToRead)
	if err != nil {
		fmt.Print(err)
		return
	}
	wg.Add(2)
	go utils.ReadUtterances(DirToRead, fileNames, uChan, &wg)

	go utils.GenerateUtterances(uChan, &wg)
	wg.Wait()

	if err := utils.CreateCompleteUtterances("complete.json"); err != nil {
		panic(err)
	}
	// simply running the server
	server.RunServer()
}
