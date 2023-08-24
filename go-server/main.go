package main

import (
	"fmt"
	"log"
	"net/http"
)

func HandleForm(resp http.ResponseWriter, req *http.Request) {
	//handle errros
	if req.URL.Path != "/form" {
		http.Error(resp, "PATH_NOT FOUND", http.StatusNotFound)
		return
	} else if req.Method != "POST" {
		http.Error(resp, "METHOD_NOT_ALLOWED", http.StatusMethodNotAllowed)
		return

	}
	err := req.ParseForm()
	if err != nil {
		fmt.Fprintf(resp, "Error while parsing : %v", err)
		return
	}
	fmt.Fprintln(resp, "POST Req sucuessful")
	name := req.FormValue("name")
	address := req.FormValue("address")
	fmt.Fprintln(resp, name)
	fmt.Fprintln(resp, address)
}
func HandleHello(resp http.ResponseWriter, req *http.Request) {
	//handle errors
	if req.URL.Path != "/hello" {
		http.Error(resp, "PATH_NOT FOUND", http.StatusNotFound)
		return
	} else if req.Method != "GET" {
		http.Error(resp, "METHOD_NOT_ALLOWED_BY_SERVER", http.StatusMethodNotAllowed)
		return

	}
	fmt.Fprintf(resp, "hello !")
}
func main() {
	fmt.Println("Starting main function")
	fileServer := http.FileServer(http.Dir("./static")) //returns a Handle
	// declare the routes with handler function
	http.Handle("/", fileServer)
	http.HandleFunc("/form", HandleForm)
	http.HandleFunc("/hello", HandleHello)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Server connected successful")
}
