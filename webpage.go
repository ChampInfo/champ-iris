package champiris

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type WebPage struct {
	ctx iris.Context
}

func (w *WebPage) Get() mvc.View {
	return mvc.View{
		Name: "graphql.html",
	}
}
