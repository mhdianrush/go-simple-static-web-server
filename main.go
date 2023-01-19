package main

import (
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
)

func helloHandler(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/hello" {
		http.Error(writer, "404 Not Found", http.StatusNotFound)
		return
	}
	if request.Method != "GET" {
		// for example we user GET http method
		http.Error(writer, "Method Is Not Supported", http.StatusNotFound)
		return
	}
	fmt.Fprint(writer, "Hello!")
}

func formHandler(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	// using ParseForm because user in the client side will send a data to a form
	if err != nil {
		fmt.Fprintf(writer, "ParseForm() err is%v", err)
		return
	}
	fmt.Fprint(writer, "POST request is SuccessFul\n")
	name := request.FormValue("name")
	address := request.FormValue("address")
	// name and address variable adove is from the form.html file inside the static folder format for html
	fmt.Fprintf(writer, "Name is %s\n", name)
	fmt.Fprintf(writer, "Address is %s\n", address)
}

func main() {
	logger := logrus.New()
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/form", formHandler)

	fmt.Println("Starting Server at Port 8080")

	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		logger.Fatal(err)
	}
}
