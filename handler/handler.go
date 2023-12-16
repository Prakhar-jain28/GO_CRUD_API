package handler

import (
	"gofr.dev/pkg/errors"
	"gofr.dev/pkg/gofr"

	"GO-LANG/datastore"
	"GO-LANG/model"
)

type handler struct {  
	store datastore.Student
}

func New(s datastore.Blog) handler {
	return handler{store: s}
}

func (h handler) Create(ctx *gofr.Context) (interface{}, error) {
	var blog model.Blog

	if err := ctx.Bind(&blog); err != nil {
		ctx.Logger.Errorf("error in binding: %v", err)
		return nil, errors.InvalidParam{Param: []string{"body"}}
	}

	resp, err := h.store.Create(ctx, &blog)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
