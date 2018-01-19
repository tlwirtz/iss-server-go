package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

//TODO -- Need to check for 429 (too many reqs) and back off on the requests
func getISSData() ([]byte, error) {
	resp, err := http.Get("https://api.wheretheiss.at/v1/satellites/25544")

	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return body, nil

}

func handler(w http.ResponseWriter, r *http.Request) {
	resp, err := getISSData()

	if err != nil {
		http.Error(w, "There was an issue with your request - iss-server-go", 500)
		log.Println(err)
		return
	}

	if err != nil {
		http.Error(w, "There was an error with your request - iss-server-go", 500)
		log.Println(err)
		return
	}

	fmt.Fprintf(w, "%s", resp)
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func sockHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println(err)
		return
	}

	log.Println("connected")

	for {
		time.Sleep(time.Second * 2)
		resp, err := getISSData()

		if err != nil {
			log.Println(err)
			return
		}

		if err := conn.WriteMessage(websocket.TextMessage, resp); err != nil {
			log.Println(err)
			return
		}
	}

}

func main() {
	port := ":8080"
	http.HandleFunc("/", handler)
	http.HandleFunc("/socket", sockHandler)
	log.Println("Starting server on port", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
