package services

import (
	"errors"
	"fmt"

	"jobfai-analytics/internal/models"
	"jobfai-analytics/internal/repositories"
)

// StageService handles stage-related business logic
type StageService struct {
	stageRepository            *repositories.StageRepository
	stageMetricRepository      *repositories.StageMetricRepository
	competenceMetricRepository *repositories.CompetenceMetricRepository
}

// NewStageService creates a new stage service
func NewStageService(
	stageRepository *repositories.StageRepository,
	stageMetricRepository *repositories.StageMetricRepository,
	competenceMetricRepository *repositories.CompetenceMetricRepository,
) *StageService {
	return &StageService{
		stageRepository:            stageRepository,
		stageMetricRepository:      stageMetricRepository,
		competenceMetricRepository: competenceMetricRepository,
	}
}

// GetStageByID retrieves a stage by its ID
func (s *StageService) GetStageByID(stageID int) (*models.Stage, error) {
	stage, err := s.stageRepository.FindByID(stageID)
	if err != nil {
		return nil, fmt.Errorf("error retrieving stage: %w", err)
	}

	if stage == nil {
		return nil, errors.New("stage not found")
	}

	return stage, nil
}

// CreateStage creates a new stage
func (s *StageService) CreateStage(stage *models.Stage) error {
	if stage.GameID == "" || stage.StageKey == "" || stage.StageName == "" {
		return errors.New("game ID, stage key, and name are required")
	}

	return s.stageRepository.Create(stage)
}

// UpdateStage updates an existing stage
func (s *StageService) UpdateStage(stage *models.Stage) error {
	if stage.StageID == 0 {
		return errors.New("stage ID is required")
	}

	existingStage, err := s.stageRepository.FindByID(stage.StageID)
	if err != nil {
		return fmt.Errorf("error checking existing stage: %w", err)
	}

	if existingStage == nil {
		return errors.New("stage not found")
	}

	return s.stageRepository.Update(stage)
}

// DeleteStage deletes a stage
func (s *StageService) DeleteStage(stageID int) error {
	existingStage, err := s.stageRepository.FindByID(stageID)
	if err != nil {
		return fmt.Errorf("error checking existing stage: %w", err)
	}

	if existingStage == nil {
		return errors.New("stage not found")
	}

	return s.stageRepository.Delete(stageID)
}

// AssociateMetricWithStage associates a metric with a stage
func (s *StageService) AssociateMetricWithStage(stageID, metricID int) error {
	stageMetric := &models.Metric{
		StageID:  stageID,
		MetricID: metricID,
	}

	return s.stageMetricRepository.Create(stageMetric)
}

// RemoveMetricFromStage removes a metric from a stage
func (s *StageService) RemoveMetricFromStage(stageID, metricID int) error {
	return s.stageMetricRepository.Delete(stageID, metricID)
}

// GetStageMetrics retrieves all metrics associated with a stage
func (s *StageService) GetStageMetrics(stageID int) ([]models.Metric, error) {
	// First get all stage-metric associations
	stageMetrics, err := s.stageMetricRepository.FindByStage(stageID)
	if err != nil {
		return nil, fmt.Errorf("error retrieving stage metrics: %w", err)
	}

	// If no metrics are associated, return empty slice
	if len(stageMetrics) == 0 {
		return []models.Metric{}, nil
	}

	return stageMetrics, nil
}

// GetStageMetrics retrieves all metrics associated with a stage
func (s *StageService) GetCompetenceMetrics(competenceID int) ([]models.Metric, error) {
	// First get all stage-metric associations
	competenceMetrics, err := s.stageMetricRepository.FindByCompetence(competenceID)
	if err != nil {
		return nil, fmt.Errorf("error retrieving stage metrics: %w", err)
	}

	// If no metrics are associated, return empty slice
	if len(competenceMetrics) == 0 {
		return []models.Metric{}, nil
	}

	return competenceMetrics, nil
}

// GetRequiredParametersForStage retrieves all required parameters for a stage
func (s *StageService) GetRequiredParametersForStage(stageID int) ([]models.MetricParameter, error) {
	// Get all metrics associated with the stage
	stageMetrics, err := s.stageMetricRepository.FindByStage(stageID)
	if err != nil {
		return nil, fmt.Errorf("error retrieving stage metrics: %w", err)
	}

	// If no metrics are associated, return empty slice
	if len(stageMetrics) == 0 {
		return []models.MetricParameter{}, nil
	}

	// Collect all required parameters from all metrics
	var allParameters []models.MetricParameter

	// for _, sm := range stageMetrics {
	// 	// For each metric, get its required parameters
	// 	metricID := sm.MetricID

	// 	// Use the metric parameter repository to get required parameters
	// 	// Since we don't have direct access to it, we'll need to use the competence metric repository
	// 	// to get the metric first, then extract its parameters
	// 	metric, err := s.competenceMetricRepository.FindWithParameters(metricID)
	// 	if err != nil {
	// 		return nil, fmt.Errorf("error retrieving metric parameters: %w", err)
	// 	}

	// 	if metric != nil && len(metric.Parameters) > 0 {
	// 		// Filter for required parameters only
	// 		for _, param := range metric.Parameters {
	// 			if param.IsRequired {
	// 				allParameters = append(allParameters, param)
	// 			}
	// 		}
	// 	}
	// }

	return allParameters, nil
}
