package controllers

import(
	"github.com/revel/revel"
)

type Pendidikan struct{
	App
}

func(c Pendidikan) Index() revel.Result{
	return c.Render()
}