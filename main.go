package main

import (
	"fmt"
	"log"

	"github.com/mafuyuk/graphql-realtime-chat/server"

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

	s, err := server.NewGraphQLServer(conn)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	if err := s.Serve("/graphql",8080); err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println("Lisiening on :8080")

}
