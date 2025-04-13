package models

import (
	"time"
)

type Game struct {
	GameID      string    `gorm:"column:game_id;primaryKey;type:varchar(36)" json:"gameId"`
	GameName    string    `gorm:"column:game_name;unique;type:varchar(255)" json:"gameName"`
	Description string    `gorm:"column:description;type:text" json:"description"`
	Active      bool      `gorm:"column:active;default:true" json:"active"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
	UpdatedAt   time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updatedAt"`
}

func (Game) TableName() string {
	return "games"
}
