package main

import (
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"

	"github.com/merakiVE/CEHDUN/src/controllers"
)

var (
	PORT_SERVER = ":8002"
)

func main() {

	///Iris
	app := iris.New()

	//app.Configure(iris.WithConfiguration(iris.YAML("./config_iris.yml")))

	APILogger := logger.New(logger.Config{
		// Status displays status code
		Status: true,
		// IP displays request's remote address
		IP: true,
		// Method displays the http method
		Method: true,
		// Path displays the request path
		Path: true,
	})

	app.Use(APILogger)

	//Enable cors in server
	app.WrapRouter(cors.WrapNext(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"OPTIONS", "GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
		Debug:            true,
	}))

	app.Post("/connect", controllers.ConnectDB)
	app.Post("/build", controllers.BuildAPI)

	app.Run(iris.Addr(PORT_SERVER), iris.WithCharset("UTF-8"))
}
