package main

import (
	"net/http"
)

func IndexHandle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Index Part !"))
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("First exercise !"))
	})
	http.HandleFunc("/Index", IndexHandle)

	http.ListenAndServe(":8080", nil)
}
