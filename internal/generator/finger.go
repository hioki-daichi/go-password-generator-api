package generator

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

func (f finger) keys(useNumber bool, useSign bool) string {
	var keys string

	switch f {
	case lIndex:
		keys += "tgbrfv"
		if useNumber {
			keys += "54"
		}
	case lMiddle:
		keys += "edc"
		if useNumber {
			keys += "3"
		}
	case lRing:
		keys += "wsx"
		if useNumber {
			keys += "2"
		}
	case lChild:
		keys += "qaz"
		if useNumber {
			keys += "1"
		}
	case rIndex:
		keys += "yhnujm"
		if useNumber {
			keys += "67"
		}
	case rMiddle:
		keys += "ik"
		if useNumber {
			keys += "8"
		}
		if useSign {
			keys += ","
		}
	case rRing:
		keys += "ol"
		if useNumber {
			keys += "9"
		}
		if useSign {
			keys += "."
		}
	case rChild:
		keys += "p"
		if useNumber {
			keys += "0"
		}
		if useSign {
			keys += ";/"
		}
	}

	return keys
}
