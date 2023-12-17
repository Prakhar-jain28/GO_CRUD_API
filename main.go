package main

import (
	"GO-LANG/datastore"
	"GO-LANG/handler"

	"gofr.dev/pkg/gofr"
) 

func main() {
	app := gofr.New()

	s := datastore.New()
	h := handler.New(*s)

	app.GET("/blog/{ID}", h.GetByID)   
	app.POST("/blog", h.Create)
	app.DELETE("/blog/{ID}", h.Delete) 
	app.PUT("/blog/{ID}", h.Update)    

	app.Server.HTTP.Port = 9092
	app.Start()
}
