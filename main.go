package main

import (
	"collectionneur-cli/src/infrastructure/application"
	"log"
)

func main() {
	cliApp, err := application.NewCLIApp()
	if err != nil {
		log.Fatal(err)
	}

	err = cliApp.Run()
	if err != nil {
		log.Fatal(err)
	}
}
