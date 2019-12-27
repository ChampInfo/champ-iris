package elklogger

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"time"

	"git.championtek.com.tw/go/logger/v2"
	"github.com/kataras/iris/v12/context"
)

type Logger struct {
	logType int
}

func New(cfg *logger.ELKConfig) context.Handler {
	l := &Logger{}

	l.logType = cfg.ELK.Type

	loggerCfg := logger.Config{
		Host:             cfg.ELK.URL,
		Port:             cfg.ELK.Port,
		Type:             cfg.ELK.Type,
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
	body, _ := ctx.GetBody()
	if l.logType == logger.GraphqlTemplate {
		graphqlBody := &GraphqlBody{}
		json.Unmarshal(body, graphqlBody)
		if len(graphqlBody.Query) != 0 {
			logData := logger.GraphqlService{
				IP:      port,
				Request: ctx.GetCurrentRoute().Name(),
				Result:  graphqlBody.Query + " - " + graphqlBody.Variables + " - " + graphqlBody.OperationName,
				Created: time.Now(),
				Tags:    nil,
				Remark:  "",
			}
			if err := logger.Mgr.PutLog(&logData); err != nil {
				log.Panicln(err)
			}
		}
	} else { //Restful request

	}

	ctx.Request().Body = ioutil.NopCloser(bytes.NewBuffer(body))
	ctx.Next()
}
