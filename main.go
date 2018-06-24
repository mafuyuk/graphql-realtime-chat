package main

import (
	"net/http"
	"log"
	"fmt"

	"github.com/mafuyuk/graphql-realtime-chat/server"

	"github.com/vektah/gqlgen/handler"
)

func main() {
	app := &server.MyApp{}
	http.Handle("/", handler.Playground("Todo", "/query"))
	http.Handle("/query", handler.GraphQL(server.MakeExecutableSchema(app)))

	fmt.Println("Lisiening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}