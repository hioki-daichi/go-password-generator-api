package executor

import (
	"encoding/json"
	"fmt"

	"github.com/graphql-go/graphql"
)

// Executor has the information needed to run GraphQL.
type Executor struct {
	schema graphql.Schema
}

// NewExecutor initializes Executor.
func NewExecutor() (*Executor, error) {
	schema, err := graphql.NewSchema(
		graphql.SchemaConfig{
			Query: graphql.NewObject(
				graphql.ObjectConfig{
					Name: "RootQuery",
					Fields: graphql.Fields{
						"hello": &graphql.Field{
							Type: graphql.String,
							Args: graphql.FieldConfigArgument{
								"name": &graphql.ArgumentConfig{Type: graphql.String, Description: "Name"},
							},
							Resolve: func(p graphql.ResolveParams) (interface{}, error) {
								return ("Hello, " + p.Args["name"].(string) + "!"), nil
							},
						},
					},
				},
			),
		},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create new schema, error: %v", err)
	}

	return &Executor{schema: schema}, nil
}

// Execute executes GraphQL.
func (e *Executor) Execute(requestString string) ([]byte, error) {
	params := graphql.Params{Schema: e.schema, RequestString: requestString}

	result := graphql.Do(params)

	if len(result.Errors) > 0 {
		return nil, fmt.Errorf("failed to execute graphql operation, errors: %+v", result.Errors)
	}

	return json.Marshal(result)
}
