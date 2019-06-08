package main

import ("fmt"
				"net/http"
				"html/template")

type NewsAggPage struct {
	Title string
	News string
}

func newsAggHandler(w http.ResponseWriter, r *http.Request) {
	p := NewsAggPage{Title: "Amazing News Aggregator", News: "some news"}
	t, _ := template.ParseFiles("basictemplating.html")
	t.Execute(w, p)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Woah, Go is neat!</h1>")
}

func main() {
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/agg/", newsAggHandler)
	http.ListenAndServe(":8000", nil)
}