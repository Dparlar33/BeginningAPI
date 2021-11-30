package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	var i ironman
	var w wolverine
	x1 := &messageHandler{"One message !"}
	x2 := &messageHandler{"One more message !"}
	mux := http.NewServeMux()
	mux.Handle("/ironman", i)
	mux.Handle("/wolverine", w)
	mux.Handle("/one", x1)
	mux.Handle("/onemore", x2)
	log.Println("Waiting....")

	http.ListenAndServe(":8080", mux)

}

type messageHandler struct {
	message string
}

func (x *messageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, x.message)
}

type wolverine int

func (w wolverine) ServeHTTP(res http.ResponseWriter, r *http.Request) {
	io.WriteString(res, "Mr Wolverine!")
}

type ironman int

func (i ironman) ServeHTTP(res http.ResponseWriter, r *http.Request) {
	io.WriteString(res, "Mr ironman!")
}
