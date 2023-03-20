package main

func romanToInt(s string) int {
	dataBase := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	var lSum, totalSum int

	for i := len(s) - 1; i >= 0; i-- {

		if dataBase[s[i]] < lSum {
			totalSum -= dataBase[s[i]]
		} else {
			totalSum += dataBase[s[i]]
		}
		lSum = dataBase[s[i]]
	}

	return totalSum
}

func main() {
	romanToInt("LVIII")
}
