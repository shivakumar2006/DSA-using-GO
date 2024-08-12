// package main

// import "fmt"

// func main() {
// 	a := 65
// 	var b int = 45

// 	fmt.Println("a + b = ", a+b)
// }

// package main

// import "fmt"

// func main() {
// 	for i := 1; i <= 10; i++ {
// 		if i%2 == 0 {
// 			fmt.Printf("even numbers %v \n", i)
// 		} else {
// 			fmt.Printf("odd numbers %v \n", i)
// 		}
// 	}
// }

// package main

// import "fmt"

// func main() {
// 	array := []int{2, 3, 4, 5, 8, 9, 55}

// 	result := binarySearch(array, 8)

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
