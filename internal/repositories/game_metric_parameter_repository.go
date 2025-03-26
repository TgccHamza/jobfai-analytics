package repositories

import (
	"errors"

	"gorm.io/gorm"

	"jobfai-analytics/internal/models"
)

type GameMetricParameterRepository struct {
	db *gorm.DB
}

func NewGameMetricParameterRepository(db *gorm.DB) *GameMetricParameterRepository {
	return &GameMetricParameterRepository{db: db}
}

func (r *GameMetricParameterRepository) Create(parameter *models.GameMetricParameter) error {
	return r.db.Create(parameter).Error
}

func (r *GameMetricParameterRepository) Update(parameter *models.GameMetricParameter) error {
	return r.db.Save(parameter).Error
}

func (r *GameMetricParameterRepository) Delete(paramID int) error {
	return r.db.Delete(&models.GameMetricParameter{}, "param_id = ?", paramID).Error
}

func (r *GameMetricParameterRepository) FindByID(paramID int) (*models.GameMetricParameter, error) {
	var parameter models.GameMetricParameter
	err := r.db.First(&parameter, "param_id = ?", paramID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &parameter, nil
}

func (r *GameMetricParameterRepository) FindByMetric(metricID int) ([]models.GameMetricParameter, error) {
	var parameters []models.GameMetricParameter
	err := r.db.Where("metric_id = ?", metricID).Order("param_name ASC").Find(&parameters).Error
	return parameters, err
}

func (r *GameMetricParameterRepository) FindRequiredParametersForMetric(metricID int) ([]models.GameMetricParameter, error) {
	var parameters []models.GameMetricParameter
	err := r.db.Where("metric_id = ? AND is_required = ?", metricID, true).Find(&parameters).Error
	return parameters, err
}
