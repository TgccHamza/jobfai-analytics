package repositories

import (
	"errors"

	"gorm.io/gorm"

	"jobfai-analytics/internal/models"
)

type StageRepository struct {
	db *gorm.DB
}

func NewStageRepository(db *gorm.DB) *StageRepository {
	return &StageRepository{db: db}
}

func (r *StageRepository) Create(stage *models.Stage) error {
	return r.db.Create(stage).Error
}

func (r *StageRepository) Update(stage *models.Stage) error {
	return r.db.Save(stage).Error
}

func (r *StageRepository) Delete(stageID int) error {
	return r.db.Delete(&models.Stage{}, "stage_id = ?", stageID).Error
}

func (r *StageRepository) FindByID(stageID int) (*models.Stage, error) {
	var stage models.Stage
	err := r.db.First(&stage, "stage_id = ?", stageID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &stage, nil
}

func (r *StageRepository) FindByGame(gameID string) ([]models.Stage, error) {
	var stages []models.Stage
	err := r.db.Where("game_id = ?", gameID).Order("stage_order ASC").Find(&stages).Error
	return stages, err
}
