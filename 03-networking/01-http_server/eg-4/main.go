package main

import (
	"fmt"
	"net/http"
)

func student(w http.ResponseWriter,r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		fmt.Fprintln(w, "Getting Student data")
	case http.MethodPost:
		fmt.Fprintln(w, "Creating student")
	default:
		http.Error(w, "Method is not allowed",http.StatusMethodNotAllowed)
		return
	}

}

func main() {

	http.HandleFunc("/students",student)

	fmt.Println("Server listening on port:8000")

	if err := http.ListenAndServe(":8000",nil); err != nil {
		fmt.Println("Server StartUp Error:",err)
	}

}