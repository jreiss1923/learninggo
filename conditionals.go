//go:build ignore

package main

import "fmt"

func main() {

	if 7%2 == 0 && 8%2 == 1 {
		fmt.Println("7 is even")
	} else if 7%2 == 1 {
		fmt.Println("7 is odd")
	}
}
