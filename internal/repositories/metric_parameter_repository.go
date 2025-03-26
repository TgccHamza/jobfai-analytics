package repositories

import (
	"errors"

	"gorm.io/gorm"

	"jobfai-analytics/internal/models"
)

type MetricParameterRepository struct {
	db *gorm.DB
}

func NewMetricParameterRepository(db *gorm.DB) *MetricParameterRepository {
	return &MetricParameterRepository{db: db}
}

func (r *MetricParameterRepository) Create(parameter *models.CompetenceMetricParameter) error {
	return r.db.Create(parameter).Error
}

func (r *MetricParameterRepository) Update(parameter *models.CompetenceMetricParameter) error {
	return r.db.Save(parameter).Error
}

func (r *MetricParameterRepository) Delete(paramID int) error {
	return r.db.Delete(&models.CompetenceMetricParameter{}, "param_id = ?", paramID).Error
}

func (r *MetricParameterRepository) FindByID(paramID int) (*models.CompetenceMetricParameter, error) {
	var parameter models.CompetenceMetricParameter
	err := r.db.First(&parameter, "param_id = ?", paramID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &parameter, nil
}

func (r *MetricParameterRepository) FindByMetric(metricID int) ([]models.CompetenceMetricParameter, error) {
	var parameters []models.CompetenceMetricParameter
	err := r.db.Where("metric_id = ?", metricID).Order("param_name ASC").Find(&parameters).Error
	return parameters, err
}

func (r *MetricParameterRepository) FindRequiredParametersForMetric(metricID int) ([]models.CompetenceMetricParameter, error) {
	var parameters []models.CompetenceMetricParameter
	err := r.db.Where("metric_id = ? AND is_required = ?", metricID, true).Find(&parameters).Error
	return parameters, err
}
