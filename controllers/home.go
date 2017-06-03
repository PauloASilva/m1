package controllers

import (
	"fmt"
	"net/http"
)

func HomeRegister(router *http.ServeMux) {
	router.HandleFunc("/m1/", home)
}

func home(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "This is Module 1 home")
}
