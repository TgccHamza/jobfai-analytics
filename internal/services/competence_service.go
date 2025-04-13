package services

import (
	"errors"
	"fmt"

	"jobfai-analytics/internal/models"
	"jobfai-analytics/internal/repositories"
)

// CompetenceService handles competence-related business logic
type CompetenceService struct {
	competenceRepository       *repositories.CompetenceRepository
	competenceMetricRepository *repositories.CompetenceMetricRepository
}

// NewCompetenceService creates a new competence service
func NewCompetenceService(
	competenceRepository *repositories.CompetenceRepository,
	competenceMetricRepository *repositories.CompetenceMetricRepository,
) *CompetenceService {
	return &CompetenceService{
		competenceRepository:       competenceRepository,
		competenceMetricRepository: competenceMetricRepository,
	}
}

// GetCompetenceByID retrieves a competence by its ID
func (s *CompetenceService) GetCompetenceByID(competenceID int) (*models.Competence, error) {
	competence, err := s.competenceRepository.FindByID(competenceID)
	if err != nil {
		return nil, fmt.Errorf("error retrieving competence: %w", err)
	}

	if competence == nil {
		return nil, errors.New("competence not found")
	}

	return competence, nil
}

// CreateCompetence creates a new competence
func (s *CompetenceService) CreateCompetence(competence *models.Competence) error {
	if competence.GameID == "" || competence.CompetenceKey == "" || competence.CompetenceName == "" {
		return errors.New("game ID, competence key, and name are required")
	}

	return s.competenceRepository.Create(competence)
}

// UpdateCompetence updates an existing competence
func (s *CompetenceService) UpdateCompetence(competence *models.Competence) error {
	if competence.CompetenceID == 0 {
		return errors.New("competence ID is required")
	}

	existingCompetence, err := s.competenceRepository.FindByID(competence.CompetenceID)
	if err != nil {
		return fmt.Errorf("error checking existing competence: %w", err)
	}

	if existingCompetence == nil {
		return errors.New("competence not found")
	}

	return s.competenceRepository.Update(competence)
}

// DeleteCompetence deletes a competence
func (s *CompetenceService) DeleteCompetence(competenceID int) error {
	existingCompetence, err := s.competenceRepository.FindByID(competenceID)
	if err != nil {
		return fmt.Errorf("error checking existing competence: %w", err)
	}

	if existingCompetence == nil {
		return errors.New("competence not found")
	}

	return s.competenceRepository.Delete(competenceID)
}

// GetCompetenceMetrics retrieves all metrics for a competence
func (s *CompetenceService) GetCompetenceMetrics(competenceID int) ([]models.Metric, error) {
	return s.competenceMetricRepository.FindByCompetence(competenceID)
}
