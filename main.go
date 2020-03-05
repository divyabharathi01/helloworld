package main

import (
	"fmt"

	"github.com/divyabharathi01/helloworld/supply"
)

func main() {
	// fmt.Println("my  favourite numberis", rand.Float32())
	var p supply.Person
	p = supply.Person{Name: "Divya"}
	fmt.Println(p.Name)
}
