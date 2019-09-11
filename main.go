package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// Codes struct
type Code struct {
	Name    string `json:"Name"`
	Country string `json:"Country"`
	Code    string `json:"Code"`
}

var codes []Code

// Get all codes
func getCodes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(codes)
}

// Get single code
func getCode(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get params
	var str = params["name"]
	var strUpper = strings.ToUpper(str)
	// Loop codes
	for _, item := range codes {
		if item.Name == strUpper {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Code{})
}

// Main function
func main() {
	// Init router
	r := mux.NewRouter()

	// data
	codes = append(codes, Code{Name: "BE", Country: "Belgium", Code: "32"})
	codes = append(codes, Code{Name: "BM", Country: "Bermuda", Code: "+1-4441"})
	codes = append(codes, Code{Name: "BJ", Country: "Benin", Code: "229"})
	codes = append(codes, Code{Name: "BY", Country: "Belarus", Code: "375"})
	codes = append(codes, Code{Name: "BZ", Country: "Belize", Code: "501"})

	// Route handles
	r.HandleFunc("/rest", getCodes).Methods("GET")
	r.Path("/rest/code").Queries("name", "{name}").HandlerFunc(getCode).Methods("GET")

	// Start server
	log.Fatal(http.ListenAndServe(":8000", r))
}
