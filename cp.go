package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type info struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type happeningNow struct{
	Expectation string `json:"expectation"`
	Reality     string `json:"reality"`
}

func main() {
	http.HandleFunc("/foo", func(w http.ResponseWriter, r *http.Request) {
		userName:="omar"
		password:="12345"
		var temp info
		err := json.NewDecoder(r.Body).Decode(&temp)
		if err != nil {
			fmt.Println(err)
		}
		if temp.Username == userName && temp.Password == password{
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(200)
			toSend:=happeningNow{Expectation:"omar is procrastinating", Reality:"omar is working hard"}
			json.NewEncoder(w).Encode(toSend)
		} else{
			w.WriteHeader(403)
			w.Write([]byte("403 - Something bad happened!"))
		}
	})

	err := http.ListenAndServe(":8000", nil)

	if err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}