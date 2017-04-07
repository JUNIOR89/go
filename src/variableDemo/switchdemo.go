package main

import "fmt"

func main() {
	// os := "OS X"
	os := "linux"
	switch (os) {
	case "OS X":
		fmt.Println("MAC.")
		break
	case "linux":
		fmt.Println("Linus T's system.")
		break
	default:
		fmt.Printf("%s", os)
	}
}
