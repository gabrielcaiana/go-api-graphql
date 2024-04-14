package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gabrielcaiana/api-go-graphql/graph"
	"github.com/gabrielcaiana/api-go-graphql/internal/database"
	_ "github.com/mattn/go-sqlite3"
)

const defaultPort = "8080"

func main() {
	// connect to the database
	db, err := sql.Open("sqlite3", "./data.db")

	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	defer db.Close()

	// create the category database connection
	CategoryDb := database.NewCategory(db)

	// create the course database connection
	CourseDb := database.NewCourse(db)

	// port configuration
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// create the server
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{

		// inject the category and course database into the resolver
		CategoryDb: CategoryDb,
		CourseDb:   CourseDb,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
