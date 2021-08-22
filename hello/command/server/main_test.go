package main

import (
	"encoding/json"
	"github.com/becosuke/golang-examples/hello/domain/entity"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetHello(t *testing.T) {
	req, _ := http.NewRequest("GET", "/hello", nil)
	w := httptest.NewRecorder()
	route().ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("w.Code: %d", w.Code)
	}

	var res entity.Hello
	_ = json.Unmarshal([]byte(w.Body.String()), &res)
	if res.Message != "hello" {
		t.Errorf("res.Message: %s", res.Message)
	}
}
