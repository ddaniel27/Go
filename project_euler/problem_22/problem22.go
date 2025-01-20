/**
* Problem 22 - Names scores
* @see {@link https://projecteuler.net/problem=22}
*
* Using names.txt (right click and 'Save Link/Target As...'), a 46K text file
* containing over five-thousand first names, begin by sorting it into alphabetical
* order. Then working out the alphabetical value for each name, multiply this value
* by its alphabetical position in the list to obtain a name score.
*
* For example, when the list is sorted into alphabetical order, COLIN, which is
* worth 3 + 15 + 12 + 9 + 14 = 53, is the 938th name in the list. So, COLIN would
* obtain a score of 938 × 53 = 49714.
*
* What is the total of all the name scores in the file?
*
* @author ddaniel27
 */
package problem22

import (
	"sort"
	"strings"
)

// This func expects a string similar to the contents of names.txt
func Problem22(input string) int {
	parsedInput := strings.Split(strings.ReplaceAll(input, "\"", ""), ",")

	sort.Strings(parsedInput)

	total := 0
	for i, name := range parsedInput {
		total += (i + 1) * nameValue(name)
	}

	return total
}

func nameValue(name string) int {
	total := 0
	for _, c := range name {
		total += int(c) - 64
	}

	return total
}
