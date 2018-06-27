package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mafuyuk/graphql-realtime-chat/server"

	"github.com/gomodule/redigo/redis"
	"github.com/gorilla/websocket"
	"github.com/kelseyhightower/envconfig"
	"github.com/rs/cors"
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

	s, err := server.NewGraphQLServer(conn)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	mux := http.NewServeMux()
	mux.Handle(
		"/graphql",
		handler.GraphQL(server.MakeExecutableSchema(s),
			handler.WebsocketUpgrader(websocket.Upgrader{
				CheckOrigin: func(r *http.Request) bool {
					return true
				},
			}),
		),
	)
	mux.Handle("/playground", handler.Playground("GraphQL", "/graphql"))

	h := cors.AllowAll().Handler(mux)
	http.ListenAndServe(fmt.Sprintf(":%d", 8080), h)

	fmt.Println("Lisiening on :8080")
}
