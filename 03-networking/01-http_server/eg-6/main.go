package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type student struct {
	Name 	string 		`json:"name"`
	Year 	int 		`json:"year"`
}

func handler(w http.ResponseWriter,r *http.Request) {

	switch r.Method {

	case http.MethodPost:
		var s student
		err := json.NewDecoder(r.Body).Decode(&s)
		if err != nil {
			fmt.Println("Failed to decode")
			return
		}
		fmt.Println(s)
		fmt.Println("Name:",s.Name)
		fmt.Println("Year:",s.Year)
		
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