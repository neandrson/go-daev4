package register

import "github.com/neandrson/go-daev4/orchestrator/entities"

type UserCreater interface {
	Create(user *entities.User) ([]byte, error)
}
