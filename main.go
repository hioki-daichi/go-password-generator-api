package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/hioki-daichi/password-generator-api/internal/executor"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	r := chi.NewRouter()
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
		AllowedMethods: []string{"POST", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type"},
	})
	r.Use(c.Handler)

	r.Post("/graphql", handler)

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	err = r.Body.Close()
	if err != nil {
		panic(err)
	}

	e, err := executor.NewExecutor()
	if err != nil {
		panic(err)
	}

	json, err := e.Execute(fmt.Sprintf("%s", body))
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, "%s\n", json)
}
