package bootstrap

import (
	"gin-demo/pkg/config"
	"gin-demo/pkg/database"
	"gin-demo/pkg/html"
	"gin-demo/pkg/routing"
	"gin-demo/pkg/static"
)

func Serve() {
	config.Set()
	database.Connect()
	routing.Init()
	static.LoadStatic(routing.GetRouter())
	html.LoadHtml(routing.GetRouter())
	routing.RegisterRoutes()
	routing.Serve()
}
