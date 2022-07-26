package server

import (
	"fmt"
	"unicode/utf8"
)

type Person struct {
	name string
	age  int
}

//1523123
var t = 6

func init() {
	t = 20
}

func Add(a int, b int) (res int, err error) {
	res = a + b
	return res, err
}

func ShowT() int {

	b1, b2 := 34, 43
	b3, e := Add(b1, b2)
	fmt.Println(b3, e)
	fmt.Println(t)
	//int8 int16 int32 int64
	//uint8 uint
	// 2 4 8 16 32 64 1288 256 512 2^10=1024~10^3 int32~2*10^9 int64~4*10^18
	//byte = uint8
	var r rune = 0x2654 // 2*16^3+6*16^2+5*16^1+4*16^0   a45f27
	fmt.Printf("%d %x %c\n", r, r, r)

	msg := "one 🐜"
	n1 := len(msg)
	n2 := utf8.RuneCountInString(msg)

	fmt.Println(msg)
	fmt.Println(n1)
	fmt.Println(n2)
	fmt.Print("132131 ")
	fmt.Print("132131 ")
	fmt.Printf(" \n")

	a1, a2 := true, false

	fmt.Println(a1, a2, a1 && a2, a1 || a2, !a1, !a2)

	k := 124
	fmt.Printf("%[1]b\n %O\n %d %X", k, k, k, k)
	fmt.Printf("\n\n%4d 1\n%9.7d 2\n%.17d 3\n%8.1d 4\n", k, k, k, k)

	a := uint8(255)
	fmt.Println(a)
	var num float64          //0.0
	var str string = "Hello" //""
	var bl bool = true       //false

	var kt = fmt.Sprintf("%t 1234", a1)
	fmt.Println(kt)
	fmt.Println(num, str, bl)

	var m [3]int
	fmt.Println(m[0])

	mas := [...]int{4, 6, 3, 8}

	b := mas
	b[0] = 5
	fmt.Println(mas, b)

	c := &mas
	c[3] = 0
	fmt.Println(mas, c, b)

	sl := mas[1:3]
	fmt.Println(sl)

	sl2 := make([]int, 4, 5)
	sl2[0] = 2
	fmt.Println(sl2, len(sl2), cap(sl2))
	sl3 := sl2
	sl2 = append(sl2, 13)
	fmt.Println(sl2, len(sl2), cap(sl2))
	sl2 = append(sl2, 8)
	fmt.Println("sl2", sl2, len(sl2), cap(sl2))
	fmt.Println("sl3", sl3, len(sl3), cap(sl3))

	sl3 = sl2
	fmt.Println("sl3", sl3, len(sl3), cap(sl3))
	sl2 = sl2[:len(sl2)-1]
	fmt.Println("sl2", sl2, len(sl2), cap(sl2))
	fmt.Println("sl3", sl3, len(sl3), cap(sl3))
	sl2 = append(sl2, 25)
	fmt.Println("sl2", sl2, len(sl2), cap(sl2))
	fmt.Println("sl3", sl3, len(sl3), cap(sl3))

	return sl2[2]
}
