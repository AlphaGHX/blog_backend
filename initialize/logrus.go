package initialize

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

func Logrus() (LOG *logrus.Logger) {
	logOutFile, _ := os.OpenFile("sys.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0755)

	log := logrus.New()

	log.SetOutput(io.MultiWriter(os.Stdout, logOutFile))
	log.Formatter = &logrus.TextFormatter{}

	return log
}
