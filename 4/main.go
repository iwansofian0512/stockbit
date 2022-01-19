package main

import (
	"fmt"
	"sort"
)

func main() {
	res := anagram([]string{"kita", "atik", "tika", "aku", "kia", "makan", "kua"})
	fmt.Println("res", res)
}

func anagram(str []string) [][]string {
	groups := make(map[string][]string)
	for _, s := range str {
		b := []byte(s)

		//sort to asc
		sort.Slice(b, func(i, j int) bool {
			return b[i] < b[j]
		})
		key := string(b)
		// append to group by sorted string b as key
		groups[key] = append(groups[key], s)
	}

	var res [][]string
	for _, v := range groups {
		res = append(res, v)
	}

	return res
}
