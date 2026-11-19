package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Student struct {
	Name string `json:"name"`
	Year int `json:"year"`
}

func handler(w http.ResponseWriter,r *http.Request) {
	student := Student{
		Name: "Jaison",
		Year: 2,
	}
	switch r.Method {
	case http.MethodGet:
		w.Header().Set("Content-Type","application/json")
		if err := json.NewEncoder(w).Encode(student); err != nil {
			fmt.Println("JSON encoding error",err)
			return
		}
	default:
		http.Error(w,"Method Not Allowed",http.StatusMethodNotAllowed)
		return
	}
}

func main() {
	http.HandleFunc("/student",handler)
	fmt.Println("Server is running on port:8000")
	if err := http.ListenAndServe(":8000",nil); err != nil {
		fmt.Println("Error while starting server:",err)
	}
}