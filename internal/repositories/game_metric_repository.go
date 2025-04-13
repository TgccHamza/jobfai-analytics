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

func (r *GameMetricRepository) Create(Metric *models.Metric) error {
	return r.db.Create(Metric).Error
}

func (r *GameMetricRepository) Update(Metric *models.Metric) error {
	return r.db.Save(Metric).Error
}

func (r *GameMetricRepository) Delete(metricID int) error {
	return r.db.Delete(&models.Metric{}, "metric_id = ?", metricID).Error
}

func (r *GameMetricRepository) FindByID(metricID int) (*models.Metric, error) {
	var Metric models.Metric
	err := r.db.First(&Metric, "metric_id = ?", metricID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &Metric, nil
}

func (r *GameMetricRepository) FindByGame(gameID string) ([]models.Metric, error) {
	var gameMetrics []models.Metric
	err := r.db.Where("game_id = ?", gameID).Order("metric_name ASC").Find(&gameMetrics).Error
	return gameMetrics, err
}
