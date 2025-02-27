package calculateexpression

import "github.com/neandrson/go-daev4/orchestrator/entities"

type (
	Response struct {
		Expression entities.Expression `json:"expression"`
		Message    string              `json:"message"`
	}
)
