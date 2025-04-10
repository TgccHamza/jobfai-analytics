package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.68

import (
	"context"
	"fmt"
	"jobfai-analytics/internal/graph"
	"jobfai-analytics/internal/graph/model"
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
	panic(fmt.Errorf("not implemented: Metrics - metrics"))
}

// Game is the resolver for the game field.
func (r *competenceResolver) Game(ctx context.Context, obj *models.Competence) (*models.Game, error) {
	panic(fmt.Errorf("not implemented: Game - game"))
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
	panic(fmt.Errorf("not implemented: Game - game"))
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
	panic(fmt.Errorf("not implemented: Competencies - competencies"))
}

// Stages is the resolver for the stages field.
func (r *gameResolver) Stages(ctx context.Context, obj *models.Game) ([]*models.Stage, error) {
	panic(fmt.Errorf("not implemented: Stages - stages"))
}

// GameMetrics is the resolver for the gameMetrics field.
func (r *gameResolver) GameMetrics(ctx context.Context, obj *models.Game) ([]*model.GameMetric, error) {
	panic(fmt.Errorf("not implemented: GameMetrics - gameMetrics"))
}

// ConstantParameters is the resolver for the constantParameters field.
func (r *gameResolver) ConstantParameters(ctx context.Context, obj *models.Game) ([]*models.ConstantParameter, error) {
	panic(fmt.Errorf("not implemented: ConstantParameters - constantParameters"))
}

// CreatedAt is the resolver for the createdAt field.
func (r *metricResolver) CreatedAt(ctx context.Context, obj *models.Metric) (*string, error) {
	panic(fmt.Errorf("not implemented: CreatedAt - createdAt"))
}

// UpdatedAt is the resolver for the updatedAt field.
func (r *metricResolver) UpdatedAt(ctx context.Context, obj *models.Metric) (*string, error) {
	panic(fmt.Errorf("not implemented: UpdatedAt - updatedAt"))
}

// Parameters is the resolver for the parameters field.
func (r *metricResolver) Parameters(ctx context.Context, obj *models.Metric) ([]*models.MetricParameter, error) {
	panic(fmt.Errorf("not implemented: Parameters - parameters"))
}

// Competence is the resolver for the competence field.
func (r *metricResolver) Competence(ctx context.Context, obj *models.Metric) (*models.Competence, error) {
	panic(fmt.Errorf("not implemented: Competence - competence"))
}

// Stages is the resolver for the stages field.
func (r *metricResolver) Stages(ctx context.Context, obj *models.Metric) ([]*models.Stage, error) {
	panic(fmt.Errorf("not implemented: Stages - stages"))
}

// CreatedAt is the resolver for the createdAt field.
func (r *metricParameterResolver) CreatedAt(ctx context.Context, obj *models.MetricParameter) (*string, error) {
	panic(fmt.Errorf("not implemented: CreatedAt - createdAt"))
}

// UpdatedAt is the resolver for the updatedAt field.
func (r *metricParameterResolver) UpdatedAt(ctx context.Context, obj *models.MetricParameter) (*string, error) {
	panic(fmt.Errorf("not implemented: UpdatedAt - updatedAt"))
}

// Metric is the resolver for the metric field.
func (r *metricParameterResolver) Metric(ctx context.Context, obj *models.MetricParameter) (*models.Metric, error) {
	panic(fmt.Errorf("not implemented: Metric - metric"))
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
	panic(fmt.Errorf("not implemented: Metrics - metrics"))
}

// Game is the resolver for the game field.
func (r *stageResolver) Game(ctx context.Context, obj *models.Stage) (*models.Game, error) {
	panic(fmt.Errorf("not implemented: Game - game"))
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
