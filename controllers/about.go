package controllers

import (
	"fmt"
	"net/http"
)

func About(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "This is all about Module 1")
}
