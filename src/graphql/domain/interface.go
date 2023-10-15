package graphql_db

import (
	"context"
	"github.com/nxtcoder19/go_graphql_api/graph/model"
)

type Id string

type User struct {
	ID       string `json:"id"`
	Username string `json:"name"`
	Password string `json:"password"`
}

type Link struct {
	ID      string
	Title   string
	Address string
	User    User
}

type GraphqlDb interface {
	Init(ctx context.Context) error
	CreateLink(ctx context.Context, title string, address string, user model.User) *model.Link
	GetLinks(ctx context.Context) []*model.Link
	CreatePost(ctx context.Context, title string, content string) *model.Post
	GetPosts(ctx context.Context) []*model.Post
	GetPost(ctx context.Context, id string) *model.Post
}
