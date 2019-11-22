package champiris

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

var irisConfig IrisNetWork

type API struct {
	app      *iris.Application
	version  []string
	htmlPath string
}

func (api *API) Application() *iris.Application {
	return api.app
}

func (api *API) SetHtmlPath(htmlPath string) {
	api.htmlPath = htmlPath
}

func (api *API) SetVersion(version []string) {
	api.version = version
}

func (api *API) NewService() {
	api.app = iris.New()
	if len(api.version) == 0 {
		api.version = []string{"1"}
	}
	if len(api.htmlPath) == 0 {
		api.htmlPath = "./https/web"
	}
	requestLog, loggerClose := api.newRequestLogger()
	api.app.Use(requestLog)
	api.app.OnAnyErrorCode(requestLog, func(ctx iris.Context) {
		ctx.Values().Set("logger_message", "a dynamic message passed to the logs")
	})
	iris.RegisterOnInterrupt(func() {
		if err := loggerClose(); err != nil {
			api.app.Logger().Fatal(err)
		}
	})
	api.addHtmlDirectory(api.htmlPath)
	api.setApiVersion(api.version)
}

func (api *API) Configuration(config IrisNetWork) {
	irisConfig = config
}

func (api *API) addHtmlDirectory(path string) {
	api.app.RegisterView(iris.HTML(path, ".html").Reload(true))
}

func (api *API) setApiVersion(v []string) {
	for _, version := range v {
		mvc.Configure(api.app.Party("/api/v"+version), Routes)
	}
}

func (api *API) Run() error {
	err := api.app.Run(
		iris.Addr(":"+irisConfig.Port),
		iris.WithOptimizations,
		iris.WithoutServerError(iris.ErrServerClosed),
	)
	return err
}
