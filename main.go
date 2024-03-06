package main

import (
	"fmt"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "GO is super easy")
	if err != nil {
		return
	}
}

func contantPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Contants page!")
}
func handeRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/contacts/", contantPage)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
func main() {
	handeRequest()
}
