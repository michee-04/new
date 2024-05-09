package main

import (
	"fmt"
	"net/http"
)

const port = ":5000"


func main() {
	http.HandleFunc("/", Accueil)
	// http.HandleFunc("/footer", Footer)
	http.HandleFunc("/ecole", Ecole)

	fmt.Println("http://localhost:5000 le serveur fonctionne sur le port: ", port)

	http.ListenAndServe(port, nil)
}