package models

import "gorm.io/gorm"

type Problem struct {
	gorm.Model
	Identity   string `gorm:"column:identity;type:varchar(36);" json:"identity"`        // 问题的唯一标识符
	CategoryId string `gorm:"column:category_id;type:varchar(255);" json:"category_id"` // 分类ID，已逗号分割
	Title      string `gorm:"column:title;type:varchar(36);" json:"title"`              //标题
	Content    string `gorm:"column:column;type:text;" json:"content"`                  //正文
}

func (table *Problem) TableName() string {
	return "problem"
}
