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

	// Relationships
	Competencies       []Competence        `gorm:"foreignKey:GameID;references:GameID" json:"competencies,omitempty"`
	Stages             []Stage             `gorm:"foreignKey:GameID;references:GameID" json:"stages,omitempty"`
	GameMetrics        []GameMetric        `gorm:"foreignKey:GameID;references:GameID" json:"gameMetrics,omitempty"`
	ConstantParameters []ConstantParameter `gorm:"foreignKey:GameID;references:GameID" json:"constantParameters,omitempty"`
}

func (Game) TableName() string {
	return "games"
}
