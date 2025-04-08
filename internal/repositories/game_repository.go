package repositories

import (
	"errors"
	"jobfai-analytics/internal/models"

	"gorm.io/gorm"
)

type GameRepository struct {
	db *gorm.DB
}

func NewGameRepository(db *gorm.DB) *GameRepository {
	return &GameRepository{db: db}
}

func (r *GameRepository) Create(game *models.Game) error {
	return r.db.Create(game).Error
}

func (r *GameRepository) Update(game *models.Game) error {
	return r.db.Save(game).Error
}

func (r *GameRepository) Delete(gameID string) error {
	return r.db.Delete(&models.Game{}, "game_id = ?", gameID).Error
}

func (r *GameRepository) FindByID(gameID string) (*models.Game, error) {
	var game models.Game
	err := r.db.First(&game, "game_id = ?", gameID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &game, nil
}

func (r *GameRepository) FindAll() ([]models.Game, error) {
	var games []models.Game
	err := r.db.Find(&games).Error
	return games, err
}

func (r *GameRepository) FindAllActive() ([]models.Game, error) {
	var games []models.Game
	err := r.db.Where("active = ?", true).Order("created_at DESC").Find(&games).Error
	return games, err
}
