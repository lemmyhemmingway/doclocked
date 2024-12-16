package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Doclocked init..")
	})

	fmt.Println("listening on port 3000")
	http.ListenAndServe(":3000", nil)
}