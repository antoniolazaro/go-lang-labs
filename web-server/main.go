package main

import (
				"encoding/json"
				"fmt"
				"log"
        "net/http"
)

type profile struct {
	Name    string
	Hobbies []string
}

func main() {
				http.HandleFunc("/welcome", welcome)

				fs := http.FileServer(http.Dir("static"))
				http.Handle("/", fs)

				http.HandleFunc("/profile", getProfile)
				
				if err := http.ListenAndServe(":8080", nil); err != nil {
					log.Fatal(err)
	}
}


func getProfile(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
					http.Error(w, "Invalid request method.", http.StatusMethodNotAllowed)
					return
	}

	profile := profile{
					Name:    "Yashish",
					Hobbies: []string{"sports", "programming"},
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(profile); err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func welcome(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Welcome to the Go webserver codelab")
}