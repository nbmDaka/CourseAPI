package main

import (
	"fmt"
	"net/http"
)

func (app *application) getModuleInfoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "status: available")
	fmt.Fprintf(w, "environment: %s\n", app.config.env)
	fmt.Fprintf(w, "version: %s\n", version)
}
func (app *application) createModuleInfoHandler(http.ResponseWriter, *http.Request) {

}
func (app *application) showModuleInfoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("THIS IS THE HANDLER FOR HEALTHCHECK")
}
func (app *application) editModuleInfoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("THIS IS THE HANDLER FOR HEALTHCHECK")
}
func (app *application) deleteModuleInfoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("THIS IS THE HANDLER FOR HEALTHCHECK")
}
