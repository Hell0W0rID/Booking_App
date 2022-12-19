package main

import (
	"net/http"
)

func main() {

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	_, err := fmt.Fprintf(w, "Hello World")
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// })
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.ListenAndServe(":8081", nil)
}
