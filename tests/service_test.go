package tests

import (
	"testing"

	"github.com/kataras/iris/v12"

	"github.com/kataras/iris/v12/mvc"

	"git.championtek.com.tw/go/champiris"
)

func TestApi_HelloWord(t *testing.T) {
	var service champiris.Service

	service.Default()
	service.AddRoute("/api", apiRouter)
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
	h.Ctx.Next()
}
