package main

import (
	"fmt"
)

func isLeap(year int) bool {
	return year%400 == 0 || year%4 == 0 && year%100 != 0
}

// func main() {
// 	for i := 2000; i < 2021; i++ {
// 		if isLeap(i) {
// 			fmt.Println(i)
// 		}
// 	}
// }

func main() {
	fmt.Println(isLeap(1530))
}
