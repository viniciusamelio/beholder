package crons

import (
	"beholder-api/internal/services/bucket"
	"log"
	"os/exec"

	"github.com/robfig/cron"
)

func BackupCron(c *cron.Cron, bucket bucket.Bucket) {
	err := c.AddFunc("@every 5s", func() {
		log.Default().Println("[BACKUP] JOB RUNNING 🔄")
		exec.Command("sqlite3", "./beholder.db", ".backup backup.db").Output()
		bucket.UploadFile("backup", "backup.db", "application/octet-stream")
		log.Default().Println("[BACKUP] FILE UPLOADED ✅")
		exec.Command("rm", "backup.db").Output()
	})

	if err != nil {
		log.Default().Println("Error adding cron job: ", err)
	}
}
