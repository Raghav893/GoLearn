package main

import "fmt"

func getName() string {
	var name string

	fmt.Println("Welcome to casino")
	fmt.Printf("Enter your name :")
	var err error
	_, err = fmt.Scanln(&name)
	if err != nil {
		return ""
	}
	fmt.Printf("Welcome %s", "lets play \n", name)
	return name
}

func main() {
	getName()
}
