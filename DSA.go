// Print the following pattern for the given N number of rows 
// Pattern for N = 4;
// 4444
// 4444
// 4444
// 4444

package main

import "fmt"

func main() {
	printPattern(4)
}

func printPattern() int n {
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			fmt.Println(n)
		}
		fmt.Println()
	}
}
