package main

import "fmt"

func romanToInt(s string) int {
	romans := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	value := 0

	runes := []rune(s)
	n := len(runes)

	for i := 0; i < n; i++ {
		currentVal := romans[runes[i]]

		if i < n-1 && currentVal < romans[runes[i+1]] {
			value -= currentVal
		} else {
			value += currentVal
		}
	}

	return value
}

func roman_to_int() {
	fmt.Printf("III = %d\n", romanToInt("III"))
	fmt.Printf("LVIII = %d\n", romanToInt("LVIII"))
	fmt.Printf("MCMXCIV = %d\n", romanToInt("MCMXCIV"))
}
