package models

import (
	"math"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Paginated interface {
	Count(db *gorm.DB) int64
	Take(db *gorm.DB, limit int, offset int) interface{}
}

func Paginate(db *gorm.DB, entity Paginated, page int) fiber.Map {
	limit := 15                  // TODO: this is the amount of records per page. Make it configurable, or get it from the query string.
	offset := (page - 1) * limit // From where we start the pagination

	data := entity.Take(db, limit, offset)
	total := entity.Count(db)

	return fiber.Map{
		"data": data,
		"pagination": fiber.Map{
			"totalRecords": total,
			"currentPage":  page,
			"lastPage":     math.Ceil(float64(int(total) / limit)),
		},
	}
}
