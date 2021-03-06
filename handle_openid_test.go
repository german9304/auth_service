package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestHandleOpenId(t *testing.T) {
	s := server{}
	t.Parallel()

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "http://test", nil)

	t.Run("should respond with 200 ok status", func(t *testing.T) {
		h := s.handleOpenId("8081")
		h(w, req)

		if w.Result().StatusCode != http.StatusOK {
			t.Fatalf("got: %d want %d\n", w.Result().StatusCode, http.StatusOK)
		}
	})

	t.Cleanup(func() {
		logrus.Info("done running tests")
	})
}
