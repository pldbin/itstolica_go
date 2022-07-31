package server

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"unicode/utf8"
)

const (
	ac = 1 << iota
	bc
	cc
	dc
	ec
	fc
	nc
	mc
)

type Person struct {
	Name string
	Age  int
	key  string
}

type Home struct {
	Person []Person
	Adress string
}

func (p Person) ShowAge() {
	fmt.Printf("%s's age is %d\n", p.Name, p.Age)
}

func (p *Person) SetAge(age int) {
	p.Age = age
	fmt.Println("Age changed:", p.Age)
}

//1523123
var t = 6

func init() {
	t = 20
}

func Add(nums ...int) (res int, err error) {
	for _, num := range nums {
		res += num // res = res + num
	}
	return res, err
}

func ShowT() int {

	b1, b2 := 34, 43
	// b3, e := Add(b1, b2)
	b3, e := Add(b1, b2, 454, 345, 1, 4)
	fmt.Println(b3, e)
	fmt.Println(t)
	//int8 int16 int32 int64
	//uint8 uint
	// 2 4 8 16 32 64 1288 256 512 2^10=1024~10^3 int32~2*10^9 int64~4*10^18
	//byte = uint8
	var r rune = 0x2654 // 2*16^3+6*16^2+5*16^1+4*16^0   a45f27
	fmt.Printf("%[1]d %[1]x %[1]c\n", r)

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
	// sl5 := []int{4, 6, 3, 8}
	// _ = sl5

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
	sl3 = sl2[4:]
	fmt.Println("sl2", sl2, len(sl2), cap(sl2))
	fmt.Println("sl3", sl3, len(sl3), cap(sl3))
	sl3 = append(sl3, 45, 34, 18, 4)
	sl2 = sl3
	sl3 = sl3[1:4]
	fmt.Println("sl2", sl2, len(sl2), cap(sl2))
	fmt.Println("sl3", sl3, len(sl3), cap(sl3))

	sl3 = append(sl3, 13, 84)
	fmt.Println("sl2", sl2, len(sl2), cap(sl2))
	fmt.Println("sl3", sl3, len(sl3), cap(sl3))

	sl3 = append(sl3, 13, 84)
	fmt.Println("sl2", sl2, len(sl2), cap(sl2))
	fmt.Println("sl3", sl3, len(sl3), cap(sl3))

	// fmt.Print("\n \t \" \\")

	var Andrew = Person{Name: "Andrew", Age: 25, key: "as"}
	fmt.Println(Andrew.key)
	// var Andrew = Person{"Andrew", 25,"adsd"}
	var t2 = Home{make([]Person, 0, 1), "Kyiv"}
	t2.Person = append(t2.Person, Andrew)
	fmt.Println(Andrew.Name, Andrew.Age)
	fmt.Printf("%v\n", t2)
	Andrew.ShowAge()
	Andrew.SetAge(27)
	Andrew.ShowAge()

	Andrew.Age = 15
	t2.Person = append(t2.Person, Andrew)
	var num2 = 5
	var pointer = &num2
	var pointer2 = &pointer
	fmt.Println(num2, pointer, pointer2)

	zero(&num2)
	fmt.Println(num2)

	if 5 > 13 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

	for i := 0; i < 50; i++ {
		fmt.Printf("%d\t", i)
	}
	fmt.Println()

	for {
		fmt.Println("num2 =", num2)
		if num2 > 6 {
			break
		}
		num2++
	}

	fmt.Println(t2.Person)
	for _, value := range t2.Person {
		fmt.Println(value.Age)
	}

	fruits := []string{"banana", "apple", "coconut"}

Loop:
	for _, val := range fruits {
		switch val {
		case "banana": // if val=="banana"
			fmt.Println("Good")
			fallthrough
		case "coconut":
			fmt.Println("Great")
			break Loop
		default:
			fmt.Println("I dont like it.")
		}
	}

	fmt.Println()
	return sl2[2]
}

func Math() {
	var a1, a2 = 12, 16
	// a1++//a+=1 a=a+1
	// fmt.Println(a1, a2, a1+a2, a1-a2, a1*a2, a1/a2, a1%a2)
	// var a1, a2 float64 = 8, 3
	// fmt.Println(a1, a2, a1+a2, a1-a2, a1*a2/a2, a1/a2*a2)
	// var num float64 = -1.7e+308
	// fmt.Println(num, num/32*10*3.2)
	fmt.Println(NSD(a1, a2))
	fmt.Println(NSD(32, 5))
	fmt.Println(NSD(6, 32))

}

func WorkFile() {
	//  f, err := os.Open("/tmp/dat")
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {

		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func zero(a *int) {
	*a = 0
}

func NSD(a, b int) int {
	if a%b == 0 {
		return b
	} else {
		return NSD(b%a, a)
	}
	// return NSD(a, b)
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

type Salary struct {
	Basic, HRA, TA float64
}

type Employee struct {
	FirstName, LastName, Email string
	Age                        int
	MonthlySalary              []Salary
}

func ExampleJson() {
	data := Employee{
		FirstName: "Mark",
		LastName:  "Jones",
		Email:     "mark@gmail.com",
		Age:       25,
		MonthlySalary: []Salary{
			Salary{
				Basic: 15000.00,
				HRA:   5000.00,
				TA:    2000.00,
			},
			Salary{
				Basic: 16000.00,
				HRA:   5000.00,
				TA:    2100.00,
			},
			Salary{
				Basic: 17000.00,
				HRA:   5000.00,
				TA:    2200.00,
			},
		},
	}

	file, _ := json.MarshalIndent(data, "", " ")

	_ = ioutil.WriteFile("test.json", file, 0644)
}

func ExampleMap() {
	m := make(map[string]int)
	m["andrew"] = 2
	fmt.Println(m["andrew"])
	m2 := make(map[string]map[int]int)
	m2["a"][5] = 5
}
