package controllers

import "github.com/kataras/iris"

func Login(ctx iris.Context) {
	ctx.ViewData("title","登录")
	ctx.ViewLayout(iris.NoLayout)
	err := ctx.View("customer/login.html")
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.Writef(err.Error())
	}

}