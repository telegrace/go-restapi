package main

import (
	"testing"
	"net/http"
	"net/http/httptest"
)

func TestAllArticles(t *testing.T) {
	// pointer to testing package
	req, err := http.NewRequest("GET", "/articles", nil)
	if err != nil {
		t.Fatal(err)
	}
	// creates a new recorder to record the response by / endpoint
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(allArticles)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `[{"title":"Test Title","desc":"Test Description","content":"Test Content"}]`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body \ngot %v \nwant %v",
			rr.Body.String(), expected)
	}
}