package elklogger

import (
	"bytes"
	"fmt"
	"io/ioutil"

	"git.championtek.com.tw/go/logger/v2"
	"github.com/kataras/iris/v12/context"
)

type Logger struct {
}

func New(cfg *logger.ELKConfig) context.Handler {
	l := &Logger{}

	loggerCfg := logger.Config{
		Host:             cfg.ELK.URL,
		Port:             cfg.ELK.Port,
		User:             cfg.ELK.User,
		Password:         cfg.ELK.Password,
		Index:            cfg.ELK.Index,
		NumberOfShards:   cfg.ELK.NumberOfShards,
		NumberOfReplicas: cfg.ELK.NumberOfReplicas,
	}
	_ = logger.Mgr.Init(&loggerCfg)

	return l.Serve
}

func (l *Logger) Serve(ctx context.Context) {
	port := ctx.RemoteAddr()
	fmt.Print(port + " ")
	fmt.Println(ctx.GetCurrentRoute())
	body, _ := ctx.GetBody()
	fmt.Println(string(body))
	ctx.Request().Body = ioutil.NopCloser(bytes.NewBuffer(body))
	ctx.Next()
}
