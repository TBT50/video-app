package main

import (
	"fmt"
	"strings"
)

func greeting() {
	var startIndex, endIndex int
	var loop string

	fmt.Print("Start index: ")
	fmt.Scan(&startIndex)

	fmt.Print("End index: ")
	fmt.Scan(&endIndex)

	fmt.Println("What loop would you like to use?")
	fmt.Scan(&loop)

	if strings.ToLower(loop) == "for" {
		fmt.Println("You've choosen the FOR loop!")
		for i:=startIndex; i<=endIndex; i++ {
			fmt.Println(i)
		}
	} else {
		fmt.Println("Actually, there is nothing to choose from. You have to use the 'FOR' loop. Try again")
	}
}

func main() {
	fmt.Print("The server is starting...")
	StartServer()
}

