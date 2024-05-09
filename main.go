package main

import (
	"errors"
	"fmt"
)

func division(a, b float32) (float32, error) {
	if b == 0 {
		return 0, errors.New("impossible de diviser par 0")
	}

	return a / b, nil
}

func main() {
	resultat, err := division(5, 2)
	if err != nil {
		panic(err)
	}

	fmt.Println(resultat)
}