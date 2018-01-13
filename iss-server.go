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
		http.Error(w, "There was an issue with your request", 500)
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(w, "%s", body)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
	fmt.Println("Server is up and running!")
}
