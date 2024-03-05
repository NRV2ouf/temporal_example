package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

const serverPort int = 3335

const serverPortB int = 3334

func main() {

	go func() {
		// Create a new router
		r := mux.NewRouter()

		// Define routes
		r.HandleFunc("/2", request2Handler).Methods("POST")

		// Start the server
		http.Handle("/", r)
		portStr := ":" + fmt.Sprintf("%v", serverPort)
		http.ListenAndServe(portStr, nil)
	}()

	// Loop
	for {
		time.Sleep(100 * time.Millisecond)
	}
}

func request2Handler(w http.ResponseWriter, r *http.Request) {
	str := r.FormValue("param")
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("server: could not read request body: %s\n", err)
	}
	fmt.Printf("received message: %s\n", &reqBody)

	// Create http request
	jsonBody := []byte(`{"param": "` + str + `"}`)
	bodyReader := bytes.NewReader(jsonBody)

	requestURL := fmt.Sprintf("http://localhost:%d/3", serverPortB)
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
