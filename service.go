package champiris

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type Service struct {
	app      *iris.Application
	config   NetConfig
	version  []string
	htmlPath string
}

//Application get the iris application
func (service *Service) Application() *iris.Application {
	return service.app
}

//HtmlPath set html path
func (service *Service) HtmlPath(htmlPath string) {
	service.htmlPath = htmlPath
}

//Version set version
func (service *Service) Version(version []string) {
	service.version = version
}

func (service *Service) New(config NetConfig) error {
	service.app = iris.New()
	if len(service.version) == 0 { // set the default version
		service.version = []string{"1"}
	}
	if len(service.htmlPath) == 0 { // set the default html path
		service.htmlPath = "./https/web"
	}
	if config.Port == "" {
		return errors.New("network port not set")
	}
	service.config = config

	requestLog, loggerClose := service.newRequestLogger()
	service.app.Use(requestLog)
	service.app.OnAnyErrorCode(requestLog, func(ctx iris.Context) {
		ctx.Values().Set("logger_message", "a dynamic message passed to the logs")
	})
	iris.RegisterOnInterrupt(func() {
		if err := loggerClose(); err != nil {
			service.app.Logger().Fatal(err)
		}
	})
	service.registerStaticWebPages(service.htmlPath)
	service.setVersionRoutingPath(service.version)
	return nil
}

func (service *Service) Run() error {
	err := service.app.Run(
		iris.Addr(":"+service.config.Port),
		iris.WithOptimizations,
		iris.WithoutServerError(iris.ErrServerClosed),
	)
	return err
}

func (service *Service) registerStaticWebPages(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.MkdirAll(path, os.ModePerm)
		if err != nil {
			log.Fatal(fmt.Sprintf("Create %s error: ", path), err)
		}
	}
	service.app.RegisterView(iris.HTML(path, ".html").Reload(true))
}

func (service *Service) setVersionRoutingPath(versions []string) {
	for _, version := range versions {
		mvc.Configure(service.app.Party("/service/v"+version), Routes)
	}
}
