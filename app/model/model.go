package model

import (
	"gorm.io/gorm"
)

type Model struct {
	Id        int            `json:"id,omitempty" gorm:"column:id;type:INT UNSIGNED AUTO_INCREMENT;primary_key;" comment:"主键标识"`
	CreatedAt Time           `json:"created_at,omitempty" gorm:"column:created_at" comment:"创建时间"`
	UpdatedAt Time           `json:"updated_at,omitempty" gorm:"column:updated_at" comment:"更新时间"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"column:deleted_at;index" comment:"删除时间"`
}
