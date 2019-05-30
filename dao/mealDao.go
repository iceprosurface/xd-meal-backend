package dao

import (
	"log"
)

func QueryAllMeals() (list []MealDao) {
	list = []MealDao{}
	rows, err := DB.Query("SELECT id as Id, name as Name, star as Star FROM meal LIMIT 15")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var id int
		var name string
		var star int
		var item = new(MealDao)
		if err := rows.Scan(&id, &name, &star); err != nil {
			log.Fatal(err)
		}
		item.Id = id
		item.Name = name
		item.Star = star
		list = append(list, *item)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	println(len(list))
	return
}

type MealDao struct {
	Id   int
	Name string
	Star int
}
