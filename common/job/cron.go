package job

import (
	"go-web/common/log"

	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

func init() {
	logger = log.GetLogger()
}

func StartJob() {

	c := cron.New()

	_, err := c.AddFunc("*/1 * * * *", EveryMinute)
	if err != nil {
		logger.Errorf("Error adding cron job [%s]: %v", "EveryMinute", err)
	}

	_, err = c.AddFunc("0 4 * * MON-FRI", Weekdays)
	if err != nil {
		logger.Errorf("Error adding cron job [%s]: %v", "Weekdays", err)
	}

	c.Start()
}
