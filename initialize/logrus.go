package initialize

import (
	"os"

	"github.com/sirupsen/logrus"
)

func Logrus() (FILE_LOG *logrus.Logger, STD_LOG *logrus.Logger) {
	filelog := logrus.New()
	stdlog := logrus.New()
	logOutFile, _ := os.OpenFile("blog.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0755)

	stdlog.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	filelog.SetFormatter(&logrus.JSONFormatter{})
	filelog.SetOutput(logOutFile)

	return filelog, stdlog
}
