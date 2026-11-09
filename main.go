package main

import "fmt"

func getName() string {
	var name string

	fmt.Println("Welcome to casino")
	fmt.Printf("Enter your name :")
	_, err := fmt.Scanln(&name)
	if err != nil {
		return ""
	}
	fmt.Printf("Welcome %s lets play \n", name)
	return name
}
func getBet(balance uint) uint {
	var bet uint
	for true {
		fmt.Printf("Enter your bet:")
		_, err := fmt.Scan(&bet)
		if err != nil {
			continue
		}
		if bet > balance {
			fmt.Printf("Your bet is greater than your balance")
		} else {
			break
		}
	}
	return bet
}

func main() {
	balance := uint(200)

	getName()

	bet := getBet(balance)
	fmt.Println(bet)

}
