package main

import "fmt"

func main() {
	var num = 5
	for i := 2; i <= num; i++ {
		for j := 2; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
}
