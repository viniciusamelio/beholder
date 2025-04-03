package crons

import (
	"beholder-api/internal/services/bucket"

	"github.com/robfig/cron"
)

func Crons(bucket bucket.Bucket) {
	go func() {
		c := cron.New()
		BackupCron(c, bucket)
		c.Run()
	}()
}
