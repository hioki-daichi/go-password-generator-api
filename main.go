package main

import (
	"fmt"
	"log"

	"github.com/hioki-daichi/password-generator-api/internal/executor"
)

func main() {
	e, err := executor.NewExecutor()
	if err != nil {
		log.Fatalf(err.Error())
	}

	json, err := e.Execute(`{ password }`)
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Printf("%s\n", json)
}
