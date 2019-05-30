package controller

import (
	"strconv"

	restful "github.com/emicklei/go-restful"
	DB "github.com/xd-meal/xd-meal-backend/dao"
)

type MealResource struct{}

type Meal struct {
	ID   string `json:"id" description:"id for Meal"`
	Name string `json:"name" description:"readable meal name"`
	Star int    `json:"start" description:"current star on meal"`
}

func (m MealResource) RegisterTo(container *restful.Container) {
	ws := new(restful.WebService)
	ws.
		Path("/meals").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)
	ws.Route(ws.GET("/meals").To(m.buildStrut))

	container.Add(ws)
}

func (m MealResource) buildStrut(request *restful.Request, response *restful.Response) {
	list := []Meal{}
	meals := DB.QueryAllMeals()
	for _, current := range meals {
		item := new(Meal)
		item.ID = strconv.Itoa(current.Id)
		item.Name = current.Name
		item.Star = current.Star
		list = append(list, *item)
	}

	response.WriteEntity(list)
}
