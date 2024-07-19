package main

import "fmt"

func Hello() string {
	var name string
	greeting := "Hello "
	ending := ", nice to meet you!"
	fmt.Print("What is your name?: ")
	_, err := fmt.Scanln(&name)
	if err != nil {
		fmt.Println("Couldn't get name", err)
		return ""
	}
	return greeting + name + ending
}

func main() {
	fmt.Println(Hello())
}
