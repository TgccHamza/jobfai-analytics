package resolver

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

import (
	"jobfai-analytics/internal/database"
	"jobfai-analytics/internal/repositories"
	"jobfai-analytics/internal/services"
	"jobfai-analytics/internal/subscription"
)

type Resolver struct {
	DB                       database.Service
	GameService              *services.GameService
	CompetenceService        *services.CompetenceService
	MetricService            *services.MetricService
	ConstantParameterService *services.ConstantParameterService
	StageService             *services.StageService
	PlayerPerformanceService *services.PlayerPerformanceService
	subscriptionManager      *subscription.Manager
}

// NewResolver creates a new resolver with database access and services
func NewResolver(db database.Service) *Resolver {
	// Create repositories directly using the database connection
	gameRepo := repositories.NewGameRepository(db.DB())
	competenceRepo := repositories.NewCompetenceRepository(db.DB())
	competenceMetricRepo := repositories.NewCompetenceMetricRepository(db.DB())
	metricParamRepo := repositories.NewMetricParameterRepository(db.DB())
	stageRepo := repositories.NewStageRepository(db.DB())
	stageMetricRepo := repositories.NewStageMetricRepository(db.DB())
	gameMetricRepo := repositories.NewGameMetricRepository(db.DB())
	gameMetricParamRepo := repositories.NewGameMetricParameterRepository(db.DB())
	constantParamRepo := repositories.NewConstantParameterRepository(db.DB())

	// Create services
	gameService := services.NewGameService(gameRepo, stageRepo, competenceRepo)
	competenceService := services.NewCompetenceService(competenceRepo, competenceMetricRepo)
	metricService := services.NewMetricService(
		competenceMetricRepo,
		metricParamRepo,
		gameMetricRepo,
		gameMetricParamRepo,
		constantParamRepo,
		nil, // FormulaEvaluator will be initialized in the service
	)
	stageService := services.NewStageService(stageRepo, stageMetricRepo, competenceMetricRepo)
	playerPerformanceService := services.NewPlayerPerformanceService(
		gameRepo,
		stageRepo,
		competenceRepo,
		competenceMetricRepo,
		gameMetricRepo,
		constantParamRepo,
		nil, // MetricCalculator will be initialized in the service
	)
	constantParamService := services.NewConstantParameterService(constantParamRepo)

	return &Resolver{
		DB:                       db,
		GameService:              gameService,
		CompetenceService:        competenceService,
		MetricService:            metricService,
		StageService:             stageService,
		PlayerPerformanceService: playerPerformanceService,
		ConstantParameterService: constantParamService,
		subscriptionManager:      subscription.NewManager(),
	}
}
