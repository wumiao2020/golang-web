package controllers

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"wumiao/app/services"
)

type PageController struct {
	Service services.Page
}

func (p *PageController) Get() mvc.Result {
	data := p.Service.Get("index.html")
	if data == nil {
		return mvc.View{
			Code:   iris.StatusNotFound,
			Name:   "errors/404.html",
			Layout: iris.NoLayout,
			Data: iris.Map{
				"title": "你很神，找到了不存在的页面",
			},
		}
	}
	return mvc.View{
		Name: "page.html",
		Data: iris.Map{
			"title": data.Title,
			"data":  data,
		},
	}
}

func (p *PageController) GetBy(page string) mvc.Result {
	data := p.Service.Get(page)
	if data == nil {
		return mvc.View{
			Code:   iris.StatusNotFound,
			Name:   "errors/404.html",
			Layout: iris.NoLayout,
			Data: iris.Map{
				"title": "你很神，找到了不存在的页面",
			},
		}
	}
	return mvc.View{
		Name: "page.html",
		Data: iris.Map{
			"title": data.Title,
			"data":  data,
		},
	}
}
