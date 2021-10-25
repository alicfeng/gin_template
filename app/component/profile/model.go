package profile

import "github.com/alicfeng/gin_template/app/model"

type Profile struct {
	model.Model
	Name    string `json:"name" gorm:"column:name;size:32" comment:"名称"`
	Code    int    `json:"type" gorm:"column:code;not null" comment:"编码"`
	Content string `json:"content" gorm:"column:content;type:json" comment:"配置内容"`
}

func (Profile) TableName() string {
	return "profile"
}
