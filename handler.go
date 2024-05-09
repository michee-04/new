package main

import (
	"net/http"
	"text/template"
)

func Accueil(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Bienvenue sur la page d'accueil")
	renderTmpl(w, "home")
}

// func Footer(w http.ResponseWriter, r *http.Request){
// 	fmt.Fprintf(w, "Le pied de page")
// }

func Ecole(w http.ResponseWriter, r *http.Request) {
	renderTmpl(w, "ecole")
}

func renderTmpl(w http.ResponseWriter, nomFichier string) {
	t, err := template.ParseFiles("template/" + nomFichier + ".page.tmpl")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	t.Execute(w, nil)
}