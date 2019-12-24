package tests

import (
	"testing"

	"git.championtek.com.tw/go/champiris/middleware/jwtpassport"
	"git.championtek.com.tw/go/passport"
	"github.com/kataras/iris/v12/context"

	"git.championtek.com.tw/go/champiris/middleware/elklogger"
	"git.championtek.com.tw/go/logger/v2"

	"github.com/kataras/iris/v12/mvc"

	"git.championtek.com.tw/go/champiris"
	"git.championtek.com.tw/go/champiris-contrib/graphql"
	"github.com/iris-contrib/middleware/cors"
)

var Ql graphql.Graphql

func init() {
	Ql = graphql.Graphql{
		ShowPlayground: true,
	}
	Ql.Query.New("Query", "搜尋&取得資料的相關命令")
	Ql.Mutation.New("Mutation", "主要用在建立、修改、刪除的相關命令")
}

func TestAPI_NewService(t *testing.T) {
	var service champiris.Service
	var elk context.Handler
	var psp context.Handler

	_ = service.New(&champiris.NetConfig{
		Port:         "8080",
		LoggerEnable: false,
		JWTEnable:    false,
	})

	if service.Config.LoggerEnable {
		// if LoggerEnable is true, then setup logger
		elk = elklogger.New(&logger.ELKConfig{ELK: logger.ELKInfo{
			URL:              "http://52.196.196.142",
			Port:             "9200",
			Type:             0,
			Index:            "champ_iris",
			User:             "elastic",
			Password:         "work$t/6qup3",
			NumberOfShards:   1,
			NumberOfReplicas: 0,
		}})
	}

	if service.Config.JWTEnable {
		// if JWTEnable is true, then setup jwt
		psp = jwtpassport.New(&passport.Config{
			Secret: "champ",
			Claims: passport.PublicClaims{
				Audience:    "bskini",
				Subject:     "champiris",
				Email:       "service@dishrank.com",
				PhoneNumber: "",
				Issuer:      "drc",
				Duration:    3600,
			},
		})
	}

	service.App.Logger().SetLevel("debug")
	router := champiris.RouterSet{
		Party: "/service/v1",
		Router: func(m *mvc.Application) {
			m.Router.Use(cors.AllowAll())
			if elk != nil {
				m.Router.Use(elk)
			}
			if psp != nil {
				m.Router.Use(psp)
			}
			m.Handle(&Ql)
		},
	}
	service.AddRoute(router)
	addSchema()
	if err := service.Run(); err != nil {
		t.Error(err)
	}
}
