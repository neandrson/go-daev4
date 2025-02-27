package login

import "github.com/neandrson/go-daev4/orchestrator/entities"

type UserLogin interface {
	Login(user *entities.User) ([]byte, error)
}
