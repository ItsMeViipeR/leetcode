package main

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	original := x
	reversed := 0
	for x != 0 {
		reversed = reversed*10 + x%10
		x /= 10
	}

	if original != reversed {
		return false
	}

	return true
}

func palindrome() {
	x := 121
	result := isPalindrome(x)
	println(result)

	x = -121
	result = isPalindrome(x)
	println(result)

	x = 123
	result = isPalindrome(x)
	println(result)

	x = 32123
	result = isPalindrome(x)
	println(result)
}
