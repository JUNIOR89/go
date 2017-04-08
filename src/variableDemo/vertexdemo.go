package main

import (
	"fmt"
	"vertex"
)


func main() {
	vert := vertex.Vertex{12,0,1}
	fmt.Println(vert.X, vert.Y, vert.Z)
}



/*
import (
	"fmt"
)

type vertex struct{
	x, y, z int
}



func main() {
	var vert vertex
	vert.x = 12
	vert.z = 1
	fmt.Println(vert.x, vert.y, vert.z)
}

*/