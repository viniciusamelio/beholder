package utils

import (
	"time"

	"github.com/godruoyi/go-snowflake"
)

func GenSnowflakeID() int32 {
	snowflake.SetMachineID(1)
	snowflake.SetStartTime(time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC))
	return int32(snowflake.ID())
}
