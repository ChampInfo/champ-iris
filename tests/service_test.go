package tests

import (
	"testing"

	"github.com/kataras/iris/v12"

	"github.com/kataras/iris/v12/mvc"

	"git.championtek.com.tw/go/champiris"
	"git.championtek.com.tw/go/champiris-contrib/graphql"
	"github.com/iris-contrib/middleware/cors"
)

var Ql *graphql.Graphql

func init() {
	Ql = graphql.Default()
	Ql.ShowPlayground = true
}

func TestApi_HelloWord(t *testing.T) {
	var service champiris.Service
	var config champiris.NetConfig
	config = champiris.NetConfig{
		Port: "8080"}
	service.New(&config)
	service.AddRoute(champiris.RouterSet{
		Party:  "/api",
		Router: apiRouter,
	})
	service.Run()
}

func apiRouter(app *mvc.Application) {
	app.Handle(new(ApiHandle))
}

type ApiHandle struct {
	Ctx iris.Context
}

//get localhost:port/api/hello
func (h *ApiHandle) GetHello() {
	h.Ctx.WriteString("HelloWord")
}

func TestAPI_NewService(t *testing.T) {
	var service champiris.Service

	_ = service.New(&champiris.NetConfig{
		Port: "8080"})

	service.App.Logger().SetLevel("debug")
	router := champiris.RouterSet{
		Party: "/service/v1",
		Router: func(m *mvc.Application) {
			m.Router.Use(cors.AllowAll())
			m.Handle(Ql)
		},
	}
	service.AddRoute(router)
	addSchema()
	if err := service.Run(); err != nil {
		t.Error(err)
	}
}
