package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type MyHandler struct{}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func top(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("tmpl.html")
	if err != nil {
		log.Println(err)
	}
	err = t.Execute(w, "Hello World111!")
	if err != nil {
		return
	}
}

func main() {
	http.HandleFunc("/top", top)

	//http.ListenAndServe(":8080", &MyHandler{})
	http.ListenAndServe(":8080", nil)
}
