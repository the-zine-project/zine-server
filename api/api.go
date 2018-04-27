package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/graphql-go/graphql"
	"net/http"
	"os"
	"zine/catalog"
	"zine/db"
)

type graphqlBody struct {
	Query string
}

var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name:   "Query",
		Fields: catalog.AssetsField,
	})

var schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query: queryType,
	},
)

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})

	if len(result.Errors) > 0 {
		fmt.Printf("wrong result, unexpected errors: %v", result.Errors)
	}
	return result
}

func main() {
	db.InitDB()
	defer db.DBCon.Close()

	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var body graphqlBody
		err := decoder.Decode(&body)
		if err != nil {
			panic(err)
		}
		defer r.Body.Close()
		result := executeQuery(body.Query, schema)
		json.NewEncoder(w).Encode(result)
	})

	http.ListenAndServe(":8080", handlers.LoggingHandler(os.Stdout, http.DefaultServeMux))
}
