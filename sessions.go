package main

import (
	"fmt"
	"github.com/gorilla/sessions"
	"github.com/urfave/negroni"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"log"
)

var store = sessions.NewCookieStore([]byte("Very-secret"))

func indexHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	session, err := store.Get(r, "session")
	if err != nil{
		log.Fatal(err)
	}

	session.Values["username"] = "Takkar"
	session.Values["age"] = 18

	fmt.Println(session.Values["username"])

	delete(session.Values, "username")
	delete(session.Values, "age")
	fmt.Println(session)

	session.Save(r, w)
	w.Write([]byte("Hello, World!!"))
}

func main(){
	router := httprouter.New()
	router.GET("/", indexHandler)
	server := negroni.Classic()
	server.UseHandler(router)
	server.Run(":5050")
}