package server

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// MockDatabaseQuery to mock D
type MockDatabaseQuery struct{}

func (md *MockDatabaseQuery) CreateUser(ctx context.Context, user User) (int64, error) {
	return 0, nil
}

func TestHandleCreateUser(t *testing.T) {
	user := User{
		Id:       "233",
		Name:     "frank",
		Age:      32,
		Password: "12344",
	}
	mockDatabaseQuery := MockDatabaseQuery{}
	t.Parallel()
	t.Run("should respond with 201 create status", func(t *testing.T) {
		var buf bytes.Buffer
		err := json.NewEncoder(&buf).Encode(user)
		if err != nil {
			t.Fatal(err)
		}
		handler := handleCreateUser(&mockDatabaseQuery)
		req := httptest.NewRequest("GET", "http://create-user", &buf)
		req.Header.Set("Content-type", "application/json")
		responseRecoder := httptest.NewRecorder()
		handler.ServeHTTP(responseRecoder, req)
		expectedStatusCode := responseRecoder.Result().StatusCode

		if expectedStatusCode != http.StatusCreated {
			t.Fatalf("got: %d want: %d", expectedStatusCode, http.StatusCreated)
		}
	})
}
