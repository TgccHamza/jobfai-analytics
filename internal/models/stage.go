package models

import (
	"time"
)

type Stage struct {
	StageID     int       `gorm:"column:stage_id;primaryKey;autoIncrement" json:"stageId"`
	GameID      string    `gorm:"column:game_id;type:varchar(72)" json:"gameId"`
	StageKey    string    `gorm:"column:stage_key;type:varchar(255)" json:"stageKey"`
	StageName   string    `gorm:"column:stage_name;type:varchar(255)" json:"stageName"`
	StageOrder  int       `gorm:"column:stage_order;check:stage_order > 0" json:"stageOrder"`
	Benchmark   float64   `gorm:"column:benchmark;type:decimal(5,2)" json:"benchmark"`
	Description string    `gorm:"column:description;type:text" json:"description"`
	OptimalTime int       `gorm:"column:optimal_time" json:"optimalTime"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
	UpdatedAt   time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updatedAt"`

	// Relationships
	Game    Game               `gorm:"foreignKey:GameID;references:GameID" json:"game,omitempty"`
	Metrics []CompetenceMetric `gorm:"many2many:stage_metrics;foreignKey:StageID;joinForeignKey:StageID;References:MetricID;joinReferences:MetricID" json:"metrics,omitempty"`
}

func (Stage) TableName() string {
	return "stages"
}
