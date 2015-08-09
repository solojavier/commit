package controllers

import "github.com/revel/revel"

type App struct {
	*revel.Controller
}

func (c App) Index(user string) revel.Result {
	return c.Render(user)
}
