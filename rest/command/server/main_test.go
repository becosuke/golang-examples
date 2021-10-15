package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/becosuke/golang-examples/rest/domain/entity"
	"github.com/becosuke/golang-examples/rest/registry/config"
)

func TestGetConst(t *testing.T) {
	data := url.Values{config.KeyString: {config.ConstKey}}
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("%s?%s", config.Endpoint, data.Encode()), nil)
	w := httptest.NewRecorder()
	s.Route().ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("w.Code: %d", w.Code)
	}
	t.Log(w.Body.String())

	var res entity.Entity
	_ = json.Unmarshal([]byte(w.Body.String()), &res)
	assert.Equal(t, config.ConstKey, res.GetKey())
	assert.Equal(t, config.ConstValue, res.GetValue())
}

func TestPostConst(t *testing.T) {
	data := url.Values{config.KeyString: {config.ConstKey}}
	req, _ := http.NewRequest(http.MethodPost, config.Endpoint, strings.NewReader(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	s.Route().ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("w.Code: %d", w.Code)
	}
	t.Log(w.Body.String())

	var res entity.Entity
	_ = json.Unmarshal([]byte(w.Body.String()), &res)
	assert.Equal(t, config.ConstKey, res.GetKey())
	assert.Equal(t, config.ConstValue, res.GetValue())
}

func TestPostAndGet(t *testing.T) {
	data := url.Values{config.KeyString: {"greeting"}, config.ValueString: {"hello"}}
	req, _ := http.NewRequest(http.MethodPost, config.Endpoint, strings.NewReader(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	s.Route().ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("w.Code: %d", w.Code)
	}
	t.Log(w.Body.String())

	var res entity.Entity
	_ = json.Unmarshal([]byte(w.Body.String()), &res)
	assert.Equal(t, "greeting", res.GetKey())
	assert.Equal(t, "hello", res.GetValue())

	data = url.Values{config.KeyString: {"greeting"}}
	req, _ = http.NewRequest(http.MethodGet, fmt.Sprintf("%s?%s", config.Endpoint, data.Encode()), strings.NewReader(data.Encode()))
	w = httptest.NewRecorder()
	s.Route().ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("w.Code: %d", w.Code)
	}
	t.Log(w.Body.String())

	_ = json.Unmarshal([]byte(w.Body.String()), &res)
	assert.Equal(t, "greeting", res.GetKey())
	assert.Equal(t, "hello", res.GetValue())
}
