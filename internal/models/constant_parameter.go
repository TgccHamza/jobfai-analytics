package models

import (
	"time"
)

type ConstantParameter struct {
	ConstID          int       `gorm:"column:const_id;primaryKey;autoIncrement" json:"constId"`
	GameID           string    `gorm:"column:game_id;type:varchar(72)" json:"gameId"`
	ConstKey         string    `gorm:"column:const_key;type:varchar(255)" json:"constKey"`
	ConstName        string    `gorm:"column:const_name;type:varchar(255)" json:"constName"`
	ConstDescription string    `gorm:"column:const_description;type:text" json:"constDescription"`
	ConstValue       float64   `gorm:"column:const_value;type:decimal(10,2)" json:"constValue"`
	Description      string    `gorm:"column:description;type:text" json:"description"`
	CreatedAt        time.Time `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
	UpdatedAt        time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updatedAt"`
}

func (ConstantParameter) TableName() string {
	return "constant_parameters"
}
