package main

import (
	"github.com/gari8/gqlgen-gorm-tutorial/graph/resolver"
	"github.com/gari8/gqlgen-gorm-tutorial/infra"
	"github.com/gari8/gqlgen-gorm-tutorial/model/repository"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gari8/gqlgen-gorm-tutorial/graph/generated"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	conn, _ := infra.NewDB()
	tx := repository.NewTx(conn)
	todoRepo := repository.NewTodoRepository(conn)
	userRepo := repository.NewUserRepo(conn)
	r := resolver.NewResolver(tx, todoRepo, userRepo)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: r}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
