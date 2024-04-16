package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *application) routes() *httprouter.Router {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/v1/module_info", app.getModuleInfoHandler)
	router.HandlerFunc(http.MethodPost, "/v1/module_info", app.createModuleInfoHandler)
	router.HandlerFunc(http.MethodGet, "/v1/module_info/:id", app.showModuleInfoHandler)
	router.HandlerFunc(http.MethodPatch, "/v1/module_info/:id", app.editModuleInfoHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/module_info/:id", app.deleteModuleInfoHandler)
	return router

}
