package main

import (
	"fmt"
	"nbfriends/basic/models"
)

// var name string = "Noobeeid"
// var age = 10

// var (
// 	name = "noobeeid"
// 	age  = 10
// )

var name, age = "noobeeid", 10

func main() {
	// name, age := "noobeeid", 10
	// user := "Reyhan"

	// fmt.Println(name, age, user)
	// text, twiceAge := get("main", age)
	// fmt.Println("return from get", text, "age is", twiceAge)

	students := []string{"noobeeid", "reyhan", "reza", "nurizko", "suhri", "ahmad"}

	f := find(students)
	isExist := f("reyhan")
	isExistB := f("fikri")

	func() {

		fmt.Println(isExist, isExistB)
	}()

	fmt.Println(models.User)

	coba(10, test())
}

func get(text string, age int) (string, int) {
	fmt.Println("Ini function get", text, age*2)
	post(10, 20, 40, 50, 60, 10, 20, 30, 40, 50, 60, 70, 70)

	return text, age * 2
}

func post(data ...int) {
	fmt.Println("ini post")
	var total = 0
	for i := 0; i < len(data); i++ {
		total += data[i]
	}

	fmt.Println(total)

	coba(10, timesTo20)
	coba(20, func(i int) int {
		return i * 3
	})
	coba(30, func(i int) int {
		return i
	})
}

func timesTo20(i int) int {
	return i * 20
}

func coba(num int, cb func(int) int) {
	total := cb(num)
	fmt.Println("total callback", total)
}

func test() func(int) int {
	return func(i int) int {
		fmt.Println(i)
		return i + 20
	}
}

func find(students []string) func(string) bool {
	var data = ""
	return func(s string) bool {
		data = s
		isExist := false
		for i := 0; i < len(students); i++ {
			if students[i] == s {
				isExist = true
				break
			}
		}

		fmt.Println(data)

		return isExist
	}
}
