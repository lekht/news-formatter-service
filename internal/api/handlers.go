package api

import (
	"encoding/json"
	"net/http"
)

type Comment struct {
	ID       int    `json:"id"`
	NewsID   int    `json:"news_id"`
	ParentID int    `json:"parent_id"`
	Msg      string `json:"msg"`
	PubTime  int64  `json:"pub_time"`
}

func (a *API) formatHandler(w http.ResponseWriter, r *http.Request) {
	var c Comment

	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ok := a.f.CheckWord(c.Msg)
	if ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}
