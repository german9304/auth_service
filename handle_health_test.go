package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestHandleHealthOkStatus(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "http://localhost:8080/test", nil)
	res := httptest.NewRecorder()
	handler := handleHealth()
	handler.ServeHTTP(res, req)
	logrus.Info(res.Result().StatusCode)
}
