package main

import ("fmt"
				"net/http")

func index_handler(w http.ResponseWriter, r *http.Request){
	// http.ResponseWriter is our datatype
	// it is not native to Go, but a custom type
	fmt.Fprintf(w, "Woah, Go is neat!")
}

func about_handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Learn more about Go for the web!")
}

func hello_handler(w http.ResponseWriter, r *http.Request){
	// remember each time we add a new HandleFunc we have to restart server
	fmt.Fprintf(w, "Hello World!")
}

func main() {
	// begin creating handlers
	// takes in the path we want to use and whatever function we want to run
	http.HandleFunc("/", index_handler)
	http.HandleFunc("/about/", about_handler)
	http.HandleFunc("/hello/", hello_handler)



	//this creates our server at port 8000
	http.ListenAndServe(":8000", nil)
}