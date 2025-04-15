package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.68

import (
	"context"
	"fmt"
	"jobfai-analytics/internal/graph"
	"jobfai-analytics/internal/models"
	"time"
)

// CreatedAt is the resolver for the createdAt field.
func (r *competenceResolver) CreatedAt(ctx context.Context, obj *models.Competence) (*string, error) {
	timeStr := obj.CreatedAt.Format(time.RFC3339)
	return &timeStr, nil
}

// UpdatedAt is the resolver for the updatedAt field.
func (r *competenceResolver) UpdatedAt(ctx context.Context, obj *models.Competence) (*string, error) {
	timeStr := obj.UpdatedAt.Format(time.RFC3339)
	return &timeStr, nil
}

// Metrics is the resolver for the metrics field.
func (r *competenceResolver) Metrics(ctx context.Context, obj *models.Competence) ([]*models.Metric, error) {
	metrics, err := r.StageService.GetCompetenceMetrics(obj.CompetenceID)
	if err != nil {
		return nil, fmt.Errorf("error fetching metrics: %w", err)
	}

	// Convert []models.Metric to []*models.Metric
	result := make([]*models.Metric, len(metrics))
	for i := range metrics {
		result[i] = &metrics[i]
	}

	return result, nil
}

// Game is the resolver for the game field.
func (r *competenceResolver) Game(ctx context.Context, obj *models.Competence) (*models.Game, error) {
	return r.GameService.GetGameByID(obj.GameID)
}

// ParentCompetence is the resolver for the parentCompetence field.
func (r *competenceResolver) ParentCompetence(ctx context.Context, obj *models.Competence) (*models.Competence, error) {
	return r.CompetenceService.GetCompetenceByID(*obj.ParentID)
}

// CreatedAt is the resolver for the createdAt field.
func (r *constantParameterResolver) CreatedAt(ctx context.Context, obj *models.ConstantParameter) (*string, error) {
	timeStr := obj.CreatedAt.Format(time.RFC3339)
	return &timeStr, nil
}

// UpdatedAt is the resolver for the updatedAt field.
func (r *constantParameterResolver) UpdatedAt(ctx context.Context, obj *models.ConstantParameter) (*string, error) {
	timeStr := obj.UpdatedAt.Format(time.RFC3339)
	return &timeStr, nil
}

// Game is the resolver for the game field.
func (r *constantParameterResolver) Game(ctx context.Context, obj *models.ConstantParameter) (*models.Game, error) {
	return r.GameService.GetGameByID(obj.GameID)
}

// CreatedAt is the resolver for the createdAt field.
func (r *gameResolver) CreatedAt(ctx context.Context, obj *models.Game) (*string, error) {
	timeStr := obj.CreatedAt.Format(time.RFC3339)
	return &timeStr, nil
}

// UpdatedAt is the resolver for the updatedAt field.
func (r *gameResolver) UpdatedAt(ctx context.Context, obj *models.Game) (*string, error) {
	timeStr := obj.UpdatedAt.Format(time.RFC3339)
	return &timeStr, nil
}

// Competencies is the resolver for the competencies field.
func (r *gameResolver) Competencies(ctx context.Context, obj *models.Game) ([]*models.Competence, error) {
	competencies, err := r.GameService.GetGameCompetencies(obj.GameID)
	if err != nil {
		return nil, fmt.Errorf("error fetching competencies: %w", err)
	}

	// Convert []models.Competence to []*models.Competence
	result := make([]*models.Competence, len(competencies))
	for i := range competencies {
		result[i] = &competencies[i]
	}

	return result, nil
}

// Stages is the resolver for the stages field.
func (r *gameResolver) Stages(ctx context.Context, obj *models.Game) ([]*models.Stage, error) {
	stages, err := r.GameService.GetGameStages(obj.GameID)
	if err != nil {
		return nil, fmt.Errorf("error fetching stages: %w", err)
	}

	// Convert []models.Stage to []*models.Stage
	result := make([]*models.Stage, len(stages))
	for i := range stages {
		result[i] = &stages[i]
	}

	return result, nil
}

// GameMetrics is the resolver for the gameMetrics field.
func (r *gameResolver) GameMetrics(ctx context.Context, obj *models.Game) ([]*models.Metric, error) {
	metrics, err := r.GameService.GetGameMetrics(obj.GameID)
	if err != nil {
		return nil, fmt.Errorf("error fetching metrics: %w", err)
	}

	// Convert []models.Metric to []*models.Metric
	result := make([]*models.Metric, len(metrics))
	for i := range metrics {
		result[i] = &metrics[i]
	}

	return result, nil
}

// ConstantParameters is the resolver for the constantParameters field.
func (r *gameResolver) ConstantParameters(ctx context.Context, obj *models.Game) ([]*models.ConstantParameter, error) {
	constants, err := r.MetricService.GetConstantParametersByGame(obj.GameID)
	if err != nil {
		return nil, fmt.Errorf("error fetching constants: %w", err)
	}

	// Convert []models.ConstantParameter to []*models.ConstantParameter
	result := make([]*models.ConstantParameter, len(constants))
	for i := range constants {
		result[i] = &constants[i]
	}

	return result, nil
}

// CreatedAt is the resolver for the createdAt field.
func (r *metricResolver) CreatedAt(ctx context.Context, obj *models.Metric) (*string, error) {
	timeStr := obj.CreatedAt.Format(time.RFC3339)
	return &timeStr, nil
}

// UpdatedAt is the resolver for the updatedAt field.
func (r *metricResolver) UpdatedAt(ctx context.Context, obj *models.Metric) (*string, error) {
	timeStr := obj.UpdatedAt.Format(time.RFC3339)
	return &timeStr, nil
}

// Parameters is the resolver for the parameters field.
func (r *metricResolver) Parameters(ctx context.Context, obj *models.Metric) ([]*models.MetricParameter, error) {
	parameters, err := r.MetricService.GetMetricParameters(obj.MetricID)
	if err != nil {
		return nil, fmt.Errorf("error fetching parameters: %w", err)
	}

	// Convert []models.MetricParameter to []*model.MetricParameter
	result := make([]*models.MetricParameter, len(parameters))
	for i, param := range parameters {
		result[i] = &models.MetricParameter{
			ParamID:      param.ParamID,
			MetricID:     param.MetricID,
			ParamName:    param.ParamName,
			ParamKey:     param.ParamKey,
			Description:  param.Description,
			DefaultValue: param.DefaultValue,
			ParamType:    param.ParamType,
			IsRequired:   param.IsRequired,
		}
	}

	return result, nil
}

// Competence is the resolver for the competence field.
func (r *metricResolver) Competence(ctx context.Context, obj *models.Metric) (*models.Competence, error) {
	return r.CompetenceService.GetCompetenceByID(obj.CompetenceID)
}

// Stage is the resolver for the stage field.
func (r *metricResolver) Stage(ctx context.Context, obj *models.Metric) (*models.Stage, error) {
	return r.StageService.GetStageByID(obj.StageID)
}

// Game is the resolver for the game field.
func (r *metricResolver) Game(ctx context.Context, obj *models.Metric) (*models.Game, error) {
	return r.GameService.GetGameByID(obj.GameID)
}

// CreatedAt is the resolver for the createdAt field.
func (r *metricParameterResolver) CreatedAt(ctx context.Context, obj *models.MetricParameter) (*string, error) {
	timeStr := obj.CreatedAt.Format(time.RFC3339)
	return &timeStr, nil
}

// UpdatedAt is the resolver for the updatedAt field.
func (r *metricParameterResolver) UpdatedAt(ctx context.Context, obj *models.MetricParameter) (*string, error) {
	timeStr := obj.UpdatedAt.Format(time.RFC3339)
	return &timeStr, nil
}

// Metric is the resolver for the metric field.
func (r *metricParameterResolver) Metric(ctx context.Context, obj *models.MetricParameter) (*models.Metric, error) {
	return r.MetricService.GetMetricByID(obj.MetricID)
}

// StageOrder is the resolver for the stageOrder field.
func (r *stageResolver) StageOrder(ctx context.Context, obj *models.Stage) (*int32, error) {
	order := int32(obj.StageOrder)
	return &order, nil
}

// OptimalTime is the resolver for the optimalTime field.
func (r *stageResolver) OptimalTime(ctx context.Context, obj *models.Stage) (*int32, error) {
	if obj.OptimalTime == 0 {
		return nil, nil
	}
	optimalTime := int32(obj.OptimalTime)
	return &optimalTime, nil
}

// CreatedAt is the resolver for the createdAt field.
func (r *stageResolver) CreatedAt(ctx context.Context, obj *models.Stage) (*string, error) {
	timeStr := obj.CreatedAt.Format(time.RFC3339)
	return &timeStr, nil
}

// UpdatedAt is the resolver for the updatedAt field.
func (r *stageResolver) UpdatedAt(ctx context.Context, obj *models.Stage) (*string, error) {
	timeStr := obj.UpdatedAt.Format(time.RFC3339)
	return &timeStr, nil
}

// Metrics is the resolver for the metrics field.
func (r *stageResolver) Metrics(ctx context.Context, obj *models.Stage) ([]*models.Metric, error) {
	metrics, err := r.StageService.GetStageMetrics(obj.StageID)
	if err != nil {
		return nil, fmt.Errorf("error fetching metrics: %w", err)
	}

	// Convert []models.Metric to []*models.Metric
	result := make([]*models.Metric, len(metrics))
	for i := range metrics {
		result[i] = &metrics[i]
	}

	return result, nil
}

// Game is the resolver for the game field.
func (r *stageResolver) Game(ctx context.Context, obj *models.Stage) (*models.Game, error) {
	return r.GameService.GetGameByID(obj.GameID)
}

// Competence returns graph.CompetenceResolver implementation.
func (r *Resolver) Competence() graph.CompetenceResolver { return &competenceResolver{r} }

// ConstantParameter returns graph.ConstantParameterResolver implementation.
func (r *Resolver) ConstantParameter() graph.ConstantParameterResolver {
	return &constantParameterResolver{r}
}

// Game returns graph.GameResolver implementation.
func (r *Resolver) Game() graph.GameResolver { return &gameResolver{r} }

// Metric returns graph.MetricResolver implementation.
func (r *Resolver) Metric() graph.MetricResolver { return &metricResolver{r} }

// MetricParameter returns graph.MetricParameterResolver implementation.
func (r *Resolver) MetricParameter() graph.MetricParameterResolver {
	return &metricParameterResolver{r}
}

// Stage returns graph.StageResolver implementation.
func (r *Resolver) Stage() graph.StageResolver { return &stageResolver{r} }

type competenceResolver struct{ *Resolver }
type constantParameterResolver struct{ *Resolver }
type gameResolver struct{ *Resolver }
type metricResolver struct{ *Resolver }
type metricParameterResolver struct{ *Resolver }
type stageResolver struct{ *Resolver }
