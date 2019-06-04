package main

import ("fmt"
				"net/http")

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	// we can do this, but let's be real
	// we're gonna use templates
	fmt.Fprintf(w, "<h1>Hey there</h1>")
	fmt.Fprintf(w, "<p>What is in a name?</p>")
	fmt.Fprintf(w, "<p>That which is called Go by any other name would still run so sweet</p>")
	// can & variables are added as variables in this line vvv
	fmt.Fprintf(w, "<p>You %s even add %s</p>", "can", "<string>variables</strong")

	// or we can use backticks to format it like this
	// it makes viewing source prettier, but really, we'll use templates
	fmt.Fprintf(w, `<h2>Go Saga 2</h2>
									<p>Go is fast!</p>
									<p>... and simple!</p>
									`)
}

func main() {
	http.HandleFunc("/", IndexHandler)
	http.ListenAndServe(":8000", nil)
}