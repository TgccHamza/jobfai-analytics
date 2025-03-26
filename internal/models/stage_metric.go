package models

type StageMetric struct {
	MetricID int `gorm:"column:metric_id" json:"metricId"`
	StageID  int `gorm:"column:stage_id" json:"stageId"`
}

func (StageMetric) TableName() string {
	return "stage_metrics"
}
