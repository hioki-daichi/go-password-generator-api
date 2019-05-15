package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/hioki-daichi/password-generator-api/internal/executor"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

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
