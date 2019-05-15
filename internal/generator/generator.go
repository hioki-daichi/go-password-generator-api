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

		ret += g.randomKey(f)
	}

	return ret
}

type finger int

const (
	lIndex finger = iota
	lMiddle
	lRing
	lChild
	rIndex
	rMiddle
	rRing
	rChild
)

func (g *Generator) randomKey(f finger) string {
	var keys string

	switch f {
	case lIndex:
		keys += "tgbrfv"
		if g.useNumber {
			keys += "54"
		}
	case lMiddle:
		keys += "edc"
		if g.useNumber {
			keys += "3"
		}
	case lRing:
		keys += "wsx"
		if g.useNumber {
			keys += "2"
		}
	case lChild:
		keys += "qaz"
		if g.useNumber {
			keys += "1"
		}
	case rIndex:
		keys += "yhnujm"
		if g.useNumber {
			keys += "67"
		}
	case rMiddle:
		keys += "ik"
		if g.useNumber {
			keys += "8"
		}
		if g.useSign {
			keys += ","
		}
	case rRing:
		keys += "ol"
		if g.useNumber {
			keys += "9"
		}
		if g.useSign {
			keys += "."
		}
	case rChild:
		keys += "p"
		if g.useNumber {
			keys += "0"
		}
		if g.useSign {
			keys += ";/"
		}
	}

	randomIndex := rand.Intn(len(keys))
	return keys[randomIndex : randomIndex+1]
}
