package controllers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/parfaire/marvelapi/models"
)

type Action func(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) *models.ApiError

type Controller struct{}

func (c *Controller) Perform(action Action) httprouter.Handle {
	return httprouter.Handle(func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		err := action(writer, request, params)
		if err != nil {
			http.Error(writer, err.Error(), err.HTTPStatus)
		}
	})
}
