package main

import (
	"fmt"
	"log"
	"os"

	"github.com/hioki-daichi/password-generator-api/internal/executor"
)

func main() {
	e, err := executor.NewExecutor()
	if err != nil {
		log.Fatalf(err.Error())
	}

	requestString := os.Args[1] // e.g. `{ password(useNumber: true) }`

	json, err := e.Execute(requestString)
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Printf("%s\n", json)
}
