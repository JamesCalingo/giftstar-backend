package main

import (
	"fmt"
	"net/http"
)

func (app *application) Home(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Preeow World and welcome to %s", app.Domain)
}
