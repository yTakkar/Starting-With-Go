package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/urfave/negroni"
)

// Payload type
type Payload struct {
	Stuff Data
}

// Data type
type Data struct {
	Fruit   Fruits
	Veggies Vegetables
}

// Fruits type
type Fruits map[string]int

// Vegetables type
type Vegetables map[string]int

func indexHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	d, err := renderJSON()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(d))
	w.Header().Set("Content-Type", "application/json")
	w.Write(d)
}

func main() {
	router := httprouter.New()
	router.GET("/home", indexHandler)
	server := negroni.Classic()
	server.UseHandler(router)
	server.Run(":6061")
}

func renderJSON() ([]byte, error) {
	fruits := make(map[string]int)
	fruits["Apples"] = 25
	fruits["Oranges"] = 10

	veggies := make(map[string]int)
	veggies["Lemon"] = 4
	veggies["Lemon"] = 10

	d := Data{fruits, veggies}
	p := Payload{d}
	data, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return nil, err
	}
	return data, nil
}
