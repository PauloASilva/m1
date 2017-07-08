package controllers

import (
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "This is Module 1 home")
}
