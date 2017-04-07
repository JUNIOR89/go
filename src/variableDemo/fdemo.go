package main

import "fmt"

func halfer(val, min int) (retVal int, isError bool) {
	retVal = val/2
	isError = val > min
	return

	//if(val>min){
	//	return val/2
	//} else {
	//	return 0
	//}
}

func main() {
	var whole = 99
	var half int
	var isError bool
	half, isError = halfer(whole,100)
	fmt.Println(half, isError)


}
