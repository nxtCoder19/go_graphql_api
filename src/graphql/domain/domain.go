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
)

func (i *Impl) CreateLink(ctx context.Context, title string, address string) *model.Link {
	//fmt.Println(user.Name)
	id := uuid.New().String()
	record := &model.Link{
		ID:      id,
		Title:   title,
		Address: address,
		//User:    nil,
	}
	fmt.Println("ddd", record)
	err := i.db.InsertRecord(ctx, LinkTable, record)
	if err != nil {
		fmt.Println("here is error")
		return nil
	}
	return record
}

func (i *Impl) Init(ctx context.Context) error {
	err := i.db.CreateCollection(ctx, "user")
	err = i.db.CreateCollection(ctx, "link")
	if err != nil {
		// Ignore error
		return nil
	}
	return nil
}

func NewGraphQl(db mongo.DBInterface) GraphqlDb {
	return &Impl{db: db}
}
