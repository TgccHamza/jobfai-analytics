package repositories

import (
	"errors"

	"gorm.io/gorm"

	"jobfai-analytics/internal/models"
)

type StageMetricRepository struct {
	db *gorm.DB
}

func NewStageMetricRepository(db *gorm.DB) *StageMetricRepository {
	return &StageMetricRepository{db: db}
}

func (r *StageMetricRepository) Create(stageMetric *models.StageMetric) error {
	return r.db.Create(stageMetric).Error
}

func (r *StageMetricRepository) Delete(stageID int, metricID int) error {
	return r.db.Delete(&models.StageMetric{}, "stage_id = ? AND metric_id = ?", stageID, metricID).Error
}

func (r *StageMetricRepository) FindByStage(stageID int) ([]models.StageMetric, error) {
	var stageMetrics []models.StageMetric
	err := r.db.Where("stage_id = ?", stageID).Find(&stageMetrics).Error
	return stageMetrics, err
}

func (r *StageMetricRepository) FindByStageAndMetric(stageID int, metricID int) (*models.StageMetric, error) {
	var stageMetric models.StageMetric
	err := r.db.Where("stage_id = ? AND metric_id = ?", stageID, metricID).First(&stageMetric).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &stageMetric, nil
}
