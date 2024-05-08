package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Id    int    `json:"id"`
	Nom   string `json:"nom"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func main() {
	jsonApi := `
		[ 
			{
				"id": 1,
				"nom": "michée",
				"age": 20,
				"email": "mic@gmail.com"
			},
			{
				"id": 2,
				"nom": "michou",
				"age": 21,
				"email": "mics@gmail.com"
			}
		]
	`

	var users []User

	err := json.Unmarshal([]byte(jsonApi), &users)

	if err != nil {
		fmt.Println("Erreur json", err)
	}

	fmt.Printf("Données json %v\n", users)

	var monStruct []User
	var user_1 User
	user_1.Age = 2
	user_1.Email = "mon@gmqil.com"
	user_1.Id = 3
	user_1.Nom = "Mon"

	monStruct = append(monStruct, user_1)

	jsonSlice, err := json.MarshalIndent(monStruct, "", " ")
	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println(string(jsonSlice))
}