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

func (r *MetricParameterRepository) Create(parameter *models.MetricParameter) error {
	return r.db.Create(parameter).Error
}

func (r *MetricParameterRepository) Update(parameter *models.MetricParameter) error {
	return r.db.Save(parameter).Error
}

func (r *MetricParameterRepository) Delete(paramID int) error {
	return r.db.Delete(&models.MetricParameter{}, "param_id = ?", paramID).Error
}

func (r *MetricParameterRepository) FindByID(paramID int) (*models.MetricParameter, error) {
	var parameter models.MetricParameter
	err := r.db.First(&parameter, "param_id = ?", paramID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &parameter, nil
}

func (r *MetricParameterRepository) FindByMetric(metricID int) ([]models.MetricParameter, error) {
	var parameters []models.MetricParameter
	err := r.db.Where("metric_id = ?", metricID).Order("param_name ASC").Find(&parameters).Error
	return parameters, err
}

func (r *MetricParameterRepository) FindRequiredParametersForMetric(metricID int) ([]models.MetricParameter, error) {
	var parameters []models.MetricParameter
	err := r.db.Where("metric_id = ? AND is_required = ?", metricID, true).Find(&parameters).Error
	return parameters, err
}
