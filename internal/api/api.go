package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lekht/news-formatter-service/internal/format"
)

type API struct {
	r *mux.Router
	f *format.Formatter
}

// Регистрация методов в маршрутизаторе
func (a *API) endpoints() {
	a.r.Name("formatting").Path("/format").Methods(http.MethodPost).HandlerFunc(a.formatHandler)

}

func (a *API) Router() *mux.Router {
	return a.r
}

// Конструктор API
func New(f *format.Formatter) *API {
	a := API{
		r: mux.NewRouter(),
		f: f,
	}
	a.endpoints()
	return &a
}
