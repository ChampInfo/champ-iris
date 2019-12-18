package elklogger

import (
	"encoding/json"
	"git.championtek.com.tw/go/logger/v2"
	"github.com/kataras/iris/v12/context"
	"github.com/olivere/elastic/v7"
)

type Logger struct {
	elkClient *elastic.Client
}

func New(cfg *logger.ELKConfig) context.Handler {
	l := &Logger{}

	elkMap := cfg.ELK.Mapping
	
	settings := logger.Settings{
		NumberOfShards:   elkMap.Settings.NumberOfShards,
		NumberOfReplicas: elkMap.Settings.NumberOfReplicas,
	}
	
	properties := logger.Properties{
		Service: elkMap.Mappings.Properties.Service,
		IP:      elkMap.Mappings.Properties.IP,
		Status:  elkMap.Mappings.Properties.Status,
		Method:  elkMap.Mappings.Properties.Method,
		Path:    elkMap.Mappings.Properties.Path,
		Tags:    elkMap.Mappings.Properties.Tags,
		Created: elkMap.Mappings.Properties.Created,
		Remark:  elkMap.Mappings.Properties.Remark,
	}
	
	mappings := logger.Mappings{Properties:properties}
	
	elkCfg := logger.Mapping{
		Settings: settings,
		Mappings: mappings,
	}
	
	b, _ := json.Marshal(elkCfg)
	//if err != nil {
	//	return nil, err
	//}
	
	lcfg := logger.Config{
		Host:     cfg.ELK.Host.URL,
		Port:     cfg.ELK.Host.PORT,
		User:     cfg.ELK.BasicAuth.User,
		Password: cfg.ELK.BasicAuth.Password,
		Index:    cfg.ELK.Index,
		Mapping:  string(b),
	}
	
	//if err := logger.Mgr.Init(&lcfg); err != nil {
	//	return nil, err
	//}
	_ = logger.Mgr.Init(&lcfg)

	l.elkClient = logger.Mgr.Client()

	return l.Serve
}

func (l *Logger) Serve(ctx context.Context) {

	ctx.Next()
}
