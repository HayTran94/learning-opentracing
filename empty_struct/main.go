package main

import "fmt"

type Person struct {}

type Car struct {}

func main() {
	persons := make(map[struct{}]string)
	a := Person{}
	b := Person{}
	c := Car{}

	persons[a] = "a"
	persons[b] = "b"
	persons[c] = "c"

	fmt.Println(&a)
	fmt.Println(&b)
	fmt.Println(&c)

	fmt.Println(persons[a])
	fmt.Println(persons[b])
	fmt.Println(persons[c])
}
