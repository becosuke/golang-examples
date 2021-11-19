package controllers

import (
	"encoding/json"
	"github.com/becosuke/golang-examples/grpc/internal/entities/node"
	"net/http"

	"github.com/becosuke/golang-examples/grpc/internal/registry/config"
	"github.com/becosuke/golang-examples/grpc/internal/usecases"
)

type Controller interface {
	Post(http.ResponseWriter, *http.Request)
	Get(http.ResponseWriter, *http.Request)
	Put(http.ResponseWriter, *http.Request)
	Delete(http.ResponseWriter, *http.Request)
}

func NewController(config *config.Config, interactor usecases.Interactor) Controller {
	return &controllerImpl{
		config:     config,
		interactor: interactor,
	}
}

type controllerImpl struct {
	config     *config.Config
	interactor usecases.Interactor
}

func (impl *controllerImpl) Get(w http.ResponseWriter, r *http.Request) {
	params := node.NewKey(r.FormValue(config.KeyString))
	node, err := impl.interactor.Read(r.Context(), params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	res, err := json.Marshal(node)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(res)
}

func (impl *controllerImpl) Post(w http.ResponseWriter, r *http.Request) {
	params := node.NewKeyValueNode(r.FormValue(config.KeyString), r.FormValue(config.ValueString))
	node, err := impl.interactor.Create(r.Context(), params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	res, err := json.Marshal(node)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Location", r.Host+r.URL.Path+node.Key)
	_, _ = w.Write(res)
}

func (impl *controllerImpl) Put(w http.ResponseWriter, r *http.Request) {
	params := node.NewKeyValueNode(r.FormValue(config.KeyString), r.FormValue(config.ValueString))
	node, err := impl.interactor.Update(r.Context(), params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	res, err := json.Marshal(node)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(res)
}

func (impl *controllerImpl) Delete(w http.ResponseWriter, r *http.Request) {
	params := node.NewKey(r.FormValue(config.KeyString))
	err := impl.interactor.Delete(r.Context(), params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusOK)
}
