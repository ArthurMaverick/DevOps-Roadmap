package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Profile struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func main() {
	http.HandleFunc("/profile", showProfile)
	fmt.Println("api is on :8080")

	log.Fatal(http.ListenAndServe(":8000", nil))

}

func showProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode([]Profile{{
		ID:   1,
		Name: "John",
	}})
}
