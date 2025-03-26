package evaluator

import (
	"fmt"
	"math"

	"github.com/expr-lang/expr"
	"github.com/expr-lang/expr/vm"
)

// FormulaEvaluator handles dynamic formula evaluation using expr-lang/expr
type FormulaEvaluator struct {
	env map[string]interface{}
}

// NewFormulaEvaluator creates a new formula evaluator with math functions
func NewFormulaEvaluator() *FormulaEvaluator {
	// Create environment with common math functions
	env := map[string]interface{}{
		"max":   math.Max,
		"min":   math.Min,
		"abs":   math.Abs,
		"sqrt":  math.Sqrt,
		"pow":   math.Pow,
		"round": math.Round,
		"floor": math.Floor,
		"ceil":  math.Ceil,
	}

	return &FormulaEvaluator{
		env: env,
	}
}

// Evaluate evaluates a formula string with the given parameters
func (e *FormulaEvaluator) Evaluate(formula string, params map[string]interface{}) (float64, error) {
	// Create a combined environment with built-in functions and user parameters
	env := make(map[string]interface{})

	// Copy the base environment
	for k, v := range e.env {
		env[k] = v
	}

	// Add user parameters
	for k, v := range params {
		env[k] = v
	}

	// Compile the expression
	program, err := expr.Compile(formula, expr.Env(env))
	if err != nil {
		return 0, fmt.Errorf("failed to compile formula: %w", err)
	}

	// Run the program
	result, err := expr.Run(program, env)
	if err != nil {
		return 0, fmt.Errorf("failed to evaluate formula: %w", err)
	}

	// Convert result to float64
	return convertToFloat64(result)
}

// EvaluateWithCompilation evaluates a pre-compiled formula for better performance
func (e *FormulaEvaluator) EvaluateWithCompilation(program *vm.Program, params map[string]interface{}) (float64, error) {
	// Create a combined environment
	env := make(map[string]interface{})

	// Copy the base environment
	for k, v := range e.env {
		env[k] = v
	}

	// Add user parameters
	for k, v := range params {
		env[k] = v
	}

	// Run the program
	result, err := expr.Run(program, env)
	if err != nil {
		return 0, fmt.Errorf("failed to evaluate formula: %w", err)
	}

	// Convert result to float64
	return convertToFloat64(result)
}

// CompileFormula pre-compiles a formula for repeated use
func (e *FormulaEvaluator) CompileFormula(formula string) (*vm.Program, error) {
	return expr.Compile(formula, expr.Env(e.env))
}

// AddFunction adds a custom function to the evaluator
func (e *FormulaEvaluator) AddFunction(name string, fn interface{}) {
	e.env[name] = fn
}

// convertToFloat64 converts various types to float64
func convertToFloat64(value interface{}) (float64, error) {
	switch v := value.(type) {
	case float64:
		return v, nil
	case float32:
		return float64(v), nil
	case int:
		return float64(v), nil
	case int64:
		return float64(v), nil
	case int32:
		return float64(v), nil
	case uint:
		return float64(v), nil
	case uint64:
		return float64(v), nil
	case bool:
		if v {
			return 1.0, nil
		}
		return 0.0, nil
	default:
		return 0, fmt.Errorf("cannot convert %T to float64", value)
	}
}
