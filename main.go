package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.Handle("/", &myHandler{})

	fmt.Println("Listening on :8000...")

	err := http.ListenAndServe(":8000", mux)
	if err != nil {
		log.Fatal(err)
	}
}

type myHandler struct{}

func (h *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
