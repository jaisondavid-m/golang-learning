package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"slices"
	"strconv"
	"sync"
)

type Student struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Year int    `json:"year"`
}

var students []Student
var nextID = 1
var mu sync.Mutex

func studentHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		mu.Lock()
		result := make([]Student,len(students))
		copy(result,students)
		mu.Unlock()
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(result); err != nil {
			http.Error(w, "Failed to Encode the request", http.StatusInternalServerError)
			return
		}
	case http.MethodPost:
		var s Student
		if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
			fmt.Println("Decode error:", err)
			http.Error(w, "Failed to Decode the request", http.StatusBadRequest)
			return
		}
		mu.Lock()
		s.ID = nextID
		nextID++
		students = append(students, s)
		mu.Unlock()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		response := map[string]string{"message": "User Created Successfullly"}
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, "Failed to Encode the request", http.StatusInternalServerError)
			return
		}
	case http.MethodDelete:
		idStr := r.URL.Query().Get("id")
		if idStr == "" {
			http.Error(w, "Missing id parameter", http.StatusBadRequest)
			return
		}
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid id format", http.StatusBadRequest)
			return
		}
		mu.Lock()
		oldLen := len(students)
		students = slices.DeleteFunc(students, func(s Student) bool {
			return s.ID == id
		})
		deleted := oldLen != len(students)
		mu.Unlock()
		if !deleted {
			http.Error(w,"Student Not Found",http.StatusNotFound)
			return
		}
		response := map[string]string{"message": "Student deleted Successfully"}
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			return
		}
	case http.MethodPut:
		var s Student
		idStr := r.URL.Query().Get("id")
		if idStr == "" {
			http.Error(w, "Missing id parameter", http.StatusBadRequest)
			return
		}
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid id format", http.StatusBadRequest)
			return
		}
		if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
			http.Error(w, "Failed to decode", http.StatusBadRequest)
			return
		}
		found := false
		mu.Lock()
		for i := range students {
			if students[i].ID == id {
				students[i].Name = s.Name
				students[i].Year = s.Year
				found = true
				break
			}
		}
		mu.Unlock()
		if !found {
			http.Error(w, "Student Not Found", http.StatusNotFound)
			return
		}
		response := map[string]string{"message": "Student Updated Successfully"}
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			return
		}
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
}

func main() {

	http.HandleFunc("/students", studentHandler)

	fmt.Println("Server is running on the port: 8000")

	if err := http.ListenAndServe(":8000", nil); err != nil {
		fmt.Println("Failed to run the server:", err)
	}

}
