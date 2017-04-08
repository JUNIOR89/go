package main

import "fmt"

type vertigo struct {
	x, y float64
}


// (associate the struct) (method name, then parameters) (the return type...
func (v vertigo) addVertex(v2 vertigo) (retVert vertigo) {
	retVert.y = v.y + v2.y
	retVert.x = v.x + v2.x
	return

}

func main() {

	vert := vertigo{4.0,2.0}
	vert2 := vertigo{3.0,3.0}
	vert3 := vert.addVertex(vert2)
	fmt.Println(vert, vert3)
}
