package dao
 
import "github.com/CNwhisper/gin-api/internal/model"
 
func (dao *Dao) CreateUser(name, username, password string) (model.User, error) {
    user := model.User{
        Name: name,
        Username: username,
        Password: password,
    }
 
    return user.Create(dao.Engine)
}