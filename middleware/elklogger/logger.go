package elklogger

import (
	"git.championtek.com.tw/go/logger/v2"
	"github.com/kataras/iris/v12/context"
	"github.com/olivere/elastic/v7"
)

type Logger struct {
	elkClient *elastic.Client
}

func New(cfg *logger.ELKConfig) context.Handler {
	l := &Logger{}
	
	loggerCfg := logger.Config{
		Host:     cfg.ELK.URL,
		Port:     cfg.ELK.Port,
		User:     cfg.ELK.User,
		Password: cfg.ELK.Password,
		Index:    cfg.ELK.Index,
		NumberOfShards: cfg.ELK.NumberOfShards,
		NumberOfReplicas: cfg.ELK.NumberOfReplicas,
	}
	_ = logger.Mgr.Init(&loggerCfg)

	l.elkClient = logger.Mgr.Client()

	return l.Serve
}

func (l *Logger) Serve(ctx context.Context) {
	ctx.Next()
}
