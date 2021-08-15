package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	words := []string{"kita", "atik", "tika", "aku", "kia", "makan", "kua"}
	list := make(map[string][]string)

	for _, word := range words {
		key := sortingStr(word)
		list[key] = append(list[key], word)
	}

	for _, wordResult := range list {
		for _, w := range wordResult {
			fmt.Print(w, " ")
		}
		fmt.Println()
	}
}

func sortingStr(k string) string {
	s := strings.Split(k, "")
	sort.Strings(s)

	return strings.Join(s, "")
}
