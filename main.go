package main

import (
	"log"
	"net/http"

	restful "github.com/emicklei/go-restful"
	controller "github.com/xd-meal/xd-meal-backend/controller"
	DB "github.com/xd-meal/xd-meal-backend/dao"
)

func main() {
	DB.InitDB()
	// 初始化 db
	wsContainer := restful.NewContainer()
	m := controller.MealResource{}
	m.RegisterTo(wsContainer)

	// Add container filter to enable CORS
	cors := restful.CrossOriginResourceSharing{
		ExposeHeaders:  []string{"X-My-Header"},
		AllowedHeaders: []string{"Content-Type", "Accept"},
		AllowedMethods: []string{"GET", "POST"},
		CookiesAllowed: false,
		Container:      wsContainer}
	wsContainer.Filter(cors.Filter)

	// Add container filter to respond to OPTIONS
	wsContainer.Filter(wsContainer.OPTIONSFilter)

	log.Print("start listening on localhost:8899")
	server := &http.Server{Addr: ":8899", Handler: wsContainer}
	log.Fatal(server.ListenAndServe())
}
