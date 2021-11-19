package controller

import (
	"encoding/json"
	"github.com/becosuke/golang-examples/kvstore/internal/entity/kv"
	"github.com/becosuke/golang-examples/kvstore/internal/registry/config"
	"github.com/becosuke/golang-examples/kvstore/internal/usecase"
	"net/http"
)

type Controller interface {
	Post(http.ResponseWriter, *http.Request)
	Get(http.ResponseWriter, *http.Request)
	Put(http.ResponseWriter, *http.Request)
	Delete(http.ResponseWriter, *http.Request)
}

func NewController(config *config.Config, interactor usecase.Interactor) Controller {
	return &controllerImpl{
		config:     config,
		interactor: interactor,
	}
}

type controllerImpl struct {
	config     *config.Config
	interactor usecase.Interactor
}

func (impl *controllerImpl) Get(w http.ResponseWriter, r *http.Request) {
	params := kv.NewSeal(r.FormValue(config.KeyString))
	p, err := impl.interactor.Read(r.Context(), params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	res, err := json.Marshal(p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(res)
}

func (impl *controllerImpl) Post(w http.ResponseWriter, r *http.Request) {
	params := kv.NewPack(r.FormValue(config.KeyString), r.FormValue(config.ValueString))
	pack, err := impl.interactor.Create(r.Context(), params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	res, err := json.Marshal(pack)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Location", r.URL.Host+r.URL.Path+pack.Key)
	_, _ = w.Write(res)
}

func (impl *controllerImpl) Put(w http.ResponseWriter, r *http.Request) {
	params := kv.NewPack(r.FormValue(config.KeyString), r.FormValue(config.ValueString))
	pack, err := impl.interactor.Update(r.Context(), params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	res, err := json.Marshal(pack)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(res)
}

func (impl *controllerImpl) Delete(w http.ResponseWriter, r *http.Request) {
	params := kv.NewSeal(r.FormValue(config.KeyString))
	err := impl.interactor.Delete(r.Context(), params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusOK)
}
