package tests

import (
	"git.championtek.com.tw/go/champiris"
	"testing"
)

func TestAPI_NewService(t *testing.T) {
	var service champiris.Service
	_ = service.New(&champiris.NetConfig{
		Port: "8080",
		LoggerEnable: true,
		JWTEnable: true,
	})

	service.App.Logger().SetLevel("debug")
	addSchema()
	if err := service.Run(); err != nil {
		t.Error(err)
	}
}
