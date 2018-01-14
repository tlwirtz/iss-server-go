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
		log.Println(err)
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		http.Error(w, "There was an error with your request - iss-server-go", 500)
		log.Println(err)
		return
	}

	fmt.Fprintf(w, "%s", body)
}

func main() {
	port := ":8080"
	http.HandleFunc("/", handler)
	log.Println("Starting server on port", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
