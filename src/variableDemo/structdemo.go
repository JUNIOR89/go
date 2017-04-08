package main

import "fmt"

type smile struct {
	z float64
}

type vertex struct {
	x, y float64
}

type gameCharacter struct{
	vertex
	life int
	smile
}

// Careful with having the same variables amongst different structures in case you decide to leave the variable identifier out when setting the value...
func main() {
	//vert := vertex{1.0, 2.0}
	//smiles := smile{3.5}

	//mainChar := gameCharacter {
	//	vertex:vert,
	//	life : 10,
	//	smile:smiles}

	mainChar := gameCharacter{
		vertex:vertex{1.0, 2.0},
		life:   10,
		smile:smile{3.5},
}
	//fmt.Println(vert, mainChar)
	fmt.Println(mainChar)






	/*
	type smile struct {
		x float64
	}

	type vertex struct {
		x, y float64
	}

	type gameCharacter struct{
		vertex
		life int
		smile
	}

	func main() {
		vert := vertex{1.0, 2.0}
		var mainChar gameCharacter
		mainChar.life = 10
		mainChar.vertex.x = 1.0
		mainChar.vertex.y = 2.0
		mainChar.smile.x  = 9.0
		fmt.Println(vert, mainChar)

	*/
}


/*
// if a value is not assigned then you can assign only one value and not the other because it defaults to 0 due to the value float64...
func main() {
	vert:=vertex{x:1.0}

	fmt.Println(vert)
}
*/


/*
func main() {
	vert:=vertex{1.0, 1.0}

	fmt.Println(vert)
}

*/