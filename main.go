package main

import (
	"fmt"
	"mics/functions"
)

func calculvaleur(intChan chan int) {
	randomNombre := functions.RandomValue(50)
	intChan <- randomNombre
}
 
func main() {
	fmt.Println("cool")
	random := make(chan int)

	go calculvaleur(random)

	 num := <- random
	 fmt.Println(num) 
}