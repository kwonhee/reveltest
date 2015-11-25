package controllers

import "github.com/revel/revel"

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) Me() revel.Result {
	name := "kkh"
	return c.Render(name)
}

func (c App) Menumber(id int) revel.Result {
	
	return c.Render(id)
}

func (c App) Mepost() revel.Result {
	
	return c.Render()
}

func (c App) Mesetpost() revel.Result {
	password := c.Params.Get("password")
	if password == "1234" {
		return c.Redirect(App.Index)
	}
	return c.Redirect(App.Mepost)
}