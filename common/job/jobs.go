package job

func EveryMinute() {
	logger.Info("This runs every minute")
}

func Weekdays() {
	logger.Info("This runs at 04:00 AM on weekdays")
}
