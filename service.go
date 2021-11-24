package champiris

import (
	"errors"
	"time"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"

	stdContext "context"
)

type Service struct {
	App    *iris.Application
	Config *NetConfig
}

type RoutesFunc func(m *mvc.Application)

func (service *Service) Default() error {
	return service.New(&NetConfig{
		Protocol: "tcp4",
		Host:     "0.0.0.0",
		Port:     "8080",
	})
}

func (service *Service) New(config *NetConfig) error {
	service.App = iris.New()

	if config == nil {
		return errors.New("network port not set")
	}

	service.Config = config

	//requestLog, loggerClose := service.newRequestLogger()
	//service.App.Use(requestLog)
	//service.App.OnAnyErrorCode(requestLog, func(ctx iris.Context) {
	//	ctx.Values().Set("logger_message", "a dynamic message passed to the logs")
	//})

	iris.RegisterOnInterrupt(func() {
		//if err := loggerClose(); err != nil {
		//	service.App.Logger().Fatal(err)
		//}

		timeout := 5 * time.Second
		ctx, cancel := stdContext.WithTimeout(stdContext.Background(), timeout)
		defer cancel()
		//close all hosts
		service.App.Shutdown(ctx)
	})

	return nil
}

func (service *Service) AddRoute(party string, routesFunc RoutesFunc) {
	p := service.App.Party(party)
	p.Done(setLog)
	mvc.Configure(p, routesFunc)
}

func setLog(ctx iris.Context) {
	//TODO 收集log
}

func (service *Service) Run() error {
	err := service.App.Run(
		iris.Addr(service.Config.Host+":"+service.Config.Port),
		iris.WithOptimizations,
		iris.WithPathEscape,
		iris.WithoutServerError(iris.ErrServerClosed),
	)
	return err
}

func (service *Service) Interrupt() error {
	err := service.App.Shutdown(stdContext.Background())
	return err
}
