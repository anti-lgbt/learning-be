package main

import (
	"time"

	"github.com/anti-lgbt/learning-be/config"
	"github.com/anti-lgbt/learning-be/models"
	"github.com/jasonlvhit/gocron"
)

func main() {
	s := gocron.NewScheduler()
	s.Every(1).Day().At("00:00:00").Do(release_comment_statistics)
	<-s.Start()
}

type Count struct {
	Count uint64
}

func release_comment_statistics() {
	yesterday := time.Now().AddDate(0, 0, -1).Format("2006-01-02")

	var count *Count

	config.DataBase.
		Model(&models.Comment{}).
		Select("COUNT(*) AS count").
		Group("DATE(\"created_at\")").
		Having("DATE(\"created_at\") = ?", yesterday).First(&count)

	release_date, _ := time.Parse("2006-01-02", yesterday)

	release := &models.CommentStatistic{
		Count:        count.Count,
		ReferralDate: release_date,
	}

	config.DataBase.Create(&release)
}
