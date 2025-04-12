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

func (r *StageMetricRepository) Create(Metric *models.Metric) error {
	return r.db.Create(Metric).Error
}

func (r *StageMetricRepository) Delete(stageID int, metricID int) error {
	return r.db.Delete(&models.Metric{}, "stage_id = ? AND metric_id = ?", stageID, metricID).Error
}

func (r *StageMetricRepository) FindByStage(stageID int) ([]models.Metric, error) {
	var stageMetrics []models.Metric
	err := r.db.Where("stage_id = ?", stageID).Find(&stageMetrics).Error
	return stageMetrics, err
}

func (r *StageMetricRepository) FindByCompetence(competenceID int) ([]models.Metric, error) {
	var competenceMetrics []models.Metric
	err := r.db.Where("competence_id = ?", competenceID).Find(&competenceMetrics).Error
	return competenceMetrics, err
}

func (r *StageMetricRepository) FindByStageAndMetric(stageID int, metricID int) (*models.Metric, error) {
	var Metric models.Metric
	err := r.db.Where("stage_id = ? AND metric_id = ?", stageID, metricID).First(&Metric).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &Metric, nil
}
