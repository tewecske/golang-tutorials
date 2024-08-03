package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

// someName := "aaa"

func tutorials() {

	fmt.Println("Hello, ninjas!")

	var nameOne string = "mario"
	var nameTwo = "luigi"
	var nameThree string

	// fmt.Println(nameOne, nameTwo, nameThree)

	nameOne = "peach"
	nameThree = "bwoser"
	// fmt.Println(nameOne, nameTwo, nameThree)

	nameFour := "yoshi"
	// fmt.Println(nameFour)
	fmt.Println(nameOne, nameTwo, nameThree, nameFour)

	var ageOne int = 20
	var ageTwo = 30
	ageThree := 40

	fmt.Println(ageOne, ageTwo, ageThree)

	// var numOne int8 = 25
	// var numTwo = -128
	// var numThree uint = 25

	// var scoreOne float32 = 25.98
	// var scoreTwo float64 = 5435435435435435.4
	// scoreThree := 1.5

	age := 12
	name := "Adam"

	fmt.Printf("my age is %v and my name is %v\n", age, name)
	fmt.Printf("my age is %q and my name is %q\n", age, name)
	fmt.Printf("age is of type %T\n", age)
	fmt.Printf("you scored %0.1f\n", 99.99)

	var str = fmt.Sprintf("my age is %v and my name is %v\n", age, name)
	fmt.Println("the saved string is: ", str)

	// var ages [3]int = [3]int{20, 25, 30}
	ages := [3]int{20, 25, 30}
	names := [4]string{"yoshi", "mario", "peach", "bowser"}
	names[1] = "luigi"

	fmt.Println(ages, len(ages))
	fmt.Println(names, len(names))

	// slices

	var scores = []int{100, 50, 60}
	scores[2] = 25
	scores = append(scores, 85)

	fmt.Println(scores, len(scores))

	// slice ranges

	rangeOne := names[1:3]
	rangeTwo := names[2:]
	rangeThree := names[:3]

	fmt.Println(rangeOne)
	fmt.Println(rangeTwo)
	fmt.Println(rangeThree)

	rangeOne = append(rangeOne, "koopa")
	fmt.Println(rangeOne)

	greeting := "hello there friends!"

	fmt.Println(strings.Contains(greeting, "hello"))
	fmt.Println(strings.ReplaceAll(greeting, "hello", "hi"))
	fmt.Println(strings.ToUpper(greeting))
	fmt.Println(strings.Index(greeting, "th"))
	fmt.Println(strings.Split(greeting, " "))

	fmt.Println("original string value = ", greeting)

	ages2 := []int{45, 20, 35, 30, 75, 60, 50, 25}

	sort.Ints(ages2)
	fmt.Println(ages2)

	index := sort.SearchInts(ages2, 30)
	fmt.Println(index)
	fmt.Println(sort.SearchInts(ages2, 90))

	names2 := []string{"yoshi", "mario", "peach", "bowser", "luigi"}

	sort.Strings(names2)
	fmt.Println(names2)

	x := 0

	for x < 5 {
		fmt.Println("value of x is: ", x)
		x++
	}

	for i := 0; i < 5; i++ {
		fmt.Println("value of i is: ", i)
	}

	names3 := []string{"yoshi", "mario", "peach", "bowser", "luigi"}

	for i := 0; i < len(names3); i++ {
		fmt.Println(names3[i])
	}

	for index, val := range names3 {
		fmt.Printf("the position at index %v is %v\n", index, val)
	}

	for _, val := range names3 {
		fmt.Printf("the value is %v\n", val)
		val = "new string"
	}

	fmt.Println(names)

	age2 := 45

	fmt.Println(age2 <= 50)
	fmt.Println(age2 >= 50)
	fmt.Println(age2 == 45)
	fmt.Println(age2 != 50)

	if age2 < 30 {
		fmt.Println("age is less than 30")
	} else if age2 < 40 {
		fmt.Println("age is less than 40")
	} else {
		fmt.Println("age is not less then 40")
	}

	names4 := []string{"yoshi", "mario", "peach", "bowser", "luigi"}

	for index, value := range names4 {
		if index == 1 {
			fmt.Println("continuing at pos", index)
			continue
		}
		if index > 2 {
			fmt.Println("breaking at pos", index)
			break
		}

		fmt.Printf("the value at pos %v is %v \n", index, value)
	}

	sayGreeting("mario")
	sayGreeting("luigi")
	sayBye("mario")

	cycleNames([]string{"cloud", "tifa", "barret"}, sayGreeting)
	cycleNames([]string{"cloud", "tifa", "barret"}, sayBye)

	fmt.Printf("The area of the circle with radius %v is %0.2f\n", 5, circleArea(5))

	first, last := getInitials("tifa lockhart")
	fmt.Println(first, last)

	fmt.Println(getInitials("tifa lockhart"))
	fmt.Println(getInitials("tifa"))

	sayHello("mario")

	for _, v := range points {
		fmt.Println(v)
	}

	showScore()

	menu := map[string]float64{
		"soup":           4.99,
		"pie":            7.99,
		"salad":          6.99,
		"toffee pudding": 3.55,
	}

	fmt.Println(menu)
	fmt.Println(menu["pie"])

	for k, v := range menu {
		fmt.Println(k, "-", v)
	}

	phonebook := map[int]string{
		213412312: "mario",
		325435344: "luigi",
		435435436: "peach",
	}

	fmt.Println(phonebook)
	fmt.Println(phonebook[213412312])

	phonebook[213412312] = "bowser"

	fmt.Println(phonebook)

	// non pointer types: string, ints, bools, floats, arrays, structs

	nameA := "tifa"
	updateName(nameA)
	fmt.Println(nameA)

	// pointer values types: slices, maps, functions

	menuB := map[string]float64{
		"pie":       5.95,
		"ice cream": 3.99,
	}

	updateMenu(menuB)
	fmt.Println(menuB)

	fmt.Println("memory address of name is: ", &nameA)

	m := &nameA
	fmt.Println("memory address: ", m)
	fmt.Println("value at memory address: ", *m)

	fmt.Println(nameA)
	updateNameP(m)
	fmt.Println(nameA)
}

func updateName(x string) {
	x = "wedge"
}

func updateNameP(x *string) {
	*x = "wedge"
}

func updateMenu(y map[string]float64) {
	y["coffee"] = 2.99
}

var score = 99.5

func sayGreeting(n string) {
	fmt.Printf("Good morning %v \n", n)
}

func sayBye(n string) {
	fmt.Printf("Goodbye %v \n", n)
}

func cycleNames(n []string, f func(string)) {
	for _, v := range n {
		f(v)
	}
}

func circleArea(r float64) float64 {
	return math.Pi * r * r
}

func getInitials(n string) (string, string) {
	s := strings.ToUpper(n)
	names := strings.Split(s, " ")
	var initials []string
	for _, v := range names {
		initials = append(initials, v[:1])
	}
	if len(initials) > 1 {
		return initials[0], initials[1]
	}

	return initials[0], "_"
}
