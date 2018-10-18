package controller

import (
	"html/template"
	"net/http"

	"github.com/Mitu217/template-responsor/repository"
	"github.com/go-chi/chi"
)

type AdminActions struct {
	responseRepository repository.Response
}

func NewAdminActions() (*AdminActions, error) {
	responseRepository, err := repository.NewSQLiteResponseRepository("./sqlite.db")
	if err != nil {
		panic(err)
	}
	return &AdminActions{
		responseRepository: responseRepository,
	}, nil
}

func (a *AdminActions) getResponseList(w http.ResponseWriter, r *http.Request) {
	ress, err := a.responseRepository.Gets()
	if err != nil {
		panic(err)
	}

	data := map[string]interface{}{
		"len":  len(ress),
		"ress": ress,
	}
	t := template.Must(template.ParseFiles("templates/list.html.tpl"))
	if err := t.ExecuteTemplate(w, "list.html.tpl", data); err != nil {
		panic(err)
	}
}

func (a *AdminActions) getCreateForm(w http.ResponseWriter, r *http.Request) {

}

func (a *AdminActions) getEditForm(w http.ResponseWriter, r *http.Request) {

}

func (a *AdminActions) createResponse(w http.ResponseWriter, r *http.Request) {

	// リダイレクト
}

func (a *AdminActions) updateResponse(w http.ResponseWriter, r *http.Request) {

	// リダイレクト
}

func (a *AdminActions) deleteResponse(w http.ResponseWriter, r *http.Request) {

	// リダイレクト
}

func (a *AdminActions) Router(r chi.Router) {
	r.Route("/", func(r chi.Router) {
		r.Get("/list", a.getResponseList)
		r.Get("/create", a.getCreateForm)
		r.Get("/edit", a.getEditForm)
		r.Post("/create", a.createResponse)
		r.Post("/update", a.updateResponse)
		r.Post("/delete", a.deleteResponse)
	})
}
