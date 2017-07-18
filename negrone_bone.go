package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/go-zoo/bone"
	"github.com/urfave/negroni"
)

//Page type
type Page struct {
	Title string
}

var templates = template.Must(template.ParseGlob("templates/*.html"))

func renderTemplates(w http.ResponseWriter, tmpl string, p *Page, r *http.Request) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplates(w, "index", &Page{Title: "Index Page"}, r)
	log.Println(bone.GetValue(r, "name"), bone.GetAllValues(r), bone.GetAllQueries(r))
}

func main() {
	router := bone.New()
	router.Get("/home/:name", http.HandlerFunc(indexHandler))
	server := negroni.Classic()
	server.UseHandler(router)
	server.Run(":6060")
}
