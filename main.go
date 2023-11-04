package main

import (
	"html/template"
	"net/http"
	"path"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Exp struct {
	Year    string
	Company string
	Site    string
}

type IndexData struct {
	Exps []Exp
}

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/", handler)

	http.ListenAndServe(":8080", r)
}

func handler(w http.ResponseWriter, r *http.Request) {
	filepath := path.Join("views", "index.html")
	tmpl, err := template.ParseFiles(filepath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	Exps := []Exp{
		{
			Year:    "2015 - 2015",
			Company: "Misete Solusi Teknologi",
		},
		{
			Year:    "2015 - 2016",
			Company: "Tulpar Asia Digital",
		},
		{
			Year:    "2016 - 2021",
			Company: "Tech in Asia",
			Site:    "https://id.techinasia.com",
		},
		{
			Year:    "2021 - Now",
			Company: "Brankas",
			Site:    "https://brankas.com",
		},
	}

	data := IndexData{
		Exps: Exps,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
