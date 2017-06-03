package controllers

import (
	"fmt"
	"net/http"
)

func AboutRegister(router *http.ServeMux) {
	router.HandleFunc("/m1/about", about)
}

func about(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "This is all about Module 1")
}
