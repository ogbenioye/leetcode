package main

func isvalid(s string) bool {
	paren := map[string]string{
		"(": ")",
		"[": "]",
		"{": "}",
	}

	pile := []string{}

	for _, v := range s {

		if _, ok := paren[string(v)]; ok {
			pile = append(pile, string(v))
		} else if len(pile) == 0 || paren[pile[len(pile)-1]] != string(v) {
			return false
		} else {
			pile = pile[:len(pile)-1]
		}
	}

	return len(pile) == 0
}

func main() {
	isvalid("(){}[]")
}
