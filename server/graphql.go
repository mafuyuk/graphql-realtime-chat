//go:generate gqlgen -schema ./schema.graphql
package server

import (
	"fmt"
	"net/http"

	"github.com/gomodule/redigo/redis"
	"github.com/vektah/gqlgen/handler"
	"github.com/gorilla/websocket"
	"github.com/rs/cors"
)

type graphQLServer struct {
	redisConn redis.Conn
}

func NewGraphQLServer(conn redis.Conn) (*graphQLServer, error) {
	return &graphQLServer{
		redisConn: conn,
	}, nil
}

func (s *graphQLServer) Serve(route string, port int) error {
	mux := http.NewServeMux()
	mux.Handle(
		route,
		handler.GraphQL(MakeExecutableSchema(s),
			handler.WebsocketUpgrader(websocket.Upgrader{
				CheckOrigin: func(r *http.Request) bool {
					return true
				},
			}),
		),
	)
	mux.Handle("/playground", handler.Playground("GraphQL", route))

	handler := cors.AllowAll().Handler(mux)
	return http.ListenAndServe(fmt.Sprintf(":%d", port), handler)
}