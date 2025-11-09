package rest

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDummy(t *testing.T) {}

func BenchmarkSum(b *testing.B) {
	h := &Handler{}

	for i := 0; i < b.N; i++ {
		req := httptest.NewRequest(http.MethodGet, "/sum2?a=1&b=2", nil)
		w := httptest.NewRecorder()

		h.Sum(w, req)

		if w.Code != http.StatusOK {
			b.Fatalf("unexpected status: %d", w.Code)
		}
	}
}
