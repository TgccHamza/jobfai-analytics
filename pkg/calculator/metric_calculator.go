package calculator

import (
	"fmt"
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
func (c *MetricCalculator) CalculateCompetenceMetric(
	metric *models.Metric,
	parameters map[string]interface{},
	constants []models.ConstantParameter,
) (map[string]interface{}, error) {
	// Prepare parameters for formula evaluation
	evalParams := make(map[string]interface{})

	// Add provided parameters
	for key, value := range parameters {
		evalParams[key] = c.convertParameter(value)
	}

	// Add constants
	for _, constant := range constants {
		evalParams[constant.ConstKey] = constant.ConstValue
	}

	// // Validate required parameters
	// for _, param := range metric.Parameters {
	// 	if param.IsRequired {
	// 		if _, exists := evalParams[param.ParamKey]; !exists {
	// 			if param.DefaultValue != "" {
	// 				// Use default value if available
	// 				evalParams[param.ParamKey] = c.convertParameter(param.DefaultValue)
	// 			} else {
	// 				return nil, fmt.Errorf("required parameter %s is missing", param.ParamKey)
	// 			}
	// 		}
	// 	}
	// }

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
	switch v := value.(type) {
	case string:
		// Try to convert string to number if possible
		if floatVal, err := strconv.ParseFloat(v, 64); err == nil {
			return floatVal
		}
		// If it's a boolean string
		if v == "true" {
			return true
		}
		if v == "false" {
			return false
		}
		return v
	case float64, float32, int, int64, int32, bool:
		return v
	default:
		// For other types, convert to string
		return fmt.Sprintf("%v", v)
	}
}
