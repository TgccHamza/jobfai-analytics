package repositories

import (
	"fmt"

	"jobfai-analytics/internal/models"

	"gorm.io/gorm"
)

// ConstantParameterRepository handles database operations for constant parameters using GORM
type ConstantParameterRepository struct {
	db *gorm.DB
}

// NewConstantParameterRepository creates a new constant parameter repository
func NewConstantParameterRepository(db *gorm.DB) *ConstantParameterRepository {
	return &ConstantParameterRepository{
		db: db,
	}
}

// FindByID retrieves a constant parameter by its ID
func (r *ConstantParameterRepository) FindByID(constID int) (*models.ConstantParameter, error) {
	var constant models.ConstantParameter
	result := r.db.Where("const_id = ?", constID).First(&constant)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, fmt.Errorf("error finding constant parameter by ID: %w", result.Error)
	}

	return &constant, nil
}

// FindByGame retrieves all constant parameters for a game
func (r *ConstantParameterRepository) FindByGame(gameID string) ([]models.ConstantParameter, error) {
	var constants []models.ConstantParameter
	result := r.db.Where("game_id = ?", gameID).Order("const_key").Find(&constants)

	if result.Error != nil {
		return nil, fmt.Errorf("error finding constant parameters by game: %w", result.Error)
	}

	return constants, nil
}

// FindByKey retrieves a constant parameter by its key for a specific game
func (r *ConstantParameterRepository) FindByKey(gameID, constKey string) (*models.ConstantParameter, error) {
	var constant models.ConstantParameter
	result := r.db.Where("game_id = ? AND const_key = ?", gameID, constKey).First(&constant)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, fmt.Errorf("error finding constant parameter by key: %w", result.Error)
	}

	return &constant, nil
}

// Create inserts a new constant parameter
func (r *ConstantParameterRepository) Create(constant *models.ConstantParameter) error {
	result := r.db.Create(constant)

	if result.Error != nil {
		return fmt.Errorf("error creating constant parameter: %w", result.Error)
	}

	return nil
}

// Update updates an existing constant parameter
func (r *ConstantParameterRepository) Update(constant *models.ConstantParameter) error {
	result := r.db.Save(constant)

	if result.Error != nil {
		return fmt.Errorf("error updating constant parameter: %w", result.Error)
	}

	return nil
}

// Delete removes a constant parameter
func (r *ConstantParameterRepository) Delete(constID int) error {
	result := r.db.Delete(&models.ConstantParameter{}, constID)

	if result.Error != nil {
		return fmt.Errorf("error deleting constant parameter: %w", result.Error)
	}

	return nil
}

// DeleteByGame removes all constant parameters for a game
func (r *ConstantParameterRepository) DeleteByGame(gameID string) error {
	result := r.db.Where("game_id = ?", gameID).Delete(&models.ConstantParameter{})

	if result.Error != nil {
		return fmt.Errorf("error deleting constant parameters by game: %w", result.Error)
	}

	return nil
}
