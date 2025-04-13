package models

import "time"

type Metric struct {
	MetricID          int       `gorm:"column:metric_id;primaryKey;autoIncrement" json:"metricId"`
	CompetenceID      int       `gorm:"column:competence_id" json:"competenceId"`
	StageID           int       `gorm:"column:stage_id" json:"stageId"`
	GameID            string    `gorm:"column:game_id" json:"gameId"`
	MetricKey         string    `gorm:"column:metric_key;type:varchar(255)" json:"metricKey"`
	MetricName        string    `gorm:"column:metric_name;type:varchar(255)" json:"metricName"`
	MetricDescription string    `gorm:"column:metric_description;type:text" json:"metricDescription"`
	Benchmark         float64   `gorm:"column:benchmark;type:decimal(5,2)" json:"benchmark"`
	BenchmarkMargin   float64   `gorm:"column:benchmark_margin;type:decimal(5,2)" json:"benchmark_margin"`
	Formula           string    `gorm:"column:formula;type:text" json:"formula"`
	Weight            float64   `gorm:"column:weight;type:decimal(5,2);default:1.0" json:"weight"`
	CreatedAt         time.Time `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
	UpdatedAt         time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updatedAt"`
}

func (Metric) TableName() string {
	return "metrics"
}
