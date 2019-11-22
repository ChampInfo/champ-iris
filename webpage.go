package champiris

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type webPage struct {
	ctx iris.Context
}

func (w *webPage) Get() mvc.View {
	return mvc.View{
		Name: "graphql.html",
	}
}
