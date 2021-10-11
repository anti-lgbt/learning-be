package admin

import (
	"time"

	"github.com/anti-lgbt/learning-be/config"
	"github.com/anti-lgbt/learning-be/models"
	"github.com/gofiber/fiber/v2"
)

type ProductStatistic struct {
	ProductTypeID uint64 `json:"product_type_id"`
	Count         uint64 `json:"count"`
}

func GetProductStatistics(c *fiber.Ctx) error {
	var product_statistic []*ProductStatistic

	config.DataBase.Model(&models.Product{}).Select("product_type_id, count(product_type_id) as count").Group("product_type_id").Find(&product_statistic)

	return c.Status(200).JSON(product_statistic)
}

type ProductTypeEntry struct {
	ID          uint64    `json:"id"`
	Count       uint64    `json:"count"`
	ReleaseDate time.Time `json:"release_date"`
}

func GetCommentStatistics(c *fiber.Ctx) error {
	var comment_statistics []*models.CommentStatistic

	config.DataBase.Find(&comment_statistics)

	comment_statistic_entries := make([]ProductTypeEntry, 0)

	for _, comment_statistic := range comment_statistics {
		comment_statistic_entries = append(comment_statistic_entries, ProductTypeEntry{
			ID:          comment_statistic.ID,
			Count:       comment_statistic.Count,
			ReleaseDate: comment_statistic.ReleaseDate,
		})
	}

	return c.Status(200).JSON(comment_statistic_entries)
}
