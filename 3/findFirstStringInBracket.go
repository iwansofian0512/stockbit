package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	res := findFirstStringInBracket("(asd)")
	res2 := findFirstStringInBracket2("(asd)(123)")
	fmt.Println(res, res2)
}

func findFirstStringInBracket(str string) string {
	res := ""
	if len(str) > 0 {
		indexFirstBracketFound := strings.Index(str, "(")
		if indexFirstBracketFound >= 0 {
			runes := []rune(str)
			wordsAfterFirstBracket := string(runes[indexFirstBracketFound:len(str)])

			indexClosingBracketFound := strings.Index(wordsAfterFirstBracket, ")")
			if indexClosingBracketFound >= 0 {
				runes := []rune(wordsAfterFirstBracket)
				return string(runes[1:indexClosingBracketFound])
			}
		}
	}

	return res
}

//using regex for find string inside the bracket
func findFirstStringInBracket2(str string) string {
	res := ""

	r := regexp.MustCompile(`\((.*?)\)`)
	submatchall := r.FindAllString(str, -1)

	for _, s := range submatchall {
		s = strings.Trim(s, "(")
		s = strings.Trim(s, ")")
		res += s
	}

	return res
}
