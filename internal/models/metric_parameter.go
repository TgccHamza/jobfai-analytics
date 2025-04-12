package models

import (
	"time"
)

type ParameterType string

const (
	ParameterTypeInteger  ParameterType = "INTEGER"
	ParameterTypeDecimal  ParameterType = "DECIMAL"
	ParameterTypeBoolean  ParameterType = "BOOLEAN"
	ParameterTypeString   ParameterType = "STRING"
	ParameterTypeConstant ParameterType = "CONSTANT"
	ParameterTypeGlobal   ParameterType = "GLOBAL"
)

type MetricParameter struct {
	ParamID          int           `gorm:"column:param_id;primaryKey;autoIncrement" json:"paramId"`
	MetricID         int           `gorm:"column:metric_id" json:"metricId"`
	ParamKey         string        `gorm:"column:param_key;type:varchar(255)" json:"paramKey"`
	ParamName        string        `gorm:"column:param_name;type:varchar(255)" json:"paramName"`
	ParamDescription string        `gorm:"column:param_description;type:text" json:"paramDescription"`
	ParamType        ParameterType `gorm:"column:param_type;type:enum('INTEGER','DECIMAL','BOOLEAN','STRING','CONSTANT','GLOBAL')" json:"paramType"`
	IsRequired       bool          `gorm:"column:is_required;default:true" json:"isRequired"`
	DefaultValue     string        `gorm:"column:default_value;type:varchar(255)" json:"defaultValue"`
	Description      string        `gorm:"column:description;type:text" json:"description"`
	CreatedAt        time.Time     `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
	UpdatedAt        time.Time     `gorm:"column:updated_at;autoUpdateTime" json:"updatedAt"`
}

func (MetricParameter) TableName() string {
	return "metric_parameters"
}
