package helpers

import (
	"time"

	"github.com/adwinugroho/simple-wedding-management/internals/logger"
)

func TimeHostNow(tz string) time.Time {
	logger.LogInfo("Starting to call TimeHostNow")
	// you can change Asia/Jakarta with your own location.
	// check on this https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
	location, err := time.LoadLocation(tz)
	if err != nil {
		logger.LogError("Error get time, cause:" + err.Error())
	}
	now := time.Now()
	timeInLoc := now.In(location)
	return timeInLoc
}
