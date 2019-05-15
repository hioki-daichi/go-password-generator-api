package generator

import (
	"math/rand"

	"github.com/graphql-go/graphql"
)

const (
	passwordLength = 16
)

// Generator has the information needed to generate password.
type Generator struct {
	useNumber bool
}

// NewGenerator initializes Generator.
func NewGenerator(p graphql.ResolveParams) *Generator {
	return &Generator{
		useNumber: p.Args["useNumber"].(bool),
	}
}

// Generate generates a password.
func (g *Generator) Generate() interface{} {
	var chars = "abcdefghijklmnopqrstuvwxyz"
	if g.useNumber {
		chars += "1234567890"
	}

	charsLength := len(chars)

	var ret string

	for i := 0; i < passwordLength; i++ {
		randomIndex := rand.Intn(charsLength)
		ret += chars[randomIndex : randomIndex+1]
	}

	return ret
}
