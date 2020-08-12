package controllers

import "github.com/kataras/iris"

func Register(ctx iris.Context) {
	ctx.ViewData("title","注册")
	ctx.ViewLayout(iris.NoLayout)
	err := ctx.View("customer/register.html")
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.Writef(err.Error())
	}
}