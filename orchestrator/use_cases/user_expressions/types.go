package userexpressions

import "github.com/neandrson/go-daev4/orchestrator/entities"

type (
	Response struct {
		UserExpressions []entities.Expression `json:"user_expressions"`
		Message         string                `json:"message"`
	}
)
