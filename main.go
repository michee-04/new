package main

import (
	"fmt"
	"net/http"
)

const port = ":5000"

func Home(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Bienvenue sur votre appication web")
}

func Contact(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Je suis très content de moi")
}

func Moi(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Je suis un developpeur backend, je viens de commencer le backend avec le langage Go et j'esperer devenir un boss en langage Golang")
}

func main() {
	// http permet de gerer ce qui va être écrit sur la page
	// Il prend en paramettre la route et la function 
	http.HandleFunc("/", Home)
	http.HandleFunc("/contact", Contact)
	http.HandleFunc("/moi", Moi)

	fmt.Println("http://localhost:5000 server started on port", port)

	http.ListenAndServe(port, nil)
}