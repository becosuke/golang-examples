package controller

import (
	"net/http"

	"github.com/becosuke/golang-examples/rest/application/usecase"
	"github.com/becosuke/golang-examples/rest/registry/config"
)

type Controller interface {
	Post(http.ResponseWriter, *http.Request)
	Get(http.ResponseWriter, *http.Request)
	Put(http.ResponseWriter, *http.Request)
	Delete(http.ResponseWriter, *http.Request)
}

func NewController(config *config.Config, usecase usecase.Usecase) Controller {
	return &controllerImpl{
		config:  config,
		usecase: usecase,
	}
}

type controllerImpl struct {
	config  *config.Config
	usecase usecase.Usecase
}

func (impl *controllerImpl) Get(w http.ResponseWriter, r *http.Request) {
	res, err := impl.usecase.Read(r.Context(), r.FormValue(config.KeyString))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
	} else {
		_, _ = w.Write(res)
	}
}

func (impl *controllerImpl) Post(w http.ResponseWriter, r *http.Request) {
	res, err := impl.usecase.Create(r.Context(), r.FormValue(config.KeyString), r.FormValue(config.ValueString))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
	} else {
		_, _ = w.Write(res)
	}
}

func (impl *controllerImpl) Put(w http.ResponseWriter, r *http.Request) {
	res, err := impl.usecase.Update(r.Context(), r.FormValue(config.KeyString), r.FormValue(config.ValueString))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
	} else {
		_, _ = w.Write(res)
	}
}

func (impl *controllerImpl) Delete(w http.ResponseWriter, r *http.Request) {
	err := impl.usecase.Delete(r.Context(), r.FormValue(config.KeyString))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}
