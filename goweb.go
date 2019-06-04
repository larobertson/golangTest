package main

import ("fmt"
				"net/http")

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	// we can do this, but let's be real
	// we're gonna use templates
	fmt.Fprintf(w, "<h1>Hey there</h1>")
	fmt.Fprintf(w, "<p>What is in a name?</p>")
	fmt.Fprintf(w, "<p>That which is called Go by any other name would still run so sweet</p>")


}

func main() {
	http.HandleFunc("/", IndexHandler)
	http.ListenAndServe(":8000", nil)
}