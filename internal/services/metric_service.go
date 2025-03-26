package services

import (
	"errors"
	"fmt"

	"jobfai-analytics/internal/models"
	"jobfai-analytics/internal/repositories"
	"jobfai-analytics/pkg/evaluator"
)

// MetricService handles metric-related business logic
type MetricService struct {
	competenceMetricRepository    *repositories.CompetenceMetricRepository
	metricParameterRepository     *repositories.MetricParameterRepository
	gameMetricRepository          *repositories.GameMetricRepository
	gameMetricParameterRepository *repositories.GameMetricParameterRepository
	constantParameterRepository   *repositories.ConstantParameterRepository
	formulaEvaluator              *evaluator.FormulaEvaluator
}

// NewMetricService creates a new metric service
func NewMetricService(
	competenceMetricRepository *repositories.CompetenceMetricRepository,
	metricParameterRepository *repositories.MetricParameterRepository,
	gameMetricRepository *repositories.GameMetricRepository,
	gameMetricParameterRepository *repositories.GameMetricParameterRepository,
	constantParameterRepository *repositories.ConstantParameterRepository,
	formulaEvaluator *evaluator.FormulaEvaluator,
) *MetricService {
	return &MetricService{
		competenceMetricRepository:    competenceMetricRepository,
		metricParameterRepository:     metricParameterRepository,
		gameMetricRepository:          gameMetricRepository,
		gameMetricParameterRepository: gameMetricParameterRepository,
		constantParameterRepository:   constantParameterRepository,
		formulaEvaluator:              formulaEvaluator,
	}
}

// GetCompetenceMetricByID retrieves a competence metric by its ID
func (s *MetricService) GetCompetenceMetricByID(metricID int) (*models.CompetenceMetric, error) {
	metric, err := s.competenceMetricRepository.FindByID(metricID)
	if err != nil {
		return nil, fmt.Errorf("error retrieving competence metric: %w", err)
	}

	if metric == nil {
		return nil, errors.New("competence metric not found")
	}

	return metric, nil
}

// GetCompetenceMetricWithParameters retrieves a competence metric with all its parameters
func (s *MetricService) GetCompetenceMetricWithParameters(metricID int) (*models.CompetenceMetric, error) {
	metric, err := s.competenceMetricRepository.FindWithParameters(metricID)
	if err != nil {
		return nil, fmt.Errorf("error retrieving competence metric with parameters: %w", err)
	}

	if metric == nil {
		return nil, errors.New("competence metric not found")
	}

	return metric, nil
}

// CreateCompetenceMetric creates a new competence metric
func (s *MetricService) CreateCompetenceMetric(metric *models.CompetenceMetric) error {
	if metric.CompetenceID == 0 || metric.MetricKey == "" || metric.MetricName == "" {
		return errors.New("competence ID, metric key, and name are required")
	}

	// Validate formula
	// if metric.Formula != "" {
	// 	if _, err := s.formulaEvaluator.CompileFormula(metric.Formula); err != nil {
	// 		return fmt.Errorf("invalid formula: %w", err)
	// 	}
	// }

	return s.competenceMetricRepository.Create(metric)
}

// UpdateCompetenceMetric updates an existing competence metric
func (s *MetricService) UpdateCompetenceMetric(metric *models.CompetenceMetric) error {
	if metric.MetricID == 0 {
		return errors.New("metric ID is required")
	}

	existingMetric, err := s.competenceMetricRepository.FindByID(metric.MetricID)
	if err != nil {
		return fmt.Errorf("error checking existing metric: %w", err)
	}

	if existingMetric == nil {
		return errors.New("competence metric not found")
	}

	// Validate formula
	// if metric.Formula != "" {
	// 	if _, err := s.formulaEvaluator.CompileFormula(metric.Formula); err != nil {
	// 		return fmt.Errorf("invalid formula: %w", err)
	// 	}
	// }

	return s.competenceMetricRepository.Update(metric)
}

// DeleteCompetenceMetric deletes a competence metric
func (s *MetricService) DeleteCompetenceMetric(metricID int) error {
	existingMetric, err := s.competenceMetricRepository.FindByID(metricID)
	if err != nil {
		return fmt.Errorf("error checking existing metric: %w", err)
	}

	if existingMetric == nil {
		return errors.New("competence metric not found")
	}

	return s.competenceMetricRepository.Delete(metricID)
}

// CreateCompetenceMetricParameter adds a parameter to a competence metric
func (s *MetricService) CreateCompetenceMetricParameter(parameter *models.CompetenceMetricParameter) error {
	if parameter.MetricID == 0 || parameter.ParamKey == "" || parameter.ParamName == "" {
		return errors.New("metric ID, parameter key, and name are required")
	}

	return s.metricParameterRepository.Create(parameter)
}

// UpdateMetricParameter updates an existing metric parameter
func (s *MetricService) UpdateMetricParameter(parameter *models.CompetenceMetricParameter) error {
	if parameter.ParamID == 0 {
		return errors.New("parameter ID is required")
	}

	existingParameter, err := s.metricParameterRepository.FindByID(parameter.ParamID)
	if err != nil {
		return fmt.Errorf("error checking existing parameter: %w", err)
	}

	if existingParameter == nil {
		return errors.New("parameter not found")
	}

	return s.metricParameterRepository.Update(parameter)
}

// DeleteMetricParameter deletes a metric parameter
func (s *MetricService) DeleteMetricParameter(paramID int) error {
	existingParameter, err := s.metricParameterRepository.FindByID(paramID)
	if err != nil {
		return fmt.Errorf("error checking existing parameter: %w", err)
	}

	if existingParameter == nil {
		return errors.New("parameter not found")
	}

	return s.metricParameterRepository.Delete(paramID)
}

// GetGameMetricByID retrieves a game metric by its ID
func (s *MetricService) GetGameMetricByID(metricID int) (*models.GameMetric, error) {
	metric, err := s.gameMetricRepository.FindByID(metricID)
	if err != nil {
		return nil, fmt.Errorf("error retrieving game metric: %w", err)
	}

	if metric == nil {
		return nil, errors.New("game metric not found")
	}

	return metric, nil
}

// GetGameMetricWithParameters retrieves a game metric with all its parameters
func (s *MetricService) GetGameMetricWithParameters(metricID int) (*models.GameMetric, error) {
	metric, err := s.gameMetricRepository.FindWithParameters(metricID)
	if err != nil {
		return nil, fmt.Errorf("error retrieving game metric with parameters: %w", err)
	}

	if metric == nil {
		return nil, errors.New("game metric not found")
	}

	return metric, nil
}

// CreateGameMetric creates a new game metric
func (s *MetricService) CreateGameMetric(metric *models.GameMetric) error {
	if metric.GameID == "" || metric.MetricKey == "" || metric.MetricName == "" {
		return errors.New("game ID, metric key, and name are required")
	}

	// Validate formula
	if metric.Formula != "" {
		if _, err := s.formulaEvaluator.CompileFormula(metric.Formula); err != nil {
			return fmt.Errorf("invalid formula: %w", err)
		}
	}

	return s.gameMetricRepository.Create(metric)
}

// UpdateGameMetric updates an existing game metric
func (s *MetricService) UpdateGameMetric(metric *models.GameMetric) error {
	if metric.MetricID == 0 {
		return errors.New("metric ID is required")
	}

	existingMetric, err := s.gameMetricRepository.FindByID(metric.MetricID)
	if err != nil {
		return fmt.Errorf("error checking existing metric: %w", err)
	}

	if existingMetric == nil {
		return errors.New("game metric not found")
	}

	// Validate formula
	if metric.Formula != "" {
		if _, err := s.formulaEvaluator.CompileFormula(metric.Formula); err != nil {
			return fmt.Errorf("invalid formula: %w", err)
		}
	}

	return s.gameMetricRepository.Update(metric)
}

// DeleteGameMetric deletes a game metric
func (s *MetricService) DeleteGameMetric(metricID int) error {
	existingMetric, err := s.gameMetricRepository.FindByID(metricID)
	if err != nil {
		return fmt.Errorf("error checking existing metric: %w", err)
	}

	if existingMetric == nil {
		return errors.New("game metric not found")
	}

	return s.gameMetricRepository.Delete(metricID)
}

// CreateGameMetricParameter adds a parameter to a game metric
func (s *MetricService) CreateGameMetricParameter(parameter *models.GameMetricParameter) error {
	if parameter.MetricID == 0 || parameter.ParamKey == "" || parameter.ParamName == "" {
		return errors.New("metric ID, parameter key, and name are required")
	}

	return s.gameMetricParameterRepository.Create(parameter)
}

// UpdateGameMetricParameter updates an existing game metric parameter
func (s *MetricService) UpdateGameMetricParameter(parameter *models.GameMetricParameter) error {
	if parameter.ParamID == 0 {
		return errors.New("parameter ID is required")
	}

	existingParameter, err := s.gameMetricParameterRepository.FindByID(parameter.ParamID)
	if err != nil {
		return fmt.Errorf("error checking existing parameter: %w", err)
	}

	if existingParameter == nil {
		return errors.New("parameter not found")
	}

	return s.gameMetricParameterRepository.Update(parameter)
}

// DeleteGameMetricParameter deletes a game metric parameter
func (s *MetricService) DeleteGameMetricParameter(paramID int) error {
	existingParameter, err := s.gameMetricParameterRepository.FindByID(paramID)
	if err != nil {
		return fmt.Errorf("error checking existing parameter: %w", err)
	}

	if existingParameter == nil {
		return errors.New("parameter not found")
	}

	return s.gameMetricParameterRepository.Delete(paramID)
}

// GetCompetenceMetricParameters retrieves all parameters for a competence metric
func (s *MetricService) GetCompetenceMetricParameters(metricID int) ([]models.CompetenceMetricParameter, error) {
	parameters, err := s.metricParameterRepository.FindByMetric(metricID)
	if err != nil {
		return nil, fmt.Errorf("error retrieving competence metric parameters: %w", err)
	}

	return parameters, nil
}

// GetGameMetricParameters retrieves all parameters for a game metric
func (s *MetricService) GetGameMetricParameters(metricID int) ([]models.GameMetricParameter, error) {
	parameters, err := s.gameMetricParameterRepository.FindByMetric(metricID)
	if err != nil {
		return nil, fmt.Errorf("error retrieving game metric parameters: %w", err)
	}

	return parameters, nil
}

// GetGameMetricsByGame retrieves all game metrics for a game
func (s *MetricService) GetGameMetricsByGame(gameID string) ([]models.GameMetric, error) {
	metrics, err := s.gameMetricRepository.FindByGame(gameID)
	if err != nil {
		return nil, fmt.Errorf("error retrieving game metrics: %w", err)
	}

	return metrics, nil
}

// GetConstantParametersByGame retrieves all constant parameters for a game
func (s *MetricService) GetConstantParametersByGame(gameID string) ([]models.ConstantParameter, error) {
	constants, err := s.constantParameterRepository.FindByGame(gameID)
	if err != nil {
		return nil, fmt.Errorf("error retrieving constant parameters: %w", err)
	}

	return constants, nil
}

// GetConstantParameterByID retrieves a constant parameter by its ID
func (s *MetricService) GetConstantParameterByID(constID int) (*models.ConstantParameter, error) {
	constant, err := s.constantParameterRepository.FindByID(constID)
	if err != nil {
		return nil, fmt.Errorf("error retrieving constant parameter: %w", err)
	}

	if constant == nil {
		return nil, errors.New("constant parameter not found")
	}

	return constant, nil
}

// GetGameMetricParameterByID retrieves a game metric parameter by its ID
func (s *MetricService) GetGameMetricParameterByID(paramID int) (*models.GameMetricParameter, error) {
	parameter, err := s.gameMetricParameterRepository.FindByID(paramID)
	if err != nil {
		return nil, fmt.Errorf("error retrieving game metric parameter: %w", err)
	}

	if parameter == nil {
		return nil, errors.New("game metric parameter not found")
	}

	return parameter, nil
}

// GetCompetenceMetricParameterByID retrieves a competence metric parameter by its ID
func (s *MetricService) GetCompetenceMetricParameterByID(paramID int) (*models.CompetenceMetricParameter, error) {
	parameter, err := s.metricParameterRepository.FindByID(paramID)
	if err != nil {
		return nil, fmt.Errorf("error retrieving competence metric parameter: %w", err)
	}

	if parameter == nil {
		return nil, errors.New("competence metric parameter not found")
	}

	return parameter, nil
}
