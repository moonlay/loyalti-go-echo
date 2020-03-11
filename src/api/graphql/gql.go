package graphQL

import (
	"fmt"
	"github.com/graphql-go/graphql"
)

// ExecuteQuery runs our graphql queries
func ExecuteQuery(query string) *graphql.Result {

	rootQuery := NewRoot()
	var schema, _ = graphql.NewSchema(
		graphql.SchemaConfig{Query: rootQuery.Query},
	)
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})

	// Error check
	if len(result.Errors) > 0 {
		fmt.Printf("Unexpected errors inside ExecuteQuery: %v", result.Errors)
	}

	return result
}