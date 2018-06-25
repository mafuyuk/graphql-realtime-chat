//go:generate gqlgen -schema ./schema.graphql
package server

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"

	"github.com/gomodule/redigo/redis"
)

type graphQLServer struct {
	redisConn redis.Conn
}

func NewGraphQLServer(conn redis.Conn) (*graphQLServer, error) {
	return &graphQLServer{
		redisConn: conn,
	}, nil
}

func (g *graphQLServer) Query_todos(ctx context.Context) ([]Todo, error) {
	var todos []Todo
	values, err := redis.ByteSlices(g.redisConn.Do("LRANGE", "todo", 0, -1))
	if err != nil {
		return todos, err
	}

	for _, v := range values {
		var todo Todo
		json.Unmarshal(v, &todo)
		todos = append(todos, todo)
	}

	return todos, nil
}

func (g *graphQLServer) Mutation_createTodo(ctx context.Context, text string) (Todo, error) {
	todo := Todo{
		Text: text,
		ID:   fmt.Sprintf("T%d", rand.Int()),
		User: User{
			ID: fmt.Sprintf("U%d", rand.Int()),
		},
	}

	mj, _ := json.Marshal(todo)

	g.redisConn.Do("LPUSH", "todo", mj)
	return todo, nil
}

func (g *graphQLServer) Todo_user(ctx context.Context, it *Todo) (User, error) {
	return User{ID: it.User.ID, Name: "user " + it.User.ID}, nil
}
