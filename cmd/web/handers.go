package main

import (
	"errors"
	"fmt"
	_ "html/template"
	"net/http"
	"strconv"

	"github.com/benjaminpodmore.com/web-server/internal/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	snippets, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	for _, snippet := range snippets {
		fmt.Fprintf(w, "%+v\n", snippet)
	}
	//
	// files := []string{
	// 	"ui/html/pages/base.tmpl.html",
	// 	"ui/html/pages/home.tmpl.html",
	// 	"ui/html/partials/nav.tmpl.html",
	// }
	//
	// t, err := template.ParseFiles(files...)
	// if err != nil {
	// 	app.serverError(w, r, err)
	// 	return
	// }
	// err = t.ExecuteTemplate(w, "base", nil)
	// if err != nil {
	// 	app.serverError(w, r, err)
	// }
}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.NotFound(w, r)
		return
	}
	snippet, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			http.NotFound(w, r)
		} else {
			app.serverError(w, r, err)	
		}
		return
	}
	
	fmt.Fprintf(w, "%+v", snippet)
}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {

}

func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	title := "0 snail"
	content := "0 snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n- Kobayashi Issa"
	expires := 7

	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)

}
