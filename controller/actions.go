package controller

import "github.com/go-chi/chi"

type Actions interface {
	Router() chi.Router
}
