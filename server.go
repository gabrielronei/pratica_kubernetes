package main

import "net/http"

func main() {
	http.HandleFunc("/", Hello)
	http.ListenAndServe(":8082", nil)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello Gabriel Rocks!</h1>"))
}