package seeder

import (
	"github.com/kataras/golog"
	"gorm.io/gorm"
)

func Handle(database *gorm.DB) {
	// a....

	golog.Info("database seeder finished")
}
