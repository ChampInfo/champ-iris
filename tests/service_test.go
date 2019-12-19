package tests

import (
	"testing"

	"github.com/kataras/iris/v12/mvc"

	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12/context"

	"git.championtek.com.tw/go/champiris"
)

func TestAPI_NewService(t *testing.T) {
	var service champiris.Service

	_ = service.New(&champiris.NetConfig{
		Port:         "8080",
		LoggerEnable: true,
		JWTEnable:    true,
	})
	service.App.Logger().SetLevel("debug")
	router := champiris.RouterSet{
		Party: "/service/v1",
		Router: func(m *mvc.Application) {
			m.Party("/query").Handle(new(Ql))
			m.Handle(new(WebPage))
		},
		Middleware: []context.Handler{cors.AllowAll()},
	}
	service.AddRoute(router)
	addSchema()
	if err := service.Run(); err != nil {
		t.Error(err)
	}
}
