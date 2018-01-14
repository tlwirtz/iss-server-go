package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("https://api.wheretheiss.at/v1/satellites/25544")

	if err != nil {
		http.Error(w, "There was an issue with your request - iss-server-go", 500)
		log.Fatal(err)
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		http.Error(w, "There was an error with your request - iss-server-go-2", 500)
		log.Fatal(err)
		return
	}

	fmt.Fprintf(w, "%s", body)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
	log.Println("Server is up and running!")
}
