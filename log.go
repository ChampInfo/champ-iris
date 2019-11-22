package champiris

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
	"strings"
	"time"

	"github.com/kataras/iris/v12"

	champLogger "git.championtek.com.tw/go/logger"
	"git.championtek.com.tw/go/logger/templates"

	"github.com/kataras/iris/v12/middleware/logger"
)

const (
	logFolder         = "logs"
	deleteFileOnExist = true
)

var excludeExtensions = [...]string{
	".js",
	".css",
	".jpg",
	".jpeg",
	".png",
	".ico",
	".svg",
}

func NewLoggerManager(config ELK) error {
	eklMap := config.Mapping
	settings := templates.Settings{
		NumberOfShards:   eklMap.Settings.NumberOfShards,
		NumberOfReplicas: eklMap.Settings.NumberOfReplicas,
	}
	properties := templates.Properties{
		Service: templates.Property{PropertyType: eklMap.Mappings.Properties.Service},
		IP:      templates.Property{PropertyType: eklMap.Mappings.Properties.IP},
		Status:  templates.Property{PropertyType: eklMap.Mappings.Properties.Status},
		Method:  templates.Property{PropertyType: eklMap.Mappings.Properties.Method},
		Path:    templates.Property{PropertyType: eklMap.Mappings.Properties.Path},
		Created: templates.Property{PropertyType: eklMap.Mappings.Properties.Created},
		Tags:    templates.Property{PropertyType: eklMap.Mappings.Properties.Tags},
		Remark:  templates.Property{PropertyType: eklMap.Mappings.Properties.Remark},
	}
	mappings := templates.Mappings{Properties: properties}
	template := templates.APIServiceMapping{
		Settings: settings,
		Mappings: mappings,
	}
	b, err := json.Marshal(template)
	if err != nil {
		return err
	}
	cfg := champLogger.Config{
		Host:     config.Host.URL,
		Port:     config.Host.PORT,
		User:     config.BasicAuth.User,
		Password: config.BasicAuth.Password,
		Index:    config.Index,
		Mapping:  string(b),
	}
	log.Printf("auth user:%v password:%v", config.BasicAuth.User, config.BasicAuth.Password)
	if err := champLogger.Mgr.Init(&cfg); err != nil {
		return err
	}
	return nil
}

func (api *API) todayFileName() string {
	today := time.Now()
	formatted := fmt.Sprintf("%d-%02d-%02d", today.Year(), today.Month(), today.Day())
	return path.Join(logFolder, formatted+".txt")
}
func (api *API) newLogFile() *os.File {
	fileName := api.todayFileName()

	if _, err := os.Stat(logFolder); os.IsExist(err) {
		err = os.Mkdir(logFolder, os.ModePerm)
		if err != nil {
			log.Fatal("Create logs folder error: ", err)
		}
	}

	f, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Open log file error : ", err)
	}
	return f
}

func (api *API) newRequestLogger() (h iris.Handler, close func() error) {
	close = func() error { return nil }

	c := logger.Config{
		Status:  true,
		IP:      true,
		Method:  true,
		Path:    true,
		Columns: true,
	}

	logFile := api.newLogFile()

	fmt.Println(logFile.Name())

	close = func() error {
		err := logFile.Close()
		if deleteFileOnExist {
			err = os.Remove(logFile.Name())
		}
		return err
	}
	c.LogFunc = func(now time.Time, latency time.Duration, status, ip, method, path string, message interface{}, headerMessage interface{}) {
		output := logger.Columnize(now.Format("2006/01/02 - 15:04:05"), latency, status, ip, method, path, message, headerMessage)
		_, _ = logFile.Write([]byte(output))
	}

	c.AddSkipper(func(ctx iris.Context) bool {
		p := ctx.Path()
		for _, ext := range excludeExtensions {
			if strings.HasSuffix(p, ext) {
				return true
			}
		}
		return false
	})
	h = logger.New(c)
	return
}
