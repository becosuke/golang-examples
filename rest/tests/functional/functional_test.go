package functional

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/becosuke/golang-examples/rest/internal/entities"
	"github.com/becosuke/golang-examples/rest/internal/registry/config"
	"github.com/becosuke/golang-examples/rest/internal/registry/injector"
)

func TestPost(t *testing.T) {
	data := url.Values{config.KeyString: {"kkk"}, config.ValueString: {"vvv"}}
	r, _ := http.NewRequest(http.MethodPost, config.Endpoint, strings.NewReader(data.Encode()))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()

	in := injector.NewInjector()
	in.InjectHandler().Route().ServeHTTP(w, r)

	assert.Equal(t, http.StatusCreated, w.Code)

	var res entities.Node
	err := json.Unmarshal([]byte(w.Body.String()), &res)
	require.NoError(t, err)
	node, err := in.InjectRepository().Read(context.Background(), entities.NewNodeKey("kkk"))
	require.NoError(t, err)
	assert.Equal(t, node.GetKey(), res.GetKey())
	assert.Equal(t, node.GetValue(), res.GetValue())
}

func TestGet(t *testing.T) {
	data := url.Values{config.KeyString: {"kkk"}}
	r, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("%s?%s", config.Endpoint, data.Encode()), nil)
	w := httptest.NewRecorder()

	in := injector.NewInjector()
	node, err := in.InjectRepository().Create(context.Background(), entities.NewNode("kkk", "vvv"))
	require.NoError(t, err)
	in.InjectHandler().Route().ServeHTTP(w, r)

	assert.Equal(t, http.StatusOK, w.Code)

	var res entities.Node
	err = json.Unmarshal([]byte(w.Body.String()), &res)
	require.NoError(t, err)
	assert.Equal(t, node.GetKey(), res.GetKey())
	assert.Equal(t, node.GetValue(), res.GetValue())
}

func TestPut(t *testing.T) {
	data := url.Values{config.KeyString: {"kkk"}, config.ValueString: {"vvv"}}
	r, _ := http.NewRequest(http.MethodPut, config.Endpoint, strings.NewReader(data.Encode()))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()

	in := injector.NewInjector()
	in.InjectHandler().Route().ServeHTTP(w, r)

	assert.Equal(t, http.StatusOK, w.Code)

	var res entities.Node
	err := json.Unmarshal([]byte(w.Body.String()), &res)
	require.NoError(t, err)
	node, err := in.InjectRepository().Read(context.Background(), entities.NewNodeKey("kkk"))
	require.NoError(t, err)
	assert.Equal(t, node.GetKey(), res.GetKey())
	assert.Equal(t, node.GetValue(), res.GetValue())
}

func TestDelete(t *testing.T) {
	data := url.Values{config.KeyString: {"kkk"}}
	r, _ := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s?%s", config.Endpoint, data.Encode()), nil)
	w := httptest.NewRecorder()

	in := injector.NewInjector()
	in.InjectHandler().Route().ServeHTTP(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "", w.Body.String())
}
