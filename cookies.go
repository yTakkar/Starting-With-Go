// package main

// import (
// 	"fmt"
// 	"html/template"
// 	"log"
// 	"net/http"
// 	"time"

// 	"github.com/julienschmidt/httprouter"
// 	"github.com/urfave/negroni"
// )

// // Page type but a structure
// type Page struct {
// 	Title  string
// 	Cookie *http.Cookie
// }

// var templates = template.Must(template.ParseGlob("templates/*.html"))

// func renderTemplates(w http.ResponseWriter, tmpl string, r *http.Request, p *Page) {
// 	err := templates.ExecuteTemplate(w, tmpl+".html", p)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		panic(err)
// 	}
// }

// func indexHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
// 	renderTemplates(w, "cookie", r, &Page{"Cookies Show", nil})
// }

// func writeHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
// 	expiration := time.Now().AddDate(0, 0, 1)
// 	cookie := http.Cookie{
// 		Name:    "cookie",
// 		Value:   "Takkar",
// 		Path:    "/",
// 		Expires: expiration,
// 		MaxAge:  86400,
// 	}
// 	http.SetCookie(w, &cookie)
// 	renderTemplates(w, "cookie", r, &Page{"Write Cookies", nil})
// }

// func readWriter(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
// 	cookie, err := r.Cookie("cookie")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(cookie.Value)
// 	renderTemplates(w, "cookie", r, &Page{"Read Cookies", cookie})
// }

// func deleteHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
// 	cookie := http.Cookie{
// 		Name:   "cookie",
// 		Path:   "/",
// 		MaxAge: -1,
// 	}
// 	http.SetCookie(w, &cookie)
// 	renderTemplates(w, "cookie", r, &Page{"Delete Cookies", nil})
// }

// func main() {
// 	router := httprouter.New()

// 	router.GET("/", indexHandler)
// 	router.GET("/read", readWriter)
// 	router.GET("/write", writeHandler)
// 	router.GET("/delete", deleteHandler)

// 	server := negroni.New(
// 		negroni.NewRecovery(),
// 		negroni.NewLogger(),
// 		negroni.NewStatic(http.Dir("public")),
// 		negroni.HandlerFunc(MyMiddleware),
// 	)
// 	server.UseHandler(router)
// 	server.Run(":5050")
// }

// func MyMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
// 	fmt.Println("Logging..")
// 	next(w, r)
// }
