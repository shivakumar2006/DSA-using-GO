// Print the following pattern for the given N number of rows

// PATTERN 1

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

// PATTERN 2

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

// PATTERN 3

// Pattern for N = 4
// 1234
// 1234
// 1234
// 1234

// package main

// import "fmt"

// func main() {
// 	printPattern(4)
// }

// func printPattern(n int) {
// 	for i := 1; i <= n; i++ {
// 		for j := 1; j <= n; j++ {
// 			fmt.Print(j)
// 		}
// 		fmt.Println()
// 	}
// }

//_______________________________________________________________________________________________________________________________

// PATTERN 4

// Pattern for N = 4
// 1111
// 2222
// 3333
// 4444

// package main

// import "fmt"

// func main() {
// 	printPattern(4)
// }

// func printPattern(n int) {
// 	for i := 1; i <= n; i++ {
// 		for j := 1; j <= n; j++ {
// 			fmt.Print(i)
// 		}
// 		fmt.Println()
// 	}
// }
//_______________________________________________________________________________________________________________________________

// BINARY SEARCH.....

package main

import "fmt"

func main() {
	// var array = [...]int{2, 3, 5, 7, 8, 9, 1}

	fmt.Print(binarySearch(new int[], {2, 3, 5, 7, 8, 9, 1}, 7))

}

func binarySearch(x int, int[] array) {
	l := 1
	r := array.length - 1
	for l <= r {
		mid := (l + r) / 2
		if array[mid] == x {
			return mid
		} else if x < array[mid] {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return -1
}
