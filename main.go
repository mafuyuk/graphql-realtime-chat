package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mafuyuk/graphql-realtime-chat/server"

	"github.com/gomodule/redigo/redis"
	"github.com/kelseyhightower/envconfig"
	"github.com/vektah/gqlgen/handler"
)

type redisConf struct {
	Addr string `envconfig:"REDIS_ADDR"`
}

func main() {
	// 環境変数の取得
	var conf redisConf
	if err := envconfig.Process("", &conf); err != nil {
		log.Fatal(err)
		panic(err)
	}

	fmt.Printf("環境変数: %#v", conf)

	// Redisへ接続
	conn, err := redis.Dial("tcp", conf.Addr)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	defer conn.Close()

	graphQLServer, err := server.NewGraphQLServer(conn)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	http.Handle("/", handler.Playground("Todo", "/query"))
	http.Handle("/query", handler.GraphQL(server.MakeExecutableSchema(graphQLServer)))

	fmt.Println("Lisiening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
