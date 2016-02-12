/*
 * 封装之后可以既输出到命令行，又输入到日志文件
 */

package log

import (
	"github.com/Unknwon/com"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"os"
	"path"
)

type Alog struct {
	logger *logs.BeeLogger
}

func (a *Alog) init(l *logs.BeeLogger, logfile string) {
	a.logger = l
	// if logfile == "" {
	// 	logfile = "log/logs.log"
	// }
	if !com.IsExist(logfile) {
		os.MkdirAll(path.Dir(logfile), os.ModePerm)
		os.Create(logfile)
	}
	a.logger.SetLogger("file", `{"filename":"`+logfile+`"}`)
	a.logger.EnableFuncCallDepth(true) //显示报错文件名和行号
	a.logger.SetLogFuncCallDepth(3)    //默认为2，适合于没有封装的情况
}

func Trace(message string) {
	beego.Trace(message)
	Log.logger.Trace(message)
}

func Warn(message string) {
	beego.Warn(message)
	Log.logger.Warn(message)
}

func Critical(message string) {
	beego.Critical(message)
	Log.logger.Critical(message)
}

var Log Alog

func init() {
	Log = Alog{}
	Log.init(logs.NewLogger(10000), "log/logs.log")
}
