package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Student struct {
	Name 	string 		`json:"name"`
	Year 	int 		`json:"year"`
}

func handler(w http.ResponseWriter,r *http.Request) {

	switch r.Method {

	case http.MethodPost:
		var s Student
		err := json.NewDecoder(r.Body).Decode(&s)
		if err != nil {
			http.Error(w,"Failed to decode",http.StatusBadRequest)
			return
		}
		fmt.Println(s)
		fmt.Println("Name:",s.Name)
		fmt.Println("Year:",s.Year)
		fmt.Fprintln(w,"Student registered successfully")
		
	default:
		http.Error(w,"Method Not Allowed",http.StatusMethodNotAllowed)
		return
	}

}

func main() {

	http.HandleFunc("/",handler)

	fmt.Println("Server is listening on port:8000")

	if err := http.ListenAndServe(":8000",nil); err != nil {
		fmt.Println("Server Failed to Start:",err)
	}

}