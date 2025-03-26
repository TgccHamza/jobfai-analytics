package models

import (
	"time"
)

type GameMetric struct {
	MetricID          int       `gorm:"column:metric_id;primaryKey;autoIncrement" json:"metricId"`
	GameID            string    `gorm:"column:game_id;type:varchar(72)" json:"gameId"`
	MetricKey         string    `gorm:"column:metric_key;type:varchar(255)" json:"metricKey"`
	MetricName        string    `gorm:"column:metric_name;type:varchar(255)" json:"metricName"`
	MetricDescription string    `gorm:"column:metric_description;type:text" json:"metricDescription"`
	Benchmark         string    `gorm:"column:benchmark;type:varchar(255)" json:"benchmark"`
	Formula           string    `gorm:"column:formula;type:text" json:"formula"`
	Description       string    `gorm:"column:description;type:text" json:"description"`
	CreatedAt         time.Time `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
	UpdatedAt         time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updatedAt"`

	// Relationships
	Game       Game                  `gorm:"foreignKey:GameID;references:GameID" json:"game,omitempty"`
	Parameters []GameMetricParameter `gorm:"foreignKey:MetricID" json:"parameters,omitempty"`
}

func (GameMetric) TableName() string {
	return "game_metrics"
}
