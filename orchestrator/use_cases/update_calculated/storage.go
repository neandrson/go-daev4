package updatecalculated

import "github.com/neandrson/go-daev4/orchestrator/entities"

type ExpressionUpdater interface {
	UpdateCalculated(expression *entities.Expression) error
}
