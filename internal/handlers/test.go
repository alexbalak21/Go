package handlers

import (
	"fmt"
	"net/http"
)

func GetResponseHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Get Response")
}

func PostResponseHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Post Response")
}

func PutResponseHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Put Response")
}

func DeleteResponseHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Delete Response")
}
