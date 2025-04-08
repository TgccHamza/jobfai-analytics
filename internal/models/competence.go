package models

import (
	"time"
)

type Competence struct {
	CompetenceID    int       `gorm:"column:competence_id;primaryKey;autoIncrement" json:"competenceId"`
	ParentID        int       `gorm:"column:parent_id;primaryKey;autoIncrement" json:"parentId"`
	GameID          string    `gorm:"column:game_id;type:varchar(72)" json:"gameId"`
	CompetenceKey   string    `gorm:"column:competence_key;type:varchar(255)" json:"competenceKey"`
	CompetenceName  string    `gorm:"column:competence_name;type:varchar(255)" json:"competenceName"`
	Benchmark       float64   `gorm:"column:benchmark;type:decimal(5,2)" json:"benchmark"`
	BenchmarkMargin float64   `gorm:"column:benchmark_margin;type:decimal(5,2)" json:"benchmark_margin"`
	Description     string    `gorm:"column:description;type:text" json:"description"`
	Weight          float64   `gorm:"column:weight;type:decimal(5,2);default:1.0" json:"weight"`
	CreatedAt       time.Time `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
	UpdatedAt       time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updatedAt"`
}

func (Competence) TableName() string {
	return "competencies"
}
