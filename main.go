package main

import (
	"os"

	"github.com/kataras/iris/v12"
	Router "github.com/seifer-d/app_crud/router"
)

func main() {
	app := iris.New()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081" //localhost
	}
	Router.AddRutas(app)

	tmpl := iris.HTML("./views", ".html")

	tmpl.Reload(true)

	tmpl.AddFunc("greet", func(s string) string {
		return "Greetings " + s + "!"
	})
	app.RegisterView(tmpl)
	app.Get("/", func(ctx iris.Context) {
		//ctx.ViewData("message", "Crud de estudiante")
		ctx.View("index.html")
	})

	app.Listen(":" + port)
}
