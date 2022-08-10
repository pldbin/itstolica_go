package tasks

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Game
func Game1() {
	// fmt.Println()
	fmt.Println("Start Game, enter a number of range and press enter:")
	var d int
	d = CorrectNum(0)
	r := rand.Intn(d)

	var num int
	// _ = r
	for i := 0; i < 5; i++ {
		fmt.Printf("input number 0-%d: ", d)
		num = CorrectNum(d)

		// 	_, err := fmt.Scanf("%d\n", &num)
		// 	if err != nil {
		// 		continue
		// 	}
		if num == r {
			fmt.Println("Correct number ")
			break
		} else {
			fmt.Println("Wrong number ")
			// fmt.Println("input number 0-7: ")
		}
	}
	fmt.Println("the end of this game")
}

// 1. CorrectNum
// Обрабатывать входные данные с терминала до тех пор пока не получим положительное число >0

// 2. Внедрить ограниченное количество попыток
// Задается тоже с терминала

// CorrectNum return one number from input
func CorrectNum(num int) (result int) {
	// var e int
	// fmt.Scanf("%d", &e)
	for {
		_, err := fmt.Scanf("%d\n", &result)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if result < 0 {
			fmt.Println("Invalid number(negative number)")
			continue
		}
		if num == 0 && result > 0 {
			fmt.Println("1. Get number:", result)
			break
		}
		if num > 0 && result <= num && result >= 0 {
			fmt.Println("2. Get number:", result)
			break
		}
	}
	return
}

func 