package createexpression

import "github.com/neandrson/go-daev4/orchestrator/entities"

type ExpressionCreater interface {
	CreateNew(expression *entities.Expression) error
}
