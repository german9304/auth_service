package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestHandleHealthOkStatus(t *testing.T) {
	s := server{}
	t.Parallel()
	req := httptest.NewRequest(http.MethodPost, "http://localhost:8080/test", nil)
	res := httptest.NewRecorder()
	handler := s.handleHealth()
	handler.ServeHTTP(res, req)
	if res.Result().StatusCode != http.StatusOK {
		logrus.Fatalf("got: %d expected: %d\n", res.Result().StatusCode, http.StatusOK)
	}
}
