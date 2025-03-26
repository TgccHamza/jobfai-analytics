package repositories

import (
	"errors"

	"gorm.io/gorm"

	"jobfai-analytics/internal/models"
)

type CompetenceRepository struct {
	db *gorm.DB
}

func NewCompetenceRepository(db *gorm.DB) *CompetenceRepository {
	return &CompetenceRepository{db: db}
}

func (r *CompetenceRepository) Create(competence *models.Competence) error {
	return r.db.Create(competence).Error
}

func (r *CompetenceRepository) Update(competence *models.Competence) error {
	return r.db.Save(competence).Error
}

func (r *CompetenceRepository) Delete(competenceID int) error {
	return r.db.Delete(&models.Competence{}, "competence_id = ?", competenceID).Error
}

func (r *CompetenceRepository) FindByID(competenceID int) (*models.Competence, error) {
	var competence models.Competence
	err := r.db.First(&competence, "competence_id = ?", competenceID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &competence, nil
}

func (r *CompetenceRepository) FindByGame(gameID string) ([]models.Competence, error) {
	var competencies []models.Competence
	err := r.db.Where("game_id = ?", gameID).Order("competence_name ASC").Find(&competencies).Error
	return competencies, err
}

func (r *CompetenceRepository) FindWithMetrics(competenceID int) (*models.Competence, error) {
	var competence models.Competence
	err := r.db.Preload("Metrics.Parameters").
		First(&competence, "competence_id = ?", competenceID).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &competence, nil
}
