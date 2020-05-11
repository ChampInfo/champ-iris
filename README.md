# 說明

本專案為 [kataras/iris](https://github.com/kataras/iris) 以及 [iris mvc](https://github.com/kataras/iris/tree/master/mvc) 的後台快速實作框架

系統需求為 [Go Programming Language](https://golang.org/dl/), version 1.13 and above.

##安裝
```shell script
$ go get git.championtek.com.tw/go/champiris
```

或者編輯你專案中的go.mod檔案
```
module your_project_name

go 1.13

require (
    git.championtek.com.tw/go/champiris v0.2.10
)
```

```shell script
$ go build
```

##如何開始
```go
package main

import (
    "git.championtek.com.tw/go/champiris"
    "github.com/kataras/iris/v12"
    "github.com/kataras/iris/v12/mvc"
)

func main()  {
    var service champiris.Service
    var config champiris.NetConfig
    config = champiris.NetConfig{
    	Port:         "8080",
    	LoggerEnable: false,
    	JWTEnable:    false}
    service.New(&config)
    service.AddRoute(champiris.RouterSet{
    		Party: "/api",
    		Router:apiRouter,
    	})
    	service.Run()
}
func apiRouter(app *mvc.Application)  {
    app.Handle(new(Hello))
}

type Hello struct {}

//get localhost:port/api/hello 
func (h *Hello) GetHello(Ctx iris.Context) {
	Ctx.WriteString("HelloWord")
}
```

iris MVC 相關操作請見[官方文件](https://github.com/kataras/iris/wiki/MVC)



* Local端log放置目錄為 logs
* 當設置檔內 `LoggerEnable` 為 `true` 時，紀錄會上傳到 `ELK伺服器`
