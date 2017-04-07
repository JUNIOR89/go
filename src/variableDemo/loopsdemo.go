package main

import "fmt"

// Call upon the value only....
func main() {
	for _, utf := range "Hello work"  {
		fmt.Println(utf)
	}
}

/*
// Call upon the index and the value inside of the range...
func main() {
	for i, utf := range "Hello work"  {
		fmt.Println(i, ": ", utf)
	}
}


// This creates a while loop...

func main() {
	sum := 0
	i := 0
	for i < 5 {
		sum += i
		i++
		fmt.Println(sum)
	}
}


// There is a difference on how you can call for loop...

func main() {
	sum := 0
	for i := 0 ; i < 5; {
		sum += i
		i++
		fmt.Println(sum)
	}
}

// This creates a loop without the statement declaration inside the loop...
func main() {
	sum := 0
	i := 0
	for ; i < 5; {
		sum += i
		i++
		fmt.Println(sum)
	}
}
*/