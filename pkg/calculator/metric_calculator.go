package calculator

import (
	"fmt"
	"reflect"
	"strconv"

	"jobfai-analytics/internal/models"
	"jobfai-analytics/pkg/evaluator"
)

// MetricCalculator handles the calculation of metrics using formulas
type MetricCalculator struct {
	formulaEvaluator *evaluator.FormulaEvaluator
}

// NewMetricCalculator creates a new metric calculator
func NewMetricCalculator(formulaEvaluator *evaluator.FormulaEvaluator) *MetricCalculator {
	return &MetricCalculator{
		formulaEvaluator: formulaEvaluator,
	}
}

// CalculateCompetenceMetric calculates a single competence metric based on parameters and constants
func (c *MetricCalculator) CalculateMetric(
	metric *models.Metric,
	parameters []models.MetricParameter,
	constants []models.ConstantParameter,
	playerData map[string]interface{},
) (map[string]interface{}, error) {
	// Prepare parameters for formula evaluation
	evalParams := make(map[string]interface{})

	// Add provided parameters
	for _, parameter := range parameters {
		if playerData[parameter.ParamKey] == nil {
			continue
		}
		evalParams[parameter.ParamKey] = c.convertParameter(playerData[parameter.ParamKey])
	}

	// Add constants
	for _, constant := range constants {
		evalParams[constant.ConstKey] = constant.ConstValue
	}

	// Validate required parameters
	for _, param := range parameters {
		if param.IsRequired {
			if _, exists := evalParams[param.ParamKey]; !exists {
				fmt.Printf("Using default value for parameter %s: %v\n", param.ParamKey, param.DefaultValue)
				if param.DefaultValue != "" {
					// Use default value if available
					evalParams[param.ParamKey] = c.convertParameter(param.DefaultValue)
				} else {
					evalParams[param.ParamKey] = 0
				}
			}
		}
	}

	// Evaluate the formula
	result, err := c.formulaEvaluator.Evaluate(metric.Formula, evalParams)
	if err != nil {
		return nil, fmt.Errorf("error calculating metric %s: %w", metric.MetricKey, err)
	}

	return map[string]interface{}{
		"value":   result,
		"rawData": evalParams,
	}, nil
}

// CalculateGameMetric calculates a global game metric based on stage performances and competence details
func (c *MetricCalculator) CalculateGameMetric(
	gameMetric *models.Metric,
	stagePerformance []map[string]interface{},
	competenceDetails map[string]map[string]interface{},
) (map[string]interface{}, error) {
	// Prepare parameters for formula evaluation
	evalParams := make(map[string]interface{})

	// Add stage performance data
	for i, stage := range stagePerformance {
		stageKey := fmt.Sprintf("stage_%d", i+1)
		evalParams[stageKey+"_score"] = c.convertParameter(stage["score"])
		evalParams[stageKey+"_time"] = c.convertParameter(stage["timeTaken"])
		evalParams[stageKey+"_optimal_time"] = c.convertParameter(stage["optimalTime"])

		// Add stage metrics
		metrics, ok := stage["metrics"].([]map[string]interface{})
		if ok {
			for _, metric := range metrics {
				metricKey, ok := metric["kpiId"].(string)
				if ok {
					evalParams[metricKey] = c.convertParameter(metric["value"])
				}
			}
		}
	}

	// Add competence details
	for competenceKey, competence := range competenceDetails {
		evalParams[competenceKey+"_score"] = c.convertParameter(competence["score"])
		evalParams[competenceKey+"_weight"] = c.convertParameter(competence["weight"])
	}

	// Evaluate the formula
	result, err := c.formulaEvaluator.Evaluate(gameMetric.Formula, evalParams)
	if err != nil {
		return nil, fmt.Errorf("error calculating global metric %s: %w", gameMetric.MetricKey, err)
	}

	return map[string]interface{}{
		"value":   result,
		"rawData": evalParams,
	}, nil
}

// convertParameter converts various parameter types to appropriate values for formula evaluation
func (c *MetricCalculator) convertParameter(value interface{}) interface{} {
	fmt.Println("Value: ", value)
	fmt.Println("Type:", reflect.TypeOf(value))
	switch v := value.(type) {
	case string:
		// Try to convert string to number if possible
		if floatVal, err := strconv.ParseFloat(v, 64); err == nil {
			return floatVal
		}

		if intVal, err := strconv.ParseInt(v, 10, 64); err == nil {
			return intVal
		}
		// If it's a boolean string
		if v == "true" {
			return 1
		}
		if v == "false" {
			return 0
		}
		return v
	case float64, float32, int, int64, int32, bool:
		return v
	case nil:
		return 0
	default:
		// For other types, convert to string
		return v
	}
}
