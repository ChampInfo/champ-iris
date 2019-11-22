package champiris

import (
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func Routes(m *mvc.Application) {
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"HEAD", "POST", "GET", "PATCH", "DELETE"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Authorization"},
		AllowCredentials: true,
		Debug:            true,
	})
	m.Router.AllowMethods(iris.MethodOptions)
	m.Router.Use(crs)
	m.Handle(Get)
	//m.Party("/query").Handle(new(apis.Query))
}

func Get() mvc.Result {
	return mvc.View{
		Name: "graphql.html",
	}
}
