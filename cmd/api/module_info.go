package main

import (
	"courseAPI/internal/data"
	"fmt"
	"net/http"
)

func (app *application) createModuleInfoHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		ModuleName     string `json:"moduleName"`
		ModuleDuration int32  `json:"moduleDuration"`
		ExamType       string `json:"examType"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	moduleInfo := &data.ModuleInfo{
		ModuleName:     input.ModuleName,
		ModuleDuration: input.ModuleDuration,
		ExamType:       input.ExamType,
	}

	err = app.models.ModuleInfo.Insert(moduleInfo)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/movies/%d", moduleInfo.ID))
	err = app.writeJSON(w, http.StatusCreated, envelope{"moduleInfo": moduleInfo}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

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
