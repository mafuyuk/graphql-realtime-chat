//go:generate gqlgen -schema ./schema.graphql
package server

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/gomodule/redigo/redis"
)

type graphQLServer struct {
	redisConn redis.Conn
}

func NewGraphQLServer(conn redis.Conn) (*graphQLServer, error) {
	return &graphQLServer{
		redisConn:     conn,
	}, nil
}

func (g *graphQLServer) Query_todos(ctx context.Context) ([]Todo, error) {
	var todos []Todo
	v, err := redis.Values(g.redisConn.Do("HGETALL", "todo"))
	if err != nil {
		return todos, err
	}
	fmt.Printf("%#v", v) //todo

	if err := redis.ScanStruct(v, &todos); err != nil {
		return todos, err
	}
	return todos, nil
}

func (g *graphQLServer) Mutation_createTodo(ctx context.Context, text string) (Todo, error) {
	todo := Todo{
		Text:   text,
		ID:     fmt.Sprintf("T%d", rand.Int()),
		User: User{
			ID: fmt.Sprintf("U%d", rand.Int()),
		},
	}

	g.redisConn.Do("HSET", "todo", "id", rand.Int())
	g.redisConn.Do("HSET", "todo", "text", text)
	return todo, nil
}

func (g *graphQLServer) Todo_user(ctx context.Context, it *Todo) (User, error) {
	return User{ID: it.User.ID, Name: "user " + it.User.ID}, nil
}
