package champiris

import (
	"fmt"
	"log"
	"os"
	"path"
	"strings"
	"time"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
)

var (
	logFolder = "logs"
)

const (
	deleteFileOnExist = true
)

func init() {
	workDirPath, _ := os.Getwd()
	logFolder = path.Join(workDirPath, logFolder)
}

var excludeExtensions = [...]string{
	".js",
	".css",
	".jpg",
	".jpeg",
	".png",
	".ico",
	".svg",
}

func (service *Service) todayFileName() string {
	today := time.Now()
	formatted := fmt.Sprintf("%s_%d-%02d-%02d", "champiris", today.Year(), today.Month(), today.Day())
	return path.Join(logFolder, formatted+".txt")
}

func (service *Service) newLogFile() *os.File {
	fileName := service.todayFileName()

	if _, err := os.Stat(logFolder); os.IsNotExist(err) {
		err = os.MkdirAll(logFolder, os.ModePerm)
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

func (service *Service) newRequestLogger() (h iris.Handler, close func() error) {
	close = func() error { return nil }

	c := logger.Config{
		Status:  true,
		IP:      true,
		Method:  true,
		Path:    true,
		Columns: true,
	}

	logFile := service.newLogFile()

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
