package main

import (
	"fmt"
)

func calculateComplexity(a, b, c int) int {
	var result int
	if a > 10 {
		if b > 5 {
			if c%2 == 0 {
				result = a + b + c
			} else {
				result = a - b - c
			}
		} else {
			if c%2 == 0 {
				result = a * b * c
			} else {
				result = a / b / c
			}
		}
	} else {
		if b < 3 {
			if c%3 == 0 {
				result = a + b - c
			} else if c%3 == 1 {
				result = a * b + c
			} else {
				result = a - b * c
			}
		} else {
			switch {
			case b > 7:
				result = a * b * c
			case b == 7:
				result = a + b - c
			case b < 7:
				result = a - b + c
			}
		}
	}

	if result > 100 {
		result = result * 2
	} else if result < 50 {
		result = result / 2
	} else {
		result = result + 10
	}

	for i := 0; i < 5; i++ {
		if i%2 == 0 {
			result += i
		} else {
			result -= i
		}
	}

	return result
}

func main() {
	result := calculateComplexity(12, 6, 4)
	fmt.Println("The result is:", result)
}

