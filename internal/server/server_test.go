package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer(t *testing.T) {
	mockQueryDB := MockDatabaseQuery{}
	srv := server{
		mux: http.NewServeMux(),
		db:  &mockQueryDB,
	}
	t.Parallel()
	t.Run("/handleHealth", func(t *testing.T) {
		srv.Routes()
		req := httptest.NewRequest("GET", "/health", nil)
		w := httptest.NewRecorder()
		srv.ServeHTTP(w, req)
		if w.Result().StatusCode != http.StatusOK {
			t.Fatalf("got: %d, want: %d\n", w.Result().StatusCode, http.StatusOK)
		}

	})

	t.Cleanup(func() {})
}
