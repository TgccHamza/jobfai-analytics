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

func (r *CompetenceMetricRepository) Create(metric *models.Metric) error {
	return r.db.Create(metric).Error
}

func (r *CompetenceMetricRepository) Update(metric *models.Metric) error {
	return r.db.Save(metric).Error
}

func (r *CompetenceMetricRepository) Delete(metricID int) error {
	return r.db.Delete(&models.Metric{}, "metric_id = ?", metricID).Error
}

func (r *CompetenceMetricRepository) FindByID(metricID int) (*models.Metric, error) {
	var metric models.Metric
	err := r.db.First(&metric, "metric_id = ?", metricID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &metric, nil
}

func (r *CompetenceMetricRepository) FindByCompetence(competenceID int) ([]models.Metric, error) {
	var metrics []models.Metric
	err := r.db.Where("competence_id = ?", competenceID).Order("metric_name ASC").Find(&metrics).Error
	return metrics, err
}

func (r *CompetenceMetricRepository) FindByStage(stageID int) ([]models.Metric, error) {
	var metrics []models.Metric
	err := r.db.Where("metrics.stage_id = ?", stageID).
		Find(&metrics).Error
	return metrics, err
}

// FindByIDs retrieves competence metrics by their IDs
func (r *CompetenceMetricRepository) FindByIDs(metricIDs []int) ([]models.Metric, error) {
	var metrics []models.Metric
	err := r.db.Where("metric_id IN ?", metricIDs).Find(&metrics).Error
	return metrics, err
}
