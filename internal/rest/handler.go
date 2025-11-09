package rest

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

type Resp struct {
	Sum int `json:"sum"`
}

func (h *Handler) Sum(w http.ResponseWriter, r *http.Request) {
	a, err := strconv.Atoi(r.URL.Query().Get("a"))
	if err != nil {
		http.Error(w, "'a' must be integer", http.StatusBadRequest)
		return
	}
	b, err := strconv.Atoi(r.URL.Query().Get("b"))
	if err != nil {
		http.Error(w, "'b' must be integer", http.StatusBadRequest)
		return
	}

	resp := Resp{Sum: a + b}

	w.Header().Set("Content-Type", "application/json")
	bytes, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, "internal server error, try again later", http.StatusInternalServerError)
		return
	}

	w.Write(bytes)
}
