package main

import (
	"fmt"
	"strings"
)

func main() {
	data := "apple, orange, mango"

	parts := strings.Split(data, ",")
	fmt.Println(parts) // splits contents of the string.

	str1 := "one two three four two two five"
	fmt.Println(strings.Count(str1, "two")) //counts certain substring in a main string.

	str2 := "united states of america"
	fmt.Println(strings.Contains(str2, "of")) // check if a string contains a substring/ characters.

	fmt.Println(strings.Join([]string{str1, str2}, " , ")) // joins 2 strings using a separator.

}
