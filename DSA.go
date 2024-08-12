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

// package main

// import "fmt"

// func main() {
// 	array := []int{1, 2, 3, 5, 7, 8, 9}

// 	result := binarySearch(array, 7)

// 	fmt.Println(result)

// }

// func binarySearch(array []int, x int) int {
// 	l := 0
// 	r := len(array) - 1
// 	for l <= r {
// 		mid := (l + r) / 2
// 		if array[mid] == x {
// 			return mid
// 		} else if x < array[mid] {
// 			r = mid - 1
// 		} else {
// 			l = mid + 1
// 		}
// 	}
// 	return -1
// }
// _______________________________________________________________________________________________________________________________

// PATTERN 5
// Given an array of size . The tsk is to left rotate array by D elements where D <= N.

package main

import "fmt"

func main() {
	array := []int{1, 2, 3, 4, 5}
	result := leftRotate(array, 2)
	fmt.Println(result)
}

func leftRotate(array []int, d int) []int {
	n := len(array)
	reverse(array, 0, d-1)
	reverse(array, d, n-1)
	reverse(array, 0, n-1)
	return array
}

func reverse(array []int, i int, j int) {
	for i < j {
		temp := array[i]
		array[i] = array[j]
		array[j] = temp
		i++
		j--
	}
}
