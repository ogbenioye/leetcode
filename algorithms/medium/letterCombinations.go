package main

import "fmt"

/*
- Runtime: 1ms, memory usage: 2MB
- Faster than 76.90% of Go Online subbmissions as at 08-08-2023
- max length of digits is 4. Split scenarios into four for faster runtime
- for 4 <= len(digits) >= 2:
	store mapped telephone digits into variables,
	iterate over arrays with nested loops,
	and append combinations into the result array and return.
*/

func letterCombinations(digits string) []string {
	tel := map[string][]string{
		"2": {"a", "b", "c"},
		"3": {"d", "e", "f"},
		"4": {"g", "h", "i"},
		"5": {"j", "k", "l"},
		"6": {"m", "n", "o"},
		"7": {"p", "q", "r", "s"},
		"8": {"t", "u", "v"},
		"9": {"w", "x", "y", "z"},
	}

	var result []string

	if len(digits) == 1 {
		return tel[string(digits[0])]
	}

	if len(digits) == 2 {
		sl1, sl2 := tel[string(digits[0])], tel[string(digits[1])]

		for _, v := range sl1 {
			for _, d := range sl2 {
				result = append(result, v+d)
			}
		}
	}

	if len(digits) == 3 {
		sl1, sl2, sl3 := tel[string(digits[0])], tel[string(digits[1])], tel[string(digits[2])]

		for _, v := range sl1 {
			for _, d := range sl2 {
				for _, n := range sl3 {
					result = append(result, v+d+n)
				}
			}
		}
	}

	if len(digits) == 4 {
		sl1, sl2, sl3, sl4 := tel[string(digits[0])], tel[string(digits[1])], tel[string(digits[2])], tel[string(digits[3])]

		for _, v := range sl1 {
			for _, d := range sl2 {
				for _, n := range sl3 {
					for _, s := range sl4 {
						result = append(result, v+d+n+s)
					}
				}
			}
		}
	}

	return result
}

func main() {
	fmt.Println(letterCombinations("234"))
}
