package services

import (
	"errors"
	"fmt"

	"jobfai-analytics/internal/models"
	"jobfai-analytics/internal/repositories"
)

// ConstantParameterService handles constant parameter-related business logic
type ConstantParameterService struct {
	constantParameterRepository *repositories.ConstantParameterRepository
}

// NewConstantParameterService creates a new constant parameter service
func NewConstantParameterService(
	constantParameterRepository *repositories.ConstantParameterRepository,
) *ConstantParameterService {
	return &ConstantParameterService{
		constantParameterRepository: constantParameterRepository,
	}
}

// GetConstantParameterByID retrieves a constant parameter by its ID
func (s *ConstantParameterService) GetConstantParameterByID(constID int) (*models.ConstantParameter, error) {
	constant, err := s.constantParameterRepository.FindByID(constID)
	if err != nil {
		return nil, fmt.Errorf("error retrieving constant parameter: %w", err)
	}

	if constant == nil {
		return nil, errors.New("constant parameter not found")
	}

	return constant, nil
}

// GetConstantParametersByGame retrieves all constant parameters for a game
func (s *ConstantParameterService) GetConstantParametersByGame(gameID string) ([]models.ConstantParameter, error) {
	return s.constantParameterRepository.FindByGame(gameID)
}

// CreateConstantParameter creates a new constant parameter
func (s *ConstantParameterService) CreateConstantParameter(constant *models.ConstantParameter) error {
	if constant.GameID == "" || constant.ConstKey == "" || constant.ConstName == "" {
		return errors.New("game ID, constant key, and name are required")
	}

	return s.constantParameterRepository.Create(constant)
}

// UpdateConstantParameter updates an existing constant parameter
func (s *ConstantParameterService) UpdateConstantParameter(constant *models.ConstantParameter) error {
	if constant.ConstID == 0 {
		return errors.New("constant ID is required")
	}

	existingConstant, err := s.constantParameterRepository.FindByID(constant.ConstID)
	if err != nil {
		return fmt.Errorf("error checking existing constant: %w", err)
	}

	if existingConstant == nil {
		return errors.New("constant parameter not found")
	}

	return s.constantParameterRepository.Update(constant)
}

// DeleteConstantParameter deletes a constant parameter
func (s *ConstantParameterService) DeleteConstantParameter(constID int) error {
	existingConstant, err := s.constantParameterRepository.FindByID(constID)
	if err != nil {
		return fmt.Errorf("error checking existing constant: %w", err)
	}

	if existingConstant == nil {
		return errors.New("constant parameter not found")
	}

	return s.constantParameterRepository.Delete(constID)
}
