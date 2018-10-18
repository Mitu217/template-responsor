package main

import (
	"net/http"

	"github.com/Mitu217/template-responsor/controller"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func getActions(w http.ResponseWriter, r *http.Request) {
	// データベースからルーティングルールを取得
	w.Write([]byte("test"))
}

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	adminActions, err := controller.NewAdminActions()
	if err != nil {
		panic(err)
	}
	r.Route("/admin", adminActions.Router)

	r.Route("/", func(r chi.Router) {
		r.Get("/", getActions)
	})

	http.ListenAndServe(":8888", r)
}
