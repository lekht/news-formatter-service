package api

import (
	"encoding/json"
	"net/http"
)

type Comment struct {
	Text string
}

func (a *API) formatHandler(w http.ResponseWriter, r *http.Request) {
	var c Comment

	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ok := a.f.CheckWord(c.Text)
	if ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}
