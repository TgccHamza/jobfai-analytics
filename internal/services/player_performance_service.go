package services

import (
	"errors"
	"fmt"
	"time"

	"jobfai-analytics/internal/models"
	"jobfai-analytics/internal/repositories"
	"jobfai-analytics/pkg/calculator"
)

// PlayerPerformanceInput represents the input data for calculating player performance
type PlayerPerformanceInput struct {
	PlayerID string `json:"playerId"`
	GameID   string `json:"gameId"`
	DbIndex  string `json:"dbIndex"`
}

// PlayerPerformanceOutput represents the output data for player performance
type PlayerPerformanceOutput struct {
	PlayerID          string                   `json:"playerId"`
	PlayerName        string                   `json:"playerName"`
	ProfileType       string                   `json:"profileType"`
	GameDate          time.Time                `json:"gameDate"`
	GameID            string                   `json:"gameId"`
	TotalScore        float64                  `json:"totalScore"`
	TotalTimeTaken    float64                  `json:"totalTimeTaken"`
	CompetenceDetails map[string]interface{}   `json:"competenceDetails"`
	StagePerformance  []map[string]interface{} `json:"stagePerformance"`
	GlobalMetrics     []map[string]interface{} `json:"globalMetrics"`
}

// PlayerPerformanceService handles player performance calculation
type PlayerPerformanceService struct {
	gameRepository              *repositories.GameRepository
	stageRepository             *repositories.StageRepository
	competenceRepository        *repositories.CompetenceRepository
	MetricRepository            *repositories.CompetenceMetricRepository
	MetricParameterRepository   *repositories.MetricParameterRepository
	gameMetricRepository        *repositories.GameMetricRepository
	constantParameterRepository *repositories.ConstantParameterRepository
	metricCalculator            *calculator.MetricCalculator
}

// NewPlayerPerformanceService creates a new player performance service
func NewPlayerPerformanceService(
	gameRepository *repositories.GameRepository,
	stageRepository *repositories.StageRepository,
	competenceRepository *repositories.CompetenceRepository,
	MetricRepository *repositories.CompetenceMetricRepository,
	MetricParameterRepository *repositories.MetricParameterRepository,
	gameMetricRepository *repositories.GameMetricRepository,
	constantParameterRepository *repositories.ConstantParameterRepository,
	metricCalculator *calculator.MetricCalculator,
) *PlayerPerformanceService {
	return &PlayerPerformanceService{
		gameRepository:              gameRepository,
		stageRepository:             stageRepository,
		competenceRepository:        competenceRepository,
		MetricRepository:            MetricRepository,
		gameMetricRepository:        gameMetricRepository,
		constantParameterRepository: constantParameterRepository,
		metricCalculator:            metricCalculator,
		MetricParameterRepository:   MetricParameterRepository,
	}
}

// CalculatePlayerPerformance calculates the performance of a player
func (s *PlayerPerformanceService) CalculatePlayerPerformance(input PlayerPerformanceInput) (*PlayerPerformanceOutput, error) {
	// Validate input data
	if err := s.validatePlayerData(input); err != nil {
		return nil, err
	}

	// // Get game configuration
	game, err := s.gameRepository.FindByID(input.GameID)
	if err != nil {
		return nil, fmt.Errorf("error retrieving game: %w", err)
	}

	if game == nil {
		return nil, errors.New("game not found")
	}

	playerData, err := GetPlayerData(input.DbIndex, input.PlayerID)

	if err != nil {
		return nil, fmt.Errorf("error retrieving player data: %w", err)
	}

	// Get constants
	constants, err := s.constantParameterRepository.FindByGame(input.GameID)
	if err != nil {
		return nil, fmt.Errorf("error retrieving constants: %w", err)
	}

	// Calculate stage metrics
	stagePerformance, err := s.calculateStagePerformance(game, constants, playerData)
	if err != nil {
		return nil, fmt.Errorf("error calculating stage performance: %w", err)
	}

	// Calculate global metrics
	globalMetrics, err := s.calculateGlobalMetrics(game, constants, playerData)
	if err != nil {
		return nil, fmt.Errorf("error calculating global metrics: %w", err)
	}

	// Calculate competence metrics
	competenceDetails, err := s.calculateCompetenceDetails(game, stagePerformance, globalMetrics)
	if err != nil {
		return nil, fmt.Errorf("error calculating competence details: %w", err)
	}

	// Calculate total score
	totalScore := s.calculateTotalScore(competenceDetails)

	// Calculate total time taken
	totalTimeTaken := s.calculateTotalTimeTaken(stagePerformance)

	// Build and return the complete player performance data
	return &PlayerPerformanceOutput{
		PlayerID:          input.PlayerID,
		PlayerName:        "",
		ProfileType:       "",
		GameDate:          time.Now(),
		GameID:            input.GameID,
		TotalScore:        totalScore,
		TotalTimeTaken:    totalTimeTaken,
		CompetenceDetails: competenceDetails,
		StagePerformance:  stagePerformance,
		GlobalMetrics:     globalMetrics,
	}, nil
}

// validatePlayerData validates the player input data
func (s *PlayerPerformanceService) validatePlayerData(input PlayerPerformanceInput) error {
	if input.PlayerID == "" || input.DbIndex == "" || input.GameID == "" {
		return errors.New("missing required player data")
	}

	return nil
}

// calculateStagePerformance calculates the performance for each stage
func (s *PlayerPerformanceService) calculateStagePerformance(
	game *models.Game,
	constants []models.ConstantParameter,
	playerData map[string]interface{},
) ([]map[string]interface{}, error) {
	stagePerformance := make([]map[string]interface{}, 0)
	stages, err := s.stageRepository.FindByGame(game.GameID)
	if err != nil {
		return nil, fmt.Errorf("error retrieving stages: %w", err)
	}

	for _, stage := range stages {
		stageMetrics, err := s.MetricRepository.FindByStage(stage.StageID)
		if err != nil {
			return nil, fmt.Errorf("error retrieving stage: %w", err)
		}

		resultMetrics := make([]map[string]interface{}, 0)

		for _, stageMetric := range stageMetrics {
			parameters, err := s.MetricParameterRepository.FindByMetric(stageMetric.MetricID)
			if err != nil {
				return nil, fmt.Errorf("error retrieving metric: %w", err)
			}

			competence, err := s.competenceRepository.FindByID(stageMetric.CompetenceID)
			if err != nil {
				return nil, fmt.Errorf("error retrieving competence: %w", err)
			}

			if competence == nil {
				continue
			}

			metricResult, err := s.metricCalculator.CalculateMetric(
				&stageMetric,
				parameters,
				constants,
				playerData)

			if err != nil {
				return nil, fmt.Errorf("error calculating metric: %w", err)
			}

			resultMetrics = append(resultMetrics, map[string]interface{}{
				"kpiId":     stageMetric.MetricKey,
				"kpiName":   stageMetric.MetricName,
				"category":  competence.CompetenceKey,
				"value":     metricResult["value"],
				"benchmark": stageMetric.Benchmark,
				"formula":   stageMetric.Formula,
				"rawData":   metricResult["rawData"],
			})
		}

		// Calculate stage score (simple average of metrics)
		var stageScore float64
		if len(resultMetrics) > 0 {
			var totalScore float64
			for _, metric := range resultMetrics {
				value, ok := metric["value"].(float64)
				if ok {
					totalScore += value
				}
			}
			stageScore = totalScore / float64(len(resultMetrics))
		}

		stagePerformance = append(stagePerformance, map[string]interface{}{
			"stageId":          stage.StageID,
			"stageName":        stage.StageName,
			"metrics":          resultMetrics,
			"timeTaken":        0,
			"optimalTime":      stage.OptimalTime,
			"score":            stageScore,
			"benchmark":        stage.Benchmark,
			"completionStatus": "completed",
		})
	}

	return stagePerformance, nil
}

// calculateCompetenceDetails calculates the details for each competence
func (s *PlayerPerformanceService) calculateCompetenceDetails(
	game *models.Game,
	stagePerformance []map[string]interface{},
	globalMetrics []map[string]interface{},
) (map[string]interface{}, error) {
	competenceDetails := make(map[string]interface{})
	Metrics := make(map[string][]map[string]interface{})

	// Group metrics by competence
	for _, stage := range stagePerformance {
		metrics, ok := stage["metrics"].([]map[string]interface{})
		if !ok {
			continue
		}

		for _, metric := range metrics {
			category, ok := metric["category"].(string)
			if !ok {
				continue
			}

			if _, exists := Metrics[category]; !exists {
				Metrics[category] = make([]map[string]interface{}, 0)
			}

			Metrics[category] = append(Metrics[category], metric)
		}
	}

	for _, metric := range globalMetrics {
		category, ok := metric["category"].(string)
		if !ok {
			continue
		}

		if _, exists := Metrics[category]; !exists {
			Metrics[category] = make([]map[string]interface{}, 0)
		}

		Metrics[category] = append(Metrics[category], metric)
	}

	competencies, err := s.competenceRepository.FindByGame(game.GameID)

	if err != nil {
		return nil, fmt.Errorf("error retrieving competencies: %w", err)
	}

	// Calculate competence scores
	for _, competence := range competencies {
		competenceKey := competence.CompetenceKey
		metrics, exists := Metrics[competenceKey]
		if !exists {
			continue
		}

		// Calculate competence score (simple average of metrics)
		var competenceScore float64
		if len(metrics) > 0 {
			var totalScore float64
			for _, metric := range metrics {
				value, ok := metric["value"].(float64)
				if ok {
					totalScore += value
				}
			}
			competenceScore = totalScore / float64(len(metrics))
		}

		metricDetails := make([]map[string]interface{}, 0)
		for _, metric := range metrics {
			metricDetails = append(metricDetails, map[string]interface{}{
				"kpiId":     metric["kpiId"],
				"kpiName":   metric["kpiName"],
				"value":     metric["value"],
				"benchmark": metric["benchmark"],
			})
		}

		competenceDetails[competenceKey] = map[string]interface{}{
			"name":      competence.CompetenceName,
			"score":     competenceScore,
			"benchmark": competence.Benchmark,
			"weight":    competence.Weight,
			"metrics":   metricDetails,
		}
	}

	return competenceDetails, nil
}

// calculateGlobalMetrics calculates the global metrics for the game
func (s *PlayerPerformanceService) calculateGlobalMetrics(
	game *models.Game,
	constants []models.ConstantParameter,
	playerData map[string]interface{},
) ([]map[string]interface{}, error) {
	// Convert competenceDetails to the expected format
	gameMetrics, err := s.gameRepository.FindByGame(game.GameID)
	if err != nil {
		return nil, fmt.Errorf("error retrieving metrics: %w", err)
	}

	resultMetrics := make([]map[string]interface{}, 0)
	for _, gameMetric := range gameMetrics {
		parameters, err := s.MetricParameterRepository.FindByMetric(gameMetric.MetricID)
		if err != nil {
			return nil, fmt.Errorf("error retrieving metric: %w", err)
		}

		competence, err := s.competenceRepository.FindByID(gameMetric.CompetenceID)
		if err != nil {
			return nil, fmt.Errorf("error retrieving competence: %w", err)
		}

		// Safely handle the case where competence might be nil
		var competenceKey string
		if competence != nil {
			competenceKey = competence.CompetenceKey
		}

		metricResult, err := s.metricCalculator.CalculateMetric(
			&gameMetric,
			parameters,
			constants,
			playerData)

		if err != nil {
			return nil, fmt.Errorf("error calculating metric: %w", err)
		}
		resultMetrics = append(resultMetrics, map[string]interface{}{
			"kpiId":     gameMetric.MetricKey,
			"kpiName":   gameMetric.MetricName,
			"category":  competenceKey,
			"value":     metricResult["value"],
			"benchmark": gameMetric.Benchmark,
			"formula":   gameMetric.Formula,
			"rawData":   metricResult["rawData"],
		})
	}

	return resultMetrics, nil
}

// calculateTotalScore calculates the total score based on competence scores and weights
func (s *PlayerPerformanceService) calculateTotalScore(competenceDetails map[string]interface{}) float64 {
	var totalScore, totalWeight float64

	for _, details := range competenceDetails {
		if competence, ok := details.(map[string]interface{}); ok {
			score, scoreOk := competence["score"].(float64)
			weight, weightOk := competence["weight"].(float64)

			if scoreOk && weightOk {
				totalScore += score * weight
				totalWeight += weight
			}
		}
	}

	if totalWeight == 0 {
		return 0
	}

	return totalScore / totalWeight
}

// calculateTotalTimeTaken calculates the total time taken across all stages
func (s *PlayerPerformanceService) calculateTotalTimeTaken(stagePerformance []map[string]interface{}) float64 {
	var totalTime float64

	for _, stage := range stagePerformance {
		timeTaken, ok := stage["timeTaken"].(float64)
		if ok {
			totalTime += timeTaken
		}
	}

	return totalTime
}
