package rest

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDummy(t *testing.T) {}

func BenchmarkSum(b *testing.B) {
	h := &Handler{}
	reqBody, _ := json.Marshal(Req{A: 1, B: 2})

	for i := 0; i < b.N; i++ {
		req := httptest.NewRequest(http.MethodPost, "/sum", bytes.NewReader(reqBody))
		w := httptest.NewRecorder()

		h.Sum(w, req)

		if w.Code != http.StatusOK {
			b.Fatalf("unexpected status: %d", w.Code)
		}
	}
}
