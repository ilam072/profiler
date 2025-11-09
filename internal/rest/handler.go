package rest

import (
	"encoding/json"
	"net/http"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

type Req struct {
	A int `json:"a"`
	B int `json:"b"`
}

type Resp struct {
	Sum int `json:"sum"`
}

func (h *Handler) Sum(w http.ResponseWriter, r *http.Request) {
	var req Req
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid json body", http.StatusBadRequest)
		return
	}

	resp := Resp{Sum: req.A + req.B}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "internal server error, try again later", http.StatusInternalServerError)
		return
	}
}
