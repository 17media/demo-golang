package main

import "fmt"

// Function with high cyclomatic complexity (4)
func complexFunction(a, b, c int) int {
    // #lizard forgives [SRE-9183]
    if a > 10 {
        if b > 5 {
            return a + b
        } else {
            return a - b
        }
    } else {
        if c < 5 {
            return b * c
        } else {
            return b / c
        }
    }
}

func main() {
    fmt.Println(complexFunction(12, 7, 3))  // Should print 19
    fmt.Println(complexFunction(8, 2, 6))   // Should print 12
}

