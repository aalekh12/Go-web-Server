package main

import (
	"fmt"
	"log"
	"net/http"
)

type Form struct {
	Name              string
	Phone_No          string
	Age               string
	Highest_Education string
	Tweleve_Marks     string
	Catogary          string
}

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
	var data Form
	data.Name = r.FormValue("name")
	data.Phone_No = r.FormValue("address")
	data.Age = r.FormValue("age")
	data.Catogary = r.FormValue("catogry")
	data.Tweleve_Marks = r.FormValue("marks")
	data.Highest_Education = r.FormValue("heducation")

	fmt.Fprintf(w, "Name= %s \n", data.Name)
	fmt.Fprintf(w, "Phone No= %s \n", data.Phone_No)
	fmt.Fprintf(w, "Age= %s \n", data.Age)
	fmt.Fprintf(w, "Caetgory= %s \n", data.Catogary)
	fmt.Fprintf(w, "Tweleve Marks= %s \n", data.Tweleve_Marks)
	fmt.Fprintf(w, "Highest Education= %s \n", data.Highest_Education)
	// jsondata, _ := json.Marshal(data)
	// fmt.Println(jsondata)

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
