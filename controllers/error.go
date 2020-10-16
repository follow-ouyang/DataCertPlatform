package controllers

import "github.com/astaxie/beego"

type Error struct {
	beego.Controller
}

func (m *Error) Get() {
	m.TplName = "error.html"
}