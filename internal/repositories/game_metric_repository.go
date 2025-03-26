package repositories

import (
	"errors"

	"gorm.io/gorm"

	"jobfai-analytics/internal/models"
)

type GameMetricRepository struct {
	db *gorm.DB
}

func NewGameMetricRepository(db *gorm.DB) *GameMetricRepository {
	return &GameMetricRepository{db: db}
}

func (r *GameMetricRepository) Create(gameMetric *models.GameMetric) error {
	return r.db.Create(gameMetric).Error
}

func (r *GameMetricRepository) Update(gameMetric *models.GameMetric) error {
	return r.db.Save(gameMetric).Error
}

func (r *GameMetricRepository) Delete(metricID int) error {
	return r.db.Delete(&models.GameMetric{}, "metric_id = ?", metricID).Error
}

func (r *GameMetricRepository) FindByID(metricID int) (*models.GameMetric, error) {
	var gameMetric models.GameMetric
	err := r.db.First(&gameMetric, "metric_id = ?", metricID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &gameMetric, nil
}

func (r *GameMetricRepository) FindByGame(gameID string) ([]models.GameMetric, error) {
	var gameMetrics []models.GameMetric
	err := r.db.Where("game_id = ?", gameID).Order("metric_name ASC").Find(&gameMetrics).Error
	return gameMetrics, err
}

func (r *GameMetricRepository) FindWithParameters(metricID int) (*models.GameMetric, error) {
	var gameMetric models.GameMetric
	err := r.db.Preload("Parameters").
		First(&gameMetric, "metric_id = ?", metricID).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &gameMetric, nil
}
