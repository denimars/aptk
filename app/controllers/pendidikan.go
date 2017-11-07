package controllers

import(
	"github.com/revel/revel"
	"strings"
	m "aptk/app/models"
	"aptk/app/routes"
	"aptk/app"
	"log"
)

type Pendidikan struct{
	App
}

func(c Pendidikan) Index() revel.Result{
	var pendidikan []m.Pendidikan
	err := app.DB.Find(&pendidikan)
	if err.Error != nil{
		panic(err.Error)
	}
	log.Println(pendidikan)
	return c.Render(pendidikan)
}

func(c Pendidikan) Save(list string) revel.Result{
	var pendidikan m.Pendidikan
	pendidikan.Id = 0
	pendidikan.List = strings.ToUpper(list)
	err := app.DB.Save(&pendidikan)
	if err.Error!= nil{
		log.Println(err)
		panic(err)
	}
	return c.Redirect(routes.Pendidikan.Index())
}