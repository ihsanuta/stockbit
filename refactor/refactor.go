package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Print(findFirstStringInBracket("hall(o)"))
}
func findFirstStringInBracket(str string) string {
	if len(str) > 0 {
		indexFirstBracketFound := strings.Index(str, "(")
		if indexFirstBracketFound == -1 {
			return ""
		}
		indexClosingBracketFound := strings.Index(str, ")")
		if indexClosingBracketFound == -1 {
			return ""
		}

		characterInBracket := indexFirstBracketFound + 1
		if characterInBracket >= indexClosingBracketFound {
			return ""
		}

		wordsInBracket := str[characterInBracket:indexClosingBracketFound]
		if len(wordsInBracket) < 1 {
			return ""
		}

		return wordsInBracket[0:1]
	} else {
		return ""
	}
}
