package main


import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	//val := 5
	if random := rand.Intn(10); random < 2 {
		fmt.Println("our number is < 2")
	} else if random < 7 {
		fmt.Println("our number is < 7")
	} else {
		fmt.Println("our number is >= 7")
	}
}



/*
func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	random := rand.Intn(10)
	//val := 5
	if random < 2 {
		fmt.Println("our number is < 2")
	} else if random < 7 {
		fmt.Println("our number is < 7")
	} else {
		fmt.Println("our number is >= 7")
	}
	fmt.Println("our number is ", random)
}*/