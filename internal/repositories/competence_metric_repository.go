package repositories

import (
	"errors"

	"gorm.io/gorm"

	"jobfai-analytics/internal/models"
)

type CompetenceMetricRepository struct {
	db *gorm.DB
}

func NewCompetenceMetricRepository(db *gorm.DB) *CompetenceMetricRepository {
	return &CompetenceMetricRepository{db: db}
}

func (r *CompetenceMetricRepository) Create(metric *models.CompetenceMetric) error {
	return r.db.Create(metric).Error
}

func (r *CompetenceMetricRepository) Update(metric *models.CompetenceMetric) error {
	return r.db.Save(metric).Error
}

func (r *CompetenceMetricRepository) Delete(metricID int) error {
	return r.db.Delete(&models.CompetenceMetric{}, "metric_id = ?", metricID).Error
}

func (r *CompetenceMetricRepository) FindByID(metricID int) (*models.CompetenceMetric, error) {
	var metric models.CompetenceMetric
	err := r.db.First(&metric, "metric_id = ?", metricID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &metric, nil
}

func (r *CompetenceMetricRepository) FindByCompetence(competenceID int) ([]models.CompetenceMetric, error) {
	var metrics []models.CompetenceMetric
	err := r.db.Where("competence_id = ?", competenceID).Order("metric_name ASC").Find(&metrics).Error
	return metrics, err
}

func (r *CompetenceMetricRepository) FindWithParameters(metricID int) (*models.CompetenceMetric, error) {
	var metric models.CompetenceMetric
	err := r.db.Preload("Parameters").
		First(&metric, "metric_id = ?", metricID).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &metric, nil
}

func (r *CompetenceMetricRepository) FindByStage(stageID int) ([]models.CompetenceMetric, error) {
	var metrics []models.CompetenceMetric
	err := r.db.Joins("JOIN stage_metrics ON competence_metrics.metric_id = stage_metrics.metric_id").
		Where("stage_metrics.stage_id = ?", stageID).
		Find(&metrics).Error
	return metrics, err
}

// FindByIDs retrieves competence metrics by their IDs
func (r *CompetenceMetricRepository) FindByIDs(metricIDs []int) ([]models.CompetenceMetric, error) {
	var metrics []models.CompetenceMetric
	err := r.db.Where("metric_id IN ?", metricIDs).Find(&metrics).Error
	return metrics, err
}
