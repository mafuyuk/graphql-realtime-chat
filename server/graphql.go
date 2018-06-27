//go:generate gqlgen -schema ./schema.graphql
package server

import (
	"context"
	"sync"

	"github.com/gomodule/redigo/redis"
)

type graphQLServer struct {
	redisConn       redis.Conn
	messageChannels map[string]chan Message
	userChannels    map[string]chan string
	mutex           sync.Mutex
}

func NewGraphQLServer(conn redis.Conn) (*graphQLServer, error) {
	return &graphQLServer{
		redisConn: conn,
	}, nil
}

func (s *graphQLServer) Mutation_postMessage(ctx context.Context, user string, text string) (*Message, error) {
	return nil, nil
}

func (s *graphQLServer) Query_messages(ctx context.Context) ([]Message, error) {
	return nil, nil
}

func (s *graphQLServer) Query_users(ctx context.Context) ([]string, error) {
	return nil, nil
}

func (s *graphQLServer) Subscription_messagePosted(ctx context.Context, user string) (<-chan Message, error) {
	return nil, nil
}

func (s *graphQLServer) Subscription_userJoined(ctx context.Context, user string) (<-chan string, error) {
	return nil, nil
}
