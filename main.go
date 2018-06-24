package main

import (
	"net/http"
	"log"
	"fmt"

	"github.com/mafuyuk/graphql-realtime-chat/server"

	"github.com/vektah/gqlgen/handler"
  "github.com/gomodule/redigo/redis"
  "github.com/kelseyhightower/envconfig"
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

	conn.Do("SET", "user", "tanaka")
	conn.Do("SET", "user", "oota")


	s, err := redis.String(conn.Do("GET", "user"))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("redis res: %#v", s)

	app := &server.MyApp{}
	http.Handle("/", handler.Playground("Todo", "/query"))
	http.Handle("/query", handler.GraphQL(server.MakeExecutableSchema(app)))

	fmt.Println("Lisiening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}