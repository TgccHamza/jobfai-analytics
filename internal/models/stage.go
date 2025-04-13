package models

import (
	"time"
)

type Stage struct {
	StageID         int       `gorm:"column:stage_id;primaryKey;autoIncrement" json:"stageId"`
	GameID          string    `gorm:"column:game_id;type:varchar(72)" json:"gameId"`
	StageKey        string    `gorm:"column:stage_key;type:varchar(255)" json:"stageKey"`
	StageName       string    `gorm:"column:stage_name;type:varchar(255)" json:"stageName"`
	StageOrder      int       `gorm:"column:stage_order;check:stage_order > 0" json:"stageOrder"`
	Benchmark       float64   `gorm:"column:benchmark;type:decimal(5,2)" json:"benchmark"`
	BenchmarkMargin float64   `gorm:"column:benchmark_margin;type:decimal(5,2)" json:"benchmark_margin"`
	Description     string    `gorm:"column:description;type:text" json:"description"`
	OptimalTime     int       `gorm:"column:optimal_time" json:"optimalTime"`
	CreatedAt       time.Time `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
	UpdatedAt       time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updatedAt"`
}

func (Stage) TableName() string {
	return "stages"
}
