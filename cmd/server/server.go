package main

// import (
// 	"log"
// 	"net/http"
// 	"os"

// 	"github.com/99designs/gqlgen/handler"
// 	"github.com/joseteodoro/consent-enforcement/pkg/consent_manager"
// )

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
	"github.com/joseteodoro/consent-enforcement/pkg/consent_manager"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// database := consent_manager.Connect(nil)
	// docs, err := database.QueryJSON(`
	// {
	// 	"selector": {
	// 	   "type": "DataType"
	// 	}
	//  }
	// `)
	// fmt.Println("db access", docs, err)

	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	rootHandler := handler.GraphQL(
		consent_manager.NewExecutableSchema(consent_manager.NewRootResolvers()),
		handler.ComplexityLimit(200))
	http.Handle("/query", rootHandler)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
