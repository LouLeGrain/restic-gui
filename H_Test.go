package main

import (
	"net/http"
)

func TestHandler(w http.ResponseWriter, r *http.Request) {
	pageData := PageData{Title: "App TEST"}
	pageData.Message = "TEST"

	if pageData.Err != "" {
		render(w, "error.html", pageData, "layout")
	} else {
		render(w, "test.html", pageData, "layout")
	}
}
