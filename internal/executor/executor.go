package executor

import (
	"encoding/json"
	"fmt"
	"math/rand"

	"github.com/graphql-go/graphql"
)

const (
	passwordLength = 16
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
						"password": &graphql.Field{
							Type: graphql.String,
							Args: graphql.FieldConfigArgument{
								"useNumber": &graphql.ArgumentConfig{Type: graphql.Boolean, DefaultValue: false, Description: "Use Number"},
							},
							Resolve: password,
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

func password(p graphql.ResolveParams) (interface{}, error) {
	useNumber := p.Args["useNumber"].(bool)

	var chars = "abcdefghijklmnopqrstuvwxyz"
	if useNumber {
		chars += "1234567890"
	}

	charsLength := len(chars)

	var ret string

	for i := 0; i < passwordLength; i++ {
		randomIndex := rand.Intn(charsLength)
		ret += chars[randomIndex : randomIndex+1]
	}

	return ret, nil
}
