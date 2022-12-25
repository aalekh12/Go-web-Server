package main

import (
	"fmt"
	"log"
	"net/http"
)

func hellohandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "Error 404 Page not Found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method Not Found", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "hello")
}

func formhandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Parseform() %v", err)
		return
	}
	fmt.Fprintf(w, "Post Request Successfull \n")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name= %s \n", name)
	fmt.Fprintf(w, "Phone No= %s \n", address)

}

func main() {
	fileserver := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileserver)
	http.HandleFunc("/hello", hellohandler)
	http.HandleFunc("/form", formhandler)

	fmt.Println("Starting Server at localhost 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
