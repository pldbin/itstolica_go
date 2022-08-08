package tasks

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func Compare(a, b int) (res bool) {
	res = a == b
	// if a == b {
	// 	return true
	// } else {
	// 	return false
	// }
	return
}

//	func More(a, b int) int {
//		if a > b {
//			return a
//		} else {
//			return b
//		}
//	}
func Game() {
	// fmt.Println()
	fmt.Println("Start Game, enter a number 0-7 and press enter:")
	var d int
	fmt.Scanf("%d\n", &d)
	r := rand.Intn(d)

	var num int
	// _ = r
	fmt.Println("input number 0-7: ")
	for {
		_, err := fmt.Scanf("%d\n", &num)
		if err != nil {
			continue
		}
		if num == r {
			fmt.Println("Correct number ")
			break
		} else {
			fmt.Println("Wrong number ")
			fmt.Println("input number 0-7: ")
		}
	}
	fmt.Println("the end of this game")
}

// CorrectNum return one number from input
func CorrectNum() (num int) {
	return
}
