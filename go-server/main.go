package main

import (
	"fmt"
	"log"
	"net/http"
)

func HandleForm(resp http.ResponseWriter, req *http.Request) {

}
func HandleHello(resp http.ResponseWriter, req *http.Request) {

}
func main() {
	fileServer := http.FileServer(http.Dir("./static")) //returns a Handle
	// declare the routes with handler function
	/**
	To understand the difference between Handle, HandlerFunc, HandleFunc
	Blog : https://perennialsky.medium.com/understand-handle-handler-and-handlefunc-in-go-e2c3c9ecef03
	Reddit : https://www.reddit.com/r/golang/comments/589wf4/whats_the_difference_between_httphandlefunc_and/
	*/
	http.Handle("/", fileServer)
	http.HandleFunc("/form", HandleForm)
	http.HandleFunc("/hello", HandleHello)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Server connected successful")
}
