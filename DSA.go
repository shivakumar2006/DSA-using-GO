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

// package main

// import "fmt"

// func main() {
// 	array := []int{1, 2, 3, 4, 5}
// 	result := leftRotate(array, 2)
// 	fmt.Println(result)
// }

// func leftRotate(array []int, d int) []int {
// 	n := len(array)
// 	reverse(array, 0, d-1)
// 	reverse(array, d, n-1)
// 	reverse(array, 0, n-1)
// 	return array
// }

// func reverse(array []int, i int, j int) {
// 	for i < j {
// 		temp := array[i]
// 		array[i] = array[j]
// 		array[j] = temp
// 		i++
// 		j--
// 	}
// }
// _______________________________________________________________________________________________________________________________

// Given a sorted and rotated array A of N distincts elements which is rotated at some point, and given an element key. The task is to find the index of given element key in the array A

// package main

// import "fmt"

// func main() {
// 	array := []int{5, 6, 7, 8, 9, 10, 1, 2, 3}
// 	key := 10
// 	fmt.Println(search(array, 0, len(array)-1, key))
// }

// func search(array []int, l int, r int, key int) int {
// 	pivot := getPivot(array, l, r)
// 	if pivot == -1 {
// 		return binarySearch(array, l, r, key)
// 	} else if array[pivot] == key {
// 		return pivot
// 	} else if array[l] <= key {
// 		return binarySearch(array, l, pivot-1, key)
// 	}
// 	return binarySearch(array, pivot+1, r, key)
// }

// func getPivot(array []int, l int, r int) int {
// 	for l <= r {
// 		mid := (l + r) / 2
// 		if mid < r && array[mid] > array[mid+1] {
// 			return mid
// 		}
// 		if mid > l && array[mid] < array[mid-1] {
// 			return mid - 1
// 		}
// 		if array[l] >= array[mid] {
// 			r = mid - 1
// 		} else {
// 			l = mid + 1
// 		}
// 	}
// 	return -1
// }

// func binarySearch(array []int, l int, r int, x int) int {
// 	for l <= r {
// 		mid := (l + r) / 2
// 		if array[mid] == x {
// 			return mid
// 		} else if array[mid] < x {
// 			l = mid + 1
// 		} else {
// 			r = mid - 1
// 		}
// 	}
// 	return -1
// }
//_______________________________________________________________________________________________________________________________

// Convert any Decimal number n into Binanry number.

// package main

// import (
// 	"fmt"
// 	"strconv"
// )

// func main() {
// 	fmt.Println(decimalToBinary(17))
// }

// func decimalToBinary(n int) string {
// 	if n == 0 {
// 		return "0"
// 	}

// 	var result string
// 	for n > 0 {
// 		remainder := n % 2
// 		result = strconv.Itoa(remainder) + result
// 		n = n / 2
// 	}
// 	return result
// }
// _______________________________________________________________________________________________________________________________

// Convert any Binary Number into Decimal numbers...>

// package main

// import (
// 	"fmt"
// 	"strconv"
// )

// func main() {
// 	binaryStr := "10011"
// 	decimalValue, err := strconv.ParseInt(binaryStr, 2, 0)
// 	if err != nil {
// 		fmt.Println("Error : ", err)
// 		return
// 	}
// 	fmt.Println(decimalValue)
// }

//_______________________________________________________________________________________________________________________________

// package main

// import "fmt"

// func main() {
// 	fmt.Println(5 >> 1) // Bitwise Right
// 	fmt.Println(5 << 1) // Bitwise Right
// }
//_______________________________________________________________________________________________________________________________

// package main

// import "fmt"

// func main() {
// 	fmt.Println((322468 & 1) == 1)
// 	fmt.Println((3224689 & 1) == 1)
// }
//_______________________________________________________________________________________________________________________________

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	n := 10
// 	m := 2
// 	fmt.Println(findDifference(n, m))
// }

// func findDifference(n int, m int) int {
// 	x := n / m
// 	number1 := m * x * (x + 1) / 2
// 	number2 := n*(n+1)/2 - number1
// 	return number1 - number2
// }

//_______________________________________________________________________________________________________________________________

// make pairs of an array elements..... -->

// package main

// import "fmt"

// func main() {
// 	number := []int{-1, 1, 2, 3, 1}
// 	target := 2
// 	result := 0
// 	for i := 0; i < len(number); i++ {
// 		for j := 0; j < len(number); j++ {
// 			if number[i]*number[j] < target {
// 				result++
// 			}
// 		}
// 	}
// 	fmt.Println("Result : ", result)
// }
// _______________________________________________________________________________________________________________________________

// package main

// import "fmt"

// func main() {
// 	number := []int{1, 1, 2, 3, 4, 4, 5}
// 	if len(number) == 0 {
// 		fmt.Println("unique count: ", 0)
// 	}

// 	unique := 1
// 	for i := 1; i < len(number); i++ {
// 		if number[i] != number[i-1] {
// 			number[unique] = number[i]
// 			unique++
// 		}
// 	}
// 	fmt.Println("unique count: ", unique)
// 	fmt.Println("updated numbers : ", number[:unique])
// }

// _______________________________________________________________________________________________________________________________

//FIND TH EPAIR OF TWO NUMMBERS WHICH IS EQUAL TO THE TARGET ---->>>>

// package main

// import (
// 	"fmt"
// 	"math"
// )

// func main() {
// 	array := []int{2, 3, 7, 8}
// 	target := 9
// 	i := 0
// 	j := len(array) - 1

// 	for i < j {
// 		sum := array[i] + array[j]
// 		if sum < target {
// 			i++
// 		} else if sum > target {
// 			j--
// 		} else {
// 			fmt.Printf("indices: %d and %d \n", i+1, j+1)
// 			return
// 		}
// 	}
// 	fmt.Println("pair not found :(")
// }

// _______________________________________________________________________________________________________________________________

// package main

// import (
// 	"fmt"
// 	"math"
// )

// func main() {
// 	numbers := []int{-4, -1, 0, 3, 10}
// 	l := 0
// 	r := len(numbers) - 1
// 	result := make([]int, len(numbers))

// 	for i := len(numbers) - 1; i >= 0; i-- {
// 		if math.Abs(float64(numbers[l])) > math.Abs(float64(numbers[r])) {
// 			result[i] = numbers[l] * numbers[l]
// 			l++
// 		} else {
// 			result[i] = numbers[r] * numbers[r]
// 			r--
// 		}
// 	}
// 	fmt.Println("result: ", result)
// }

// Find the square root of an array and sorted them in ascending order.

package main

import (
	"fmt"
	"math"
)

func main() {
	array := []int{-4, -2, 0, 3, 10}
	l := 0
	r := len(array) - 1
	result := make([]int, len(array))

	for i := len(array) - 1; i >= 0; i-- {
		if math.Abs(float64(array[l])) > math.Abs(float64(array[r])) {
			result[i] = array[l] * array[l]
			l++
		} else {
			result[i] = array[r] * array[r]
			r--
		}
	}
	fmt.Println("result of an array : ", result)
}
