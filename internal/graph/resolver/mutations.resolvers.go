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
	"jobfai-analytics/internal/services"
	"strconv"

	"github.com/google/uuid"
)

// CreateGame is the resolver for the createGame field.
// CreateGame is the resolver for the createGame field.
func (r *mutationResolver) CreateGame(ctx context.Context, input model.GameInput) (*models.Game, error) {
	// Generate a unique ID for the game using UUID
	gameID := uuid.New().String() // Truncate to 20 chars to match the DB column size

	game := &models.Game{
		GameID:      gameID,
		GameName:    input.GameName,
		Description: *input.Description,
		Active:      *input.Active,
	}

	err := r.GameService.CreateGame(game)
	if err != nil {
		return nil, fmt.Errorf("failed to create game: %w", err)
	}

	r.subscriptionManager.Publish("game:created", game)
	return game, nil
}

// UpdateGame is the resolver for the updateGame field.
func (r *mutationResolver) UpdateGame(ctx context.Context, input model.GameUpdateInput) (*models.Game, error) {
	// First fetch the existing game
	game, err := r.GameService.GetGameByID(input.GameID)
	if err != nil {
		return nil, fmt.Errorf("failed to find game: %w", err)
	}

	if game == nil {
		return nil, fmt.Errorf("game not found with ID: %s", input.GameID)
	}

	// Update fields
	if input.GameName != nil {
		game.GameName = *input.GameName
	}
	if input.Description != nil {
		game.Description = *input.Description
	}
	if input.Active != nil {
		game.Active = *input.Active
	}

	// Save the updated game
	err = r.GameService.UpdateGame(game)
	if err != nil {
		return nil, fmt.Errorf("failed to update game: %w", err)
	}

	// Publish update event
	r.subscriptionManager.Publish("game:updated", game)
	r.subscriptionManager.Publish(fmt.Sprintf("game:updated:%s", game.GameID), game)

	return game, nil
}

// DeleteGame is the resolver for the deleteGame field.
func (r *mutationResolver) DeleteGame(ctx context.Context, gameID string) (*bool, error) {
	err := r.GameService.DeleteGame(gameID)
	if err != nil {
		return nil, fmt.Errorf("failed to delete game: %w", err)
	}

	// Publish delete event
	r.subscriptionManager.Publish("game:deleted", &gameID)
	r.subscriptionManager.Publish(fmt.Sprintf("game:deleted:%s", gameID), &gameID)

	success := true
	return &success, nil
}

// CreateCompetence is the resolver for the createCompetence field.
func (r *mutationResolver) CreateCompetence(ctx context.Context, input model.CompetenceInput) (*models.Competence, error) {
	competence := &models.Competence{
		GameID:         input.GameID,
		CompetenceKey:  input.CompetenceKey,
		CompetenceName: input.CompetenceName,
		Benchmark:      *input.Benchmark,
		Description:    *input.Description,
		Weight:         *input.Weight,
	}

	err := r.CompetenceService.CreateCompetence(competence)
	if err != nil {
		return nil, fmt.Errorf("failed to create competence: %w", err)
	}

	// Publish create event
	r.subscriptionManager.Publish("competence:created", competence)
	r.subscriptionManager.Publish(fmt.Sprintf("competence:created:%s", competence.GameID), competence)

	return competence, nil
}

// UpdateCompetence is the resolver for the updateCompetence field.
func (r *mutationResolver) UpdateCompetence(ctx context.Context, input model.CompetenceUpdateInput) (*models.Competence, error) {
	competenceID, err := strconv.Atoi(input.CompetenceID)
	if err != nil {
		return nil, fmt.Errorf("invalid competence ID format: %w", err)
	}

	// Fetch the existing competence
	competence, err := r.CompetenceService.GetCompetenceByID(competenceID)
	if err != nil {
		return nil, fmt.Errorf("failed to find competence: %w", err)
	}

	if competence == nil {
		return nil, fmt.Errorf("competence not found with ID: %s", input.CompetenceID)
	}

	// Update fields
	if input.CompetenceKey != nil {
		competence.CompetenceKey = *input.CompetenceKey
	}
	if input.CompetenceName != nil {
		competence.CompetenceName = *input.CompetenceName
	}
	if input.Benchmark != nil {
		competence.Benchmark = *input.Benchmark
	}
	if input.Description != nil {
		competence.Description = *input.Description
	}
	if input.Weight != nil {
		competence.Weight = *input.Weight
	}

	// Save the updated competence
	err = r.CompetenceService.UpdateCompetence(competence)
	if err != nil {
		return nil, fmt.Errorf("failed to update competence: %w", err)
	}

	// Publish update event
	r.subscriptionManager.Publish("competence:updated", competence)
	r.subscriptionManager.Publish(fmt.Sprintf("competence:updated:%s", input.CompetenceID), competence)

	return competence, nil
}

// DeleteCompetence is the resolver for the deleteCompetence field.
func (r *mutationResolver) DeleteCompetence(ctx context.Context, competenceID string) (*bool, error) {
	id, err := strconv.Atoi(competenceID)
	if err != nil {
		return nil, fmt.Errorf("invalid competence ID format: %w", err)
	}

	err = r.CompetenceService.DeleteCompetence(id)
	if err != nil {
		return nil, fmt.Errorf("failed to delete competence: %w", err)
	}

	// Publish delete event
	r.subscriptionManager.Publish("competence:deleted", &competenceID)
	r.subscriptionManager.Publish(fmt.Sprintf("competence:deleted:%s", competenceID), &competenceID)

	success := true
	return &success, nil
}

// CreateStageMetric is the resolver for the createStageMetric field.
func (r *mutationResolver) CreateStageMetric(ctx context.Context, input model.StageMetricInput) (*models.Metric, error) {
	panic(fmt.Errorf("not implemented: CreateStageMetric - createStageMetric"))
}

// UpdateStageMetric is the resolver for the updateStageMetric field.
func (r *mutationResolver) UpdateStageMetric(ctx context.Context, input model.StageMetricUpdateInput) (*models.Metric, error) {
	panic(fmt.Errorf("not implemented: UpdateStageMetric - updateStageMetric"))
}

// DeleteStageMetric is the resolver for the deleteStageMetric field.
func (r *mutationResolver) DeleteStageMetric(ctx context.Context, metricID string) (*bool, error) {
	panic(fmt.Errorf("not implemented: DeleteStageMetric - deleteStageMetric"))
}

// CreateMetricParameter is the resolver for the createMetricParameter field.
func (r *mutationResolver) CreateMetricParameter(ctx context.Context, input model.MetricParameterInput) (*models.MetricParameter, error) {
	metricID, err := strconv.Atoi(input.MetricID)
	if err != nil {
		return nil, fmt.Errorf("invalid metric ID format: %w", err)
	}

	parameter := &models.MetricParameter{
		MetricID:    metricID,
		ParamKey:    input.ParamKey,
		ParamName:   input.ParamName,
		Description: *input.Description,
		ParamType:   input.ParamType,
		IsRequired:  *input.IsRequired,
	}

	err = r.MetricService.CreateCompetenceMetricParameter(parameter)
	if err != nil {
		return nil, fmt.Errorf("failed to create metric parameter: %w", err)
	}
	// Publish create event
	r.subscriptionManager.Publish("metric_parameter:created", parameter)
	r.subscriptionManager.Publish(fmt.Sprintf("metric_parameter:created:%s", input.MetricID), parameter)

	return parameter, nil
}

// UpdateMetricParameter is the resolver for the updateMetricParameter field.
func (r *mutationResolver) UpdateMetricParameter(ctx context.Context, input model.MetricParameterUpdateInput) (*models.MetricParameter, error) {
	paramID, err := strconv.Atoi(input.ParamID)
	if err != nil {
		return nil, fmt.Errorf("invalid parameter ID format: %w", err)
	}

	// Fetch the existing parameter
	parameter, err := r.MetricService.GetCompetenceMetricParameterByID(paramID)
	if err != nil {
		return nil, fmt.Errorf("failed to find metric parameter: %w", err)
	}

	if parameter == nil {
		return nil, fmt.Errorf("metric parameter not found with ID: %s", input.ParamID)
	}

	// Update fields
	if input.ParamKey != nil {
		parameter.ParamKey = *input.ParamKey
	}
	if input.ParamName != nil {
		parameter.ParamName = *input.ParamName
	}
	if input.Description != nil {
		parameter.Description = *input.Description
	}
	if input.ParamType != nil {
		parameter.ParamType = *input.ParamType
	}
	if input.IsRequired != nil {
		parameter.IsRequired = *input.IsRequired
	}

	// Save the updated parameter
	err = r.MetricService.UpdateMetricParameter(parameter)
	if err != nil {
		return nil, fmt.Errorf("failed to update metric parameter: %w", err)
	}

	// Publish update event
	r.subscriptionManager.Publish("metric_parameter:updated", parameter)
	r.subscriptionManager.Publish(fmt.Sprintf("metric_parameter:updated:%s", input.ParamID), parameter)

	return parameter, nil
}

// DeleteMetricParameter is the resolver for the deleteMetricParameter field.
func (r *mutationResolver) DeleteMetricParameter(ctx context.Context, paramID string) (*bool, error) {
	id, err := strconv.Atoi(paramID)
	if err != nil {
		return nil, fmt.Errorf("invalid parameter ID format: %w", err)
	}

	err = r.MetricService.DeleteMetricParameter(id)
	if err != nil {
		return nil, fmt.Errorf("failed to delete metric parameter: %w", err)
	}

	// Publish delete event
	r.subscriptionManager.Publish("metric_parameter:deleted", &paramID)
	r.subscriptionManager.Publish(fmt.Sprintf("metric_parameter:deleted:%s", paramID), &paramID)

	success := true
	return &success, nil
}

// CreateStage is the resolver for the createStage field.
func (r *mutationResolver) CreateStage(ctx context.Context, input model.StageInput) (*models.Stage, error) {
	stage := &models.Stage{
		GameID:      input.GameID,
		StageName:   input.StageName,
		StageKey:    input.StageKey,
		StageOrder:  int(input.StageOrder),
		Description: *input.Description,
	}

	err := r.StageService.CreateStage(stage)
	if err != nil {
		return nil, fmt.Errorf("failed to create stage: %w", err)
	}

	// Publish create event
	r.subscriptionManager.Publish("stage:created", stage)
	r.subscriptionManager.Publish(fmt.Sprintf("stage:created:%s", input.GameID), stage)

	return stage, nil
}

// UpdateStage is the resolver for the updateStage field.
func (r *mutationResolver) UpdateStage(ctx context.Context, input model.StageUpdateInput) (*models.Stage, error) {
	stageID, err := strconv.Atoi(input.StageID)
	if err != nil {
		return nil, fmt.Errorf("invalid stage ID format: %w", err)
	}

	// Fetch the existing stage
	stage, err := r.StageService.GetStageByID(stageID)
	if err != nil {
		return nil, fmt.Errorf("failed to find stage: %w", err)
	}

	if stage == nil {
		return nil, fmt.Errorf("stage not found with ID: %s", input.StageID)
	}

	// Update fields
	if input.StageName != nil {
		stage.StageName = *input.StageName
	}
	if input.StageKey != nil {
		stage.StageKey = *input.StageKey
	}
	if input.StageOrder != nil {
		stage.StageOrder = int(*input.StageOrder)
	}
	if input.Description != nil {
		stage.Description = *input.Description
	}

	// Save the updated stage
	err = r.StageService.UpdateStage(stage)
	if err != nil {
		return nil, fmt.Errorf("failed to update stage: %w", err)
	}

	// Publish update event
	r.subscriptionManager.Publish("stage:updated", stage)
	r.subscriptionManager.Publish(fmt.Sprintf("stage:updated:%s", input.StageID), stage)

	return stage, nil
}

// DeleteStage is the resolver for the deleteStage field.
func (r *mutationResolver) DeleteStage(ctx context.Context, stageID string) (*bool, error) {
	id, err := strconv.Atoi(stageID)
	if err != nil {
		return nil, fmt.Errorf("invalid stage ID format: %w", err)
	}

	err = r.StageService.DeleteStage(id)
	if err != nil {
		return nil, fmt.Errorf("failed to delete stage: %w", err)
	}

	// Publish delete event
	r.subscriptionManager.Publish("stage:deleted", &stageID)
	r.subscriptionManager.Publish(fmt.Sprintf("stage:deleted:%s", stageID), &stageID)

	success := true
	return &success, nil
}

// CreateGameMetric is the resolver for the createGameMetric field.
func (r *mutationResolver) CreateGameMetric(ctx context.Context, input model.GameMetricInput) (*model.GameMetric, error) {
	metric := &models.Metric{
		GameID:            input.GameID,
		MetricKey:         input.MetricKey,
		MetricName:        input.MetricName,
		Formula:           input.Formula,
		MetricDescription: *input.MetricDescription, // Fixed: using MetricDescription instead of Description
	}

	if input.Benchmark != nil {
		benchmark, err := strconv.ParseFloat(*input.Benchmark, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid benchmark value: %w", err)
		}
		metric.Benchmark = benchmark
	}

	err := r.MetricService.CreateGameMetric(metric)
	if err != nil {
		return nil, fmt.Errorf("failed to create game metric: %w", err)
	}

	// Convert to GraphQL model
	result := &model.GameMetric{
		MetricID:          strconv.Itoa(metric.MetricID),
		GameID:            &metric.GameID,
		MetricKey:         &metric.MetricKey,
		MetricName:        &metric.MetricName,
		MetricDescription: &metric.MetricDescription,
		Formula:           &metric.Formula,
	}

	if metric.Benchmark != 0 {
		benchmarkStr := strconv.FormatFloat(metric.Benchmark, 'f', -1, 64)
		result.Benchmark = &benchmarkStr
	}

	// Publish create event
	r.subscriptionManager.Publish("game_metric:created", metric)
	r.subscriptionManager.Publish(fmt.Sprintf("game_metric:created:%s", input.GameID), metric)

	return result, nil
}

// UpdateGameMetric is the resolver for the updateGameMetric field.
func (r *mutationResolver) UpdateGameMetric(ctx context.Context, input model.GameMetricUpdateInput) (*model.GameMetric, error) {
	metricID, err := strconv.Atoi(input.MetricID)
	if err != nil {
		return nil, fmt.Errorf("invalid metric ID format: %w", err)
	}

	// Fetch the existing metric
	metric, err := r.MetricService.GetGameMetricByID(metricID)
	if err != nil {
		return nil, fmt.Errorf("failed to find game metric: %w", err)
	}

	if metric == nil {
		return nil, fmt.Errorf("game metric not found with ID: %s", input.MetricID)
	}

	// Update fields
	if input.MetricKey != nil {
		metric.MetricKey = *input.MetricKey
	}
	if input.MetricName != nil {
		metric.MetricName = *input.MetricName
	}
	if input.Formula != nil {
		metric.Formula = *input.Formula
	}
	if input.Description != nil {
		metric.MetricDescription = *input.Description
	}
	if input.Benchmark != nil {
		benchmark, err := strconv.ParseFloat(*input.Benchmark, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid benchmark value: %w", err)
		}
		metric.Benchmark = benchmark
	}

	// Save the updated metric
	err = r.MetricService.UpdateGameMetric(metric)
	if err != nil {
		return nil, fmt.Errorf("failed to update game metric: %w", err)
	}

	// Convert to GraphQL model
	result := &model.GameMetric{
		MetricID:          input.MetricID,
		GameID:            &metric.GameID,
		MetricKey:         &metric.MetricKey,
		MetricName:        &metric.MetricName,
		MetricDescription: &metric.MetricDescription,
		Formula:           &metric.Formula,
	}

	if metric.Benchmark != 0 {
		benchmarkStr := strconv.FormatFloat(metric.Benchmark, 'f', -1, 64)
		result.Benchmark = &benchmarkStr
	}

	// Publish update event
	r.subscriptionManager.Publish("game_metric:updated", metric)
	r.subscriptionManager.Publish(fmt.Sprintf("game_metric:updated:%s", input.MetricID), metric)

	return result, nil
}

// DeleteGameMetric is the resolver for the deleteGameMetric field.
func (r *mutationResolver) DeleteGameMetric(ctx context.Context, metricID string) (*bool, error) {
	id, err := strconv.Atoi(metricID)
	if err != nil {
		return nil, fmt.Errorf("invalid metric ID format: %w", err)
	}

	err = r.MetricService.DeleteGameMetric(id)
	if err != nil {
		return nil, fmt.Errorf("failed to delete game metric: %w", err)
	}

	// Publish delete event
	r.subscriptionManager.Publish("game_metric:deleted", &metricID)
	r.subscriptionManager.Publish(fmt.Sprintf("game_metric:deleted:%s", metricID), &metricID)

	success := true
	return &success, nil
}

// CreateConstantParameter is the resolver for the createConstantParameter field.
func (r *mutationResolver) CreateConstantParameter(ctx context.Context, input model.ConstantParameterInput) (*models.ConstantParameter, error) {
	constant := &models.ConstantParameter{
		GameID:      input.GameID,
		ConstKey:    input.ConstKey,
		ConstName:   input.ConstName,
		ConstValue:  input.ConstValue,
		Description: *input.Description,
	}

	err := r.ConstantParameterService.CreateConstantParameter(constant)
	if err != nil {
		return nil, fmt.Errorf("failed to create constant parameter: %w", err)
	}

	// Publish create event
	r.subscriptionManager.Publish("constant_parameter:created", constant)
	r.subscriptionManager.Publish(fmt.Sprintf("constant_parameter:created:%s", input.GameID), constant)

	return constant, nil
}

// UpdateConstantParameter is the resolver for the updateConstantParameter field.
func (r *mutationResolver) UpdateConstantParameter(ctx context.Context, input model.ConstantParameterUpdateInput) (*models.ConstantParameter, error) {
	constID, err := strconv.Atoi(input.ConstID)
	if err != nil {
		return nil, fmt.Errorf("invalid constant ID format: %w", err)
	}

	// Fetch the existing constant
	constant, err := r.MetricService.GetConstantParameterByID(constID)
	if err != nil {
		return nil, fmt.Errorf("failed to find constant parameter: %w", err)
	}

	if constant == nil {
		return nil, fmt.Errorf("constant parameter not found with ID: %s", input.ConstID)
	}

	// Update fields
	if input.ConstKey != nil {
		constant.ConstKey = *input.ConstKey
	}
	if input.ConstName != nil {
		constant.ConstName = *input.ConstName
	}
	if input.ConstValue != nil {
		constant.ConstValue = *input.ConstValue
	}
	if input.Description != nil {
		constant.Description = *input.Description
	}

	// Save the updated constant
	err = r.ConstantParameterService.UpdateConstantParameter(constant)
	if err != nil {
		return nil, fmt.Errorf("failed to update constant parameter: %w", err)
	}

	// Publish update event
	r.subscriptionManager.Publish("constant_parameter:updated", constant)
	r.subscriptionManager.Publish(fmt.Sprintf("constant_parameter:updated:%s", input.ConstID), constant)

	return constant, nil
}

// DeleteConstantParameter is the resolver for the deleteConstantParameter field.
func (r *mutationResolver) DeleteConstantParameter(ctx context.Context, constID string) (*bool, error) {
	id, err := strconv.Atoi(constID)
	if err != nil {
		return nil, fmt.Errorf("invalid constant ID format: %w", err)
	}

	err = r.ConstantParameterService.DeleteConstantParameter(id)
	if err != nil {
		return nil, fmt.Errorf("failed to delete constant parameter: %w", err)
	}

	// Publish delete event
	r.subscriptionManager.Publish("constant_parameter:deleted", &constID)
	r.subscriptionManager.Publish(fmt.Sprintf("constant_parameter:deleted:%s", constID), &constID)

	success := true
	return &success, nil
}

// CalculatePlayerPerformance is the resolver for the calculatePlayerPerformance field.
func (r *mutationResolver) CalculatePlayerPerformance(ctx context.Context, input model.PlayerPerformanceInput) (*model.PlayerPerformance, error) {
	// Convert input parameters to the format expected by the service
	playerPerformanceInput := &services.PlayerPerformanceInput{
		PlayerID:    input.PlayerID,
		PlayerName:  input.PlayerName,
		ProfileType: *input.ProfileType,
		GameID:      input.GameID,
	}

	// Convert stage parameters
	for _, stage := range input.StageParameters {
		// Convert StageID from string to int
		stageID, err := strconv.Atoi(stage.StageID)
		if err != nil {
			return nil, fmt.Errorf("invalid stage ID format: %w", err)
		}

		stageParams := services.StageParametersInput{
			StageID:    stageID,
			Parameters: make(map[string]interface{}),
			TimeTaken:  stage.TimeTaken,
		}

		// Copy parameters
		for _, param := range stage.Parameters {
			stageParams.Parameters[param.ParamID] = param.Value
		}

		playerPerformanceInput.StageParameters = append(playerPerformanceInput.StageParameters, stageParams)
	}

	// Call the service to calculate performance
	result, err := r.PlayerPerformanceService.CalculatePlayerPerformance(*playerPerformanceInput)
	if err != nil {
		return nil, fmt.Errorf("failed to calculate player performance: %w", err)
	}

	// Convert the result to the GraphQL model
	performance := &model.PlayerPerformance{
		GameID:     input.GameID,
		TotalScore: &result.TotalScore,
	}

	// Add competence scores
	for key, comp := range result.CompetenceDetails {
		compData, ok := comp.(map[string]interface{})
		if !ok {
			continue
		}

		competenceScore := &model.CompetenceDetail{
			CompetenceKey: key,
		}

		if name, ok := compData["name"].(string); ok {
			competenceScore.Name = &name
		}
		if score, ok := compData["score"].(float64); ok {
			competenceScore.Score = &score
		}
		if benchmark, ok := compData["benchmark"].(float64); ok {
			competenceScore.Benchmark = &benchmark
		}
		if weight, ok := compData["weight"].(float64); ok {
			competenceScore.Weight = &weight
		}

		performance.CompetenceDetails = append(performance.CompetenceDetails, competenceScore)
	}

	return performance, nil
}

// Mutation returns graph.MutationResolver implementation.
func (r *Resolver) Mutation() graph.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
