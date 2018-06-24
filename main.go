package main

import (
	"net/http"
	"log"
	"fmt"

	"github.com/mafuyuk/graphql-realtime-chat/server"

	"github.com/vektah/gqlgen/handler"
  "github.com/garyburd/redigo/redis"
  "github.com/kelseyhightower/envconfig"
)

type redisConf struct {
	Addr string `envconfig:"REDIS_ADDR"`
}

func main() {
	// 環境変数の取得
	c, err := getEnvConfig()
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	fmt.Printf("環境変数: %#v", c)

	// Redisへ接続
	redis, err := redisConnection(c.Addr)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	app := &server.MyApp{}
	http.Handle("/", handler.Playground("Todo", "/query"))
	http.Handle("/query", handler.GraphQL(server.MakeExecutableSchema(app)))

	fmt.Println("Lisiening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func getEnvConfig() (*redisConf, error) {
	var conf redisConf
	err := envconfig.Process("", &conf)
	return &conf, err
}

func redisConnection(addr string) (redis.Conn, error) {
	c, err := redis.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}
	return &c, nil
}