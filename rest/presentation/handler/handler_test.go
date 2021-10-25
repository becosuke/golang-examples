package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"

	"github.com/becosuke/golang-examples/rest/domain/entity"
	"github.com/becosuke/golang-examples/rest/registry/config"
	"github.com/becosuke/golang-examples/rest/registry/injector"

	"github.com/stretchr/testify/assert"
)

var impl *handlerImpl

func TestMain(m *testing.M) {
	if err := setup(); err != nil {
		log.Fatal(err)
	}
	ret := m.Run()
	teardown()
	os.Exit(ret)
}

func setup() error {
	in := injector.NewInjector()
	impl = &handlerImpl{
		config:     in.InjectConfig(),
		controller: in.InjectController(),
	}
	return nil
}

func teardown() {
}

func TestGetConst(t *testing.T) {
	data := url.Values{config.KeyString: {config.ConstKey}}
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("%s?%s", config.Endpoint, data.Encode()), nil)
	w := httptest.NewRecorder()
	impl.route().ServeHTTP(w, req)

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
	impl.route().ServeHTTP(w, req)

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
	impl.route().ServeHTTP(w, req)

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
	impl.route().ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("w.Code: %d", w.Code)
	}
	t.Log(w.Body.String())

	_ = json.Unmarshal([]byte(w.Body.String()), &res)
	assert.Equal(t, "greeting", res.GetKey())
	assert.Equal(t, "hello", res.GetValue())
}
