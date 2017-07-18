package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
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

func indexHandler(w http.ResponseWriter, r *http.Request, pm httprouter.Params) {
	renderTemplates(w, "index", &Page{Title: "Index Page"}, r)
	log.Println(pm.ByName("name"))
}

func main() {
	router := httprouter.New()
	router.GET("/home/:name", indexHandler)
	server := negroni.New(
		negroni.NewRecovery(),
		negroni.NewLogger(),
		negroni.NewStatic(http.Dir("public")),
		negroni.HandlerFunc(MyMiddleware),
	)
	server.UseHandler(router)
	server.Run(":6060")
}

//MyMiddleware function
func MyMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	log.Println("Logging..")
	next(w, r)
}
