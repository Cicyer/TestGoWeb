package main

import"github.com/kataras/iris"

func main(){
	app:=iris.New()
	app.RegisterView(iris.HTML("./views",".html"))
	app.Handle("GET","/", func(ctx iris.Context) {
		ctx.ViewData("message","Hello world!")
		ctx.View("index.html")
	})

	app.Run(iris.Addr(":8080"))

}
