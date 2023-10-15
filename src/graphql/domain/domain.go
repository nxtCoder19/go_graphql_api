package graphql_db

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/nxtcoder19/go_graphql_api/graph/model"
	"github.com/nxtcoder19/go_graphql_api/package/mongo"
)

type Impl struct {
	db mongo.DBInterface
}

const (
	LinkTable string = "link"
	UserTable string = "user"
	PostTable string = "post"
)

func (i *Impl) CreateLink(ctx context.Context, title string, address string, user model.User) *model.Link {
	//fmt.Println(user.Name)
	id := uuid.New().String()
	record := &model.Link{
		ID:      id,
		Title:   title,
		Address: address,
		User:    &user,
	}
	fmt.Println("ddd", record)
	err := i.db.InsertRecord(ctx, LinkTable, record)
	if err != nil {
		fmt.Println("here is error")
		return nil
	}
	return record
}

func (i *Impl) CreatePost(ctx context.Context, title string, content string) *model.Post {
	id := uuid.New().String()
	record := &model.Post{
		ID:      id,
		Title:   title,
		Content: content,
	}
	err := i.db.InsertRecord(ctx, PostTable, record)
	if err != nil {
		return nil
	}
	return record
}

func (i *Impl) GetLinks(ctx context.Context) []*model.Link {
	links := make([]*model.Link, 0)
	//var links = []*model.Link
	cursor, err := i.db.GetAllRecords(ctx, LinkTable)
	if err != nil {
		return nil
	}
	defer func() {
		if cer := cursor.Close(ctx); cer != nil {
			fmt.Println(cer)
		}
	}()
	if err = cursor.All(context.TODO(), &links); err != nil {
		panic(err)
	}
	return links
}

func (i *Impl) GetPosts(ctx context.Context) []*model.Post {
	posts := make([]*model.Post, 0)
	cursor, err := i.db.GetAllRecords(ctx, PostTable)
	if err != nil {
		return nil
	}
	defer func() {
		if cer := cursor.Close(ctx); cer != nil {
			fmt.Println(cer)
		}
	}()
	if err = cursor.All(context.TODO(), &posts); err != nil {
		panic(err)
	}
	return posts
}

func (i *Impl) GetPost(ctx context.Context, id string) *model.Post {
	post := &model.Post{}
	fmt.Println("kk", id)
	err := i.db.GetByID(ctx, PostTable, id, post)
	if err != nil {
		fmt.Println("here")
		return nil
	}
	fmt.Println(post)
	return post
}

func (i *Impl) Init(ctx context.Context) error {
	err := i.db.CreateCollection(ctx, "user")
	err = i.db.CreateCollection(ctx, LinkTable)
	err = i.db.CreateCollection(ctx, PostTable)
	if err != nil {
		// Ignore error
		return nil
	}
	return nil
}

func NewGraphQl(db mongo.DBInterface) GraphqlDb {
	return &Impl{db: db}
}
