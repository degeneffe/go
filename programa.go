package main

import (
	"fmt"
	"os"
)

func main() {
	for i := 0; i < 15; i++ {
		fmt.Println("Mi programa con GO")
	}
	fmt.Println(os.Args)
	edad := 18
	if edad >= 18 {
		fmt.Println("Eres mayor de edad")
	} else {
		fmt.Println("Eres menor de edad")
	}

}
