package main

import "fmt"

func main() {
	// os := "OS X"
	os := "OS X"

	switch (os) {
	case "OS X":
		fmt.Println("MAC.")
		// Can use what's called fallthrough to show the next case value...
		fallthrough
		// break
	case "linux":
		fmt.Println("Linus T's system.")
		//break
	default:
		fmt.Printf("%s", os)
	}
}
