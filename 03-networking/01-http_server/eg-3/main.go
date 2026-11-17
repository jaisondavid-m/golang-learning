package main

import (
	"fmt"
	"net/http"
)

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This about route of the Go Server")
}

func welcome(w http.ResponseWriter,r *http.Request) {
	fmt.Fprintln(w,"Welcome to the Go Server")
}

func greet(w http.ResponseWriter,r *http.Request){

	name := r.URL.Query().Get("name")
	year := r.URL.Query().Get("year")

	if name == "" || year == "" {
		fmt.Fprintln(w,"Name and year parameters are required")
		return
	}

	fmt.Fprintln(w,"Hello!",name,"Your are a",year,"Year Student")
}

func main() {
	
	http.HandleFunc("/", welcome)
	http.HandleFunc("/about", about)
	http.HandleFunc("/greet", greet)

	fmt.Println("Server is running on http://localhost:8000")

	if err := http.ListenAndServe(":8000", nil); err != nil {
		fmt.Println("Server Error:", err)
	}

}