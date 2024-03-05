package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

const serverPort int = 3333

const serverPortB int = 3334

func main() {

	// Create a new router
	r := mux.NewRouter()

	// Define routes
	r.HandleFunc("/0", request0Handler).Methods("POST")
	r.HandleFunc("/4", request4Handler).Methods("POST")

	// Start the server
	http.Handle("/", r)
	portStr := ":" + fmt.Sprintf("%v", serverPort)
	http.ListenAndServe(portStr, nil)

}

func request0Handler(w http.ResponseWriter, r *http.Request) {
	str := r.FormValue("param")
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("server: could not read request body: %s\n", err)
	}
	fmt.Printf("received message: %s\n", &reqBody)

	// Create http request
	jsonBody := []byte(`{"param": "` + str + `"}`)
	bodyReader := bytes.NewReader(jsonBody)

	requestURL := fmt.Sprintf("http://localhost:%d/1", serverPortB)
	req, err := http.NewRequest(http.MethodPost, requestURL, bodyReader)
	if err != nil {
		fmt.Printf("error creating http request: %s\n", err)
	}

	// Send request 1 to B
	_, err = http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("error sending http request: %s\n", err)
	}
}

func request4Handler(w http.ResponseWriter, r *http.Request) {
	str := r.FormValue("param")
	fmt.Fprintf(w, "received message: %s", str)
}
