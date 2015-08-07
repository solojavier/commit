package controllers

import "github.com/revel/revel"

type Commitment struct {
	*revel.Controller
}

func (c Commitment) Create(user string) revel.Result {
	response := make(map[string]string)
	response["user"] = user

	return c.RenderJson(response)
}

func (c Commitment) Get(user string) revel.Result {
	response := make(map[string]string)
	response["user"] = user

	return c.RenderJson(response)
}
