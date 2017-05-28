package main

import (
	"net/http"
)


func main() {
	http.HandleFunc("/blog/", blogHandler)
	http.HandleFunc("/", indexHandler)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	http.Handle("/fonts/",  http.StripPrefix("/fonts/",  http.FileServer(http.Dir("fonts"))))
	http.ListenAndServe(":8080", nil)
}

