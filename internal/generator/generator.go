package generator

import (
	"math"
	"math/rand"

	"github.com/graphql-go/graphql"
)

const (
	passwordLength = 16
)

// Generator has the information needed to generate password.
type Generator struct {
	useNumber bool
	useSign   bool
}

// NewGenerator initializes Generator.
func NewGenerator(p graphql.ResolveParams) *Generator {
	return &Generator{
		useNumber: p.Args["useNumber"].(bool),
		useSign:   p.Args["useSign"].(bool),
	}
}

// Generate generates a password.
func (g *Generator) Generate() interface{} {
	lFingers := []finger{lIndex, lMiddle, lRing, lChild}
	rFingers := []finger{rIndex, rMiddle, rRing, rChild}

	lenLFingers := len(lFingers)
	lenRFingers := len(rFingers)

	rand.Shuffle(lenLFingers, func(i, j int) { lFingers[i], lFingers[j] = lFingers[j], lFingers[i] })
	rand.Shuffle(lenRFingers, func(i, j int) { rFingers[i], rFingers[j] = rFingers[j], rFingers[i] })

	var fingers []finger
	for i := 0; i < lenLFingers; i++ {
		fingers = append(fingers, lFingers[i], rFingers[i])
	}

	lenFingers := len(fingers)

	var ret string

	for i := 0; i < passwordLength; i++ {
		f := fingers[int(
			math.Mod(
				float64(i),
				float64(lenFingers),
			),
		)]

		keys := f.keys(g.useNumber, g.useSign)
		idx := rand.Intn(len(keys))
		ret += keys[idx : idx+1]
	}

	return ret
}
