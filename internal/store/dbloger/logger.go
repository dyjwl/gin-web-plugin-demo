package dbloger

import (
	"fmt"

	"github.com/dyjwl/gin-web-plugin-demo/configs"
	"github.com/dyjwl/gin-web-plugin-demo/pkg/log"
	"gorm.io/gorm/logger"
)

type writer struct {
	logger.Writer
}

// NewWriter writer 构造函数
// Author [SliverHorn](https://github.com/SliverHorn)
func NewWriter(w logger.Writer) *writer {
	return &writer{Writer: w}
}

// Printf 格式化打印日志
// Author [SliverHorn](https://github.com/SliverHorn)
func (w *writer) Printf(message string, data ...interface{}) {
	logZap := configs.Config.Database.LogZap
	if logZap {
		log.Info(fmt.Sprintf(message+"\n", data...))
	} else {
		w.Writer.Printf(message, data...)
	}
}
