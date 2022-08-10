package main
 
import (
    "fmt"
    "github.com/CNwhisper/gin-api/global"
    "github.com/CNwhisper/gin-api/internal/model"
    "github.com/CNwhisper/gin-api/pkg/setting"
)
 
func main(){
    setting, _ := setting.NewSetting()
    _ = setting.ReadSection("Database", &global.DatabaseSetting)
 
    db, err := model.NewDBEngine(global.DatabaseSetting)
    err = db.AutoMigrate(&model.User{})
    if err != nil {
        fmt.Println(err)
    }
}
