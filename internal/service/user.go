package service
 
import (
    "github.com/CNwhisper/gin-api/internal/model"
    "github.com/CNwhisper/gin-api/internal/requests"
)
 
func (s Service) CreateUser(params *requests.RegisterUserRequest) (model.User, error) {
    return s.Dao.CreateUser(params.Name, params.Username, params.Password)
}