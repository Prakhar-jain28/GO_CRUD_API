package handler

import (
	"strconv"

	"gofr.dev/pkg/errors"
	"gofr.dev/pkg/gofr"

	"GO-LANG/datastore"
	"GO-LANG/model"
)

type handler struct {  
	store datastore.Blog
}

func New(s datastore.Blog) handler {
	return handler{store: s}
}

func validateID(id string) (int, error) {
	res, err := strconv.Atoi(id)
	if err != nil {
		return 0, err
	}

	return res, err
}

func (h handler) GetByID(ctx *gofr.Context) (interface{}, error) {
	ID := ctx.PathParam("ID")
	if ID == "" {
		return nil, errors.MissingParam{Param: []string{"id"}}
	}

	if _, err := validateID(ID); err != nil {
		return nil, errors.InvalidParam{Param: []string{"id"}}
	}

	resp, err := h.store.GetByID(ctx, ID)
	if err != nil {
		return nil, errors.EntityNotFound{
			Entity: "blog",
			ID:     ID,
		}
	}

	return resp, nil
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

func (h handler) Update(ctx *gofr.Context) (interface{}, error) {
	i := ctx.PathParam("ID")
	if i == "" {
		return nil, errors.MissingParam{Param: []string{"id"}}
	}

	id, err := validateID(i)
	if err != nil {
		return nil, errors.InvalidParam{Param: []string{"id"}}
	}

	var blog model.Blog
	if err = ctx.Bind(&blog); err != nil {
		ctx.Logger.Errorf("error in binding: %v", err)
		return nil, errors.InvalidParam{Param: []string{"body"}}
	}

	blog.ID = id

	resp, err := h.store.Update(ctx, &blog)
	if err != nil {
		return nil, err
	}

	return resp, nil
}


func (h handler) Delete(ctx *gofr.Context) (interface{}, error) {
	i := ctx.PathParam("ID")
	if i == "" {
		return nil, errors.MissingParam{Param: []string{"id"}}
	}

	id, err := validateID(i)
	if err != nil {
		return nil, errors.InvalidParam{Param: []string{"id"}}
	}

	if err := h.store.Delete(ctx, id); err != nil {
		return nil, err
	}

	return "Deleted successfully", nil
}