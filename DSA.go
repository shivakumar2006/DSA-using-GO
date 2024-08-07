// Print the following pattern for the given N number of rows
// Pattern for N = 4;
// 4444
// 4444
// 4444
// 4444

// package main

// import "fmt"

// func main() {
// 	var n int
// 	fmt.Print("Enter the number of rows : ")
// 	fmt.Scan(&n)
// 	printPattern(n)
// }

// func printPattern(n int) {
// 	for i := 1; i <= n; i++ {
// 		for j := 1; j <= n; j++ {
// 			fmt.Print(n)
// 		}
// 		fmt.Println()
// 	}
// }
//_______________________________________________________________________________________________________________________________

// Pattern for N = 4
// ****
// ****
// ****
// ****

// package main

// import "fmt"

// func main() {
// 	var n int
// 	fmt.Print("Enter the number of rows : ")
// 	fmt.Scan(&n)

// 	printPattern(n)
// }

// func printPattern(n int) {
// 	for i := 1; i <= n; i++ {
// 		for j := 1; j <= n; j++ {
// 			fmt.Print("*")
// 		}
// 		fmt.Println()
// 	}
// }
//_______________________________________________________________________________________________________________________________

// Pattern for N = 4
// 1234
// 1234
// 1234
// 1234

package main

import "fmt"

func main() {
	printPattern(4)
}

func printPattern(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			fmt.Print(j)
		}
		fmt.Println()
	}
}
