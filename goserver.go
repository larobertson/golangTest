package main

import ("fmt"
				"net/http")

func index_handler(w http.ResponseWriter, r *http.Request){
	//http.ResponseWriter is our datatype
	//it is not native to Go, but a custom type
	fmt.Fprintf(w, "Woah, Go is neat!")

}

func main() {
	//begin creating handlers
	//takes in the path we want to use and whatever function we want to run
	http.HandleFunc("/", index_handler)

	//this creates our server at port 8000
	http.ListenAndServe(":8000", nil)
}