package main

import ("fmt"
				"net/http")

func IndexHandler(w http.ResponseWriter, r *http.Request){
	// http.ResponseWriter is our datatype
	// it is not native to Go, but a custom type
	fmt.Fprintf(w, "Woah, Go is neat!")
}

func AboutHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Learn more about Go for the web!")
}

func HelloHandler(w http.ResponseWriter, r *http.Request){
	// remember each time we add a new HandleFunc we have to restart server
	fmt.Fprintf(w, "Hello World!")
}

func main() {
	// begin creating handlers
	// takes in the path we want to use and whatever function we want to run
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/about/", AboutHandler)
	http.HandleFunc("/hello/", HelloHandler)



	//this creates our server at port 8000
	http.ListenAndServe(":8000", nil)
}