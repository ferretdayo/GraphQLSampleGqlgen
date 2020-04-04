package main

import (
	"log"
	"net/http"
	"os"

	"github.com/GraphQLSampleGqlgen/src/usecases/users"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/GraphQLSampleGqlgen/src/db"
	"github.com/GraphQLSampleGqlgen/src/generated"
	"github.com/GraphQLSampleGqlgen/src/interfaces/repositories"
	"github.com/GraphQLSampleGqlgen/src/resolver"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	mysql := db.NewMysql()
	db := mysql.Open()

	srv := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{
				Resolvers: &resolver.Resolver{
					UserUsecase: &users.UserUsecase{
						UserRepository: &repositories.UserRepository{},
						DB:             db,
					},
					UserTodoUsecase: &users.UserTodoUsecase{
						UserTodoRepository: &repositories.UserTodoRepository{},
						DB:                 db,
					},
				},
			},
		),
	)

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
