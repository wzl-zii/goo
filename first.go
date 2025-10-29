package main

import "fmt"

func main() {
	var num = 5
	for i := 1; i <= num; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println("")
		fmt.Println("")
	}
}
