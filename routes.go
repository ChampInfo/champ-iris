package champiris

import (
	"git.championtek.com.tw/go/champiris/middleware/elklogger"
	"git.championtek.com.tw/go/logger/v2"
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

var (
	crs = cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"HEAD", "POST", "GET", "PATCH", "DELETE"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Authorization"},
		AllowCredentials: true,
		Debug:            true,
	})

	elk = elklogger.New(&logger.ELKConfig{ELK:logger.ELKInfo{
		URL:              "http://52.196.196.142",
		Port:             "9200",
		Index:            "champ_iris",
		User:             "elastic",
		Password:         "work$t/6qup3",
		NumberOfShards:   1,
		NumberOfReplicas: 0,
	}})
)

func Routes(m *mvc.Application) {
	//crs := cors.New(cors.Options{
	//	AllowedOrigins:   []string{"*"},
	//	AllowedMethods:   []string{"HEAD", "POST", "GET", "PATCH", "DELETE"},
	//	AllowedHeaders:   []string{"*"},
	//	ExposedHeaders:   []string{"Authorization"},
	//	AllowCredentials: true,
	//	Debug:            true,
	//})
	m.Router.AllowMethods(iris.MethodOptions)
	m.Router.Use(crs)
	m.Handle(new(WebPage))
	m.Party("/query").Handle(new(Ql))
}

func RoutesWithLogger(m *mvc.Application) {
	m.Router.AllowMethods(iris.MethodOptions)
	m.Router.Use(crs, elk)
	m.Handle(new(WebPage))
	m.Party("/query").Handle(new(Ql))
}
