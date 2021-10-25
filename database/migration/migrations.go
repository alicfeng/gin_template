package migration

import (
	"github.com/kataras/golog"
	"gorm.io/gorm"
)

func Handle(database *gorm.DB) {
	err := database.
		Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4").
		AutoMigrate()
	if nil != err {
		golog.Error("database migration failure", err)
		panic("database migration failure")
	}

	golog.Info("database migration finished")
}
