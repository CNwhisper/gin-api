package service
 
import (
    "context"
    "github.com/CNwhisper/gin-api/global"
    "github.com/CNwhisper/gin-api/internal/dao"
)
 
type Service struct {
    Ctx context.Context
    Dao *dao.Dao
}
 
func New(ctx context.Context) Service{
    service := Service{Ctx: ctx}
    service.Dao = dao.New(global.DBEngine)
    return service
}