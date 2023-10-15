package main

import (
	"context"
	"github.com/nxtcoder19/go_graphql_api/package/mongo"
	graphql_db "github.com/nxtcoder19/go_graphql_api/src/graphql/domain"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/nxtcoder19/go_graphql_api/graph"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	mongoUrl := os.Getenv("MONGO_URI")
	db := mongo.NewDB("xyz", mongoUrl)
	err := db.ConnectDB(context.TODO())
	if err != nil {
		panic(err)
	}

	graphql := graphql_db.NewGraphQl(db)
	err = graphql.Init(context.TODO())
	if err != nil {
		panic(err)
	}
	//graphql.CreateLink(context.TODO(), "ppp", "qqq")

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
