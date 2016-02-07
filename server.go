package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type errorHandler struct {
	message string
}

// http get - `/`
func indexHandler(res http.ResponseWriter, req *http.Request) {

	// set the content type
	res.Header().Set("Content-Type", "text/html")
	data, err := ioutil.ReadFile("assets/index.html")
	if err != nil {
		panic(err)
	}

	// return the data along with a Status Ok header
	//res.WriteHeader(http.StatusOk)
	res.Write(data)
}

// http get - `/about`
func aboutHandler(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-Type", "text/html")
	data, err := ioutil.ReadFile("assets/about.html")
	if err != nil {
		panic(err)
	}

	//res.WriteHeader(http.StatusOk)
	res.Write(data)
}

func contactHandler(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-Type", "text/html")
	data, err := ioutil.ReadFile("assets/contact.html")
	if err != nil {
		panic(err)
	}

	//res.WriteHeader(http.StatusOk)
	res.Write(data)
}

func main() {

	// We are using the function handlers for '/, /about, /contact'
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/contact", contactHandler)

	server := &http.Server{
		Addr:           ":8000",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Println("Front End Server is up and running")
	server.ListenAndServe()
}
