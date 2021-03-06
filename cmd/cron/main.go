package main

import (
	"time"

	"github.com/anti-lgbt/learning-be/config"
	"github.com/anti-lgbt/learning-be/models"
	"github.com/jasonlvhit/gocron"
)

func main() {
	if err := config.InitializeConfig(); err != nil {
		config.Logger.Error(err.Error())
		return
	}

	config.DataBase.AutoMigrate(&models.User{}, &models.ProductType{}, &models.Product{}, &models.Comment{}, &models.CommentStatistic{})

	s := gocron.NewScheduler()
	s.Every(1).Day().At("00:00:00").Do(release_comment_statistics)
	<-s.Start()
}

func release_comment_statistics() {
	yesterday := time.Now().AddDate(0, 0, -1).Format("2006-01-02")

	var count uint64

	config.DataBase.
		Model(&models.Comment{}).
		Select("COUNT(*) AS count").
		Group("DATE(\"created_at\")").
		Having("DATE(\"created_at\") = ?", yesterday).Scan(&count)

	release_date, _ := time.Parse("2006-01-02", yesterday)

	release := &models.CommentStatistic{
		Count:       count,
		ReleaseDate: release_date,
	}

	config.DataBase.Create(&release)
}
