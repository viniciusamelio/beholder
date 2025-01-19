package utils

import (
	"github.com/godruoyi/go-snowflake"
)

func GenSnowflakeID() int {
	id := snowflake.ID()
	return int(id)
}
