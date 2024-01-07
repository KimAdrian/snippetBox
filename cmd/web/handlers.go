package main

import (
	"errors"
	"fmt"
	"github.com/KimAdrian/snippetBox/internal/model"
	"net/http"
	"strconv"
)

func (app *application) homeHandler(writer http.ResponseWriter, request *http.Request) {
	//Restrict root pattern from being a fallback option
	//Send 404 response of url specified is not recognized and return
	if request.URL.Path != "/" {
		app.notFound(writer)
		return
	}

	snippets, err := app.snippets.Latest()
	if err != nil {
		app.serverError(writer, err)
		return
	}

	for _, snippet := range snippets {
		fmt.Fprintf(writer, "%+v\n", snippet)
	}
	//Initialise slice containing path to templates
	//templateFiles := []string{
	//	"./ui/html/templates/base.html",
	//	"./ui/html/templates/nav.html",
	//	"./ui/html/templates/pages/home.html",
	//}
	////Read template file into template set
	//ts, err := template.ParseFiles(templateFiles...)
	//if err != nil {
	//	app.serverError(writer, err)
	//	return
	//}
	//
	////Write template content as the response body
	//err = ts.ExecuteTemplate(writer, "base", nil)
	//if err != nil {
	//	app.serverError(writer, err)
	//}
}

func (app *application) viewSnippetsHandler(writer http.ResponseWriter, request *http.Request) {
	//Fetch id parameter from url and convert it to int
	id, err := strconv.Atoi(request.URL.Query().Get("id"))
	//If id cannot be converted or is less than 1 respond with 404 and return
	if err != nil || id < 1 {
		app.notFound(writer)
		return
	}

	snippet, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, model.ErrNoRecord) {
			app.notFound(writer)
		} else {
			app.serverError(writer, err)
		}
		return
	}

	//Write the snippet data as a plain textHTTP response body
	fmt.Fprintf(writer, "%+v", snippet)
}

func (app *application) createSnippetHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		writer.Header().Set("Allow", http.MethodPost)
		app.clientError(writer, http.StatusMethodNotAllowed)
		return
	}

	//dummy data
	title := "O snail"
	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi Issa"
	expires := 7

	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(writer, err)
		return
	}
	// Redirect the user to the relevant page for the snippet.
	http.Redirect(writer, request, fmt.Sprintf("/snippet/view?id=%d", id), http.StatusSeeOther)
}
