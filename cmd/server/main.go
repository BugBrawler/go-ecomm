package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Product struct {
	Id   string
	Name string
}

func main() {
	r := chi.NewRouter()
	products := []Product{
		{Id: "1", Name: "Product 1"},
		{Id: "2", Name: "Product 2"},
		{Id: "3", Name: "Product 3"},
	}
	r.Get("/", Index(products))
	http.ListenAndServe(":4321", r)
}

func Index(products []Product) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("templates/pages/index.html")
		if err != nil {
			log.Fatal(err)
		}

		err = tmpl.Execute(w, products)
		if err != nil {
			log.Fatal(err)
		}
	}
}
