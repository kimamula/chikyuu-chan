package controllers

import (
	"github.com/revel/revel"
)

type Map struct {
	Application
}

func (c Map) Index() revel.Result {
	return c.Render()
}
