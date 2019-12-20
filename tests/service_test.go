package tests

import (
	"testing"

	"git.championtek.com.tw/go/champiris/middleware/elklogger"
	"git.championtek.com.tw/go/logger/v2"

	"github.com/kataras/iris/v12/mvc"

	"github.com/iris-contrib/middleware/cors"

	"git.championtek.com.tw/go/champiris"
)

func TestAPI_NewService(t *testing.T) {
	var service champiris.Service

	_ = service.New(&champiris.NetConfig{
		Port:         "8080",
		LoggerEnable: true,
		JWTEnable:    true,
	})

	// if LoggerEnable is true, then setup logger
	elk := elklogger.New(&logger.ELKConfig{ELK: logger.ELKInfo{
		URL:              "http://52.196.196.142",
		Port:             "9200",
		Index:            "champ_iris",
		User:             "elastic",
		Password:         "work$t/6qup3",
		NumberOfShards:   1,
		NumberOfReplicas: 0,
	}})

	service.App.Logger().SetLevel("debug")
	router := champiris.RouterSet{
		Party: "/service/v1",
		Router: func(m *mvc.Application) {
			m.Router.Use(cors.AllowAll())
			m.Router.Use(elk)
			m.Party("/query").Handle(new(Ql))
			m.Handle(new(WebPage))
		},
	}
	service.AddRoute(router)
	addSchema()
	if err := service.Run(); err != nil {
		t.Error(err)
	}
}
