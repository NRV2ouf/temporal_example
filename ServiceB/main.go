package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

const serverPort int = 3334

const serverPortA int = 3333
const serverPortC int = 3335

func main() {

	go func() {
		// Create a new router
		r := mux.NewRouter()

		// Define routes
		r.HandleFunc("/1", request1Handler).Methods("POST")
		r.HandleFunc("/3", request3Handler).Methods("POST")

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

func request1Handler(w http.ResponseWriter, r *http.Request) {
	str := r.FormValue("param")
	fmt.Printf("received message: %s", str)

	// Create http request
	jsonBody := []byte(`{"param": "` + str + `"}`)
	bodyReader := bytes.NewReader(jsonBody)

	requestURL := fmt.Sprintf("http://localhost:%d/2", serverPortC)
	req, err := http.NewRequest(http.MethodPost, requestURL, bodyReader)
	if err != nil {
		fmt.Printf("error creating http request: %s\n", err)
	}

	// Send request 2 to C
	_, err = http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("error sending http request: %s\n", err)
	}
}

func request3Handler(w http.ResponseWriter, r *http.Request) {
	str := r.FormValue("param")
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("server: could not read request body: %s\n", err)
	}
	fmt.Printf("received message: %s\n", &reqBody)

	// Create http request
	jsonBody := []byte(`{"param": "` + str + `"}`)
	bodyReader := bytes.NewReader(jsonBody)

	requestURL := fmt.Sprintf("http://localhost:%d/4", serverPortA)
	req, err := http.NewRequest(http.MethodPost, requestURL, bodyReader)
	if err != nil {
		fmt.Printf("error creating http request: %s\n", err)
	}

	// Send request 4 to A
	_, err = http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("error sending http request: %s\n", err)
	}
}
