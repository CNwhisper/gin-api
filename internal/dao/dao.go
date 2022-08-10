package dao
 
import "gorm.io/gorm"
 
type Dao struct {
    Engine *gorm.DB
}
 
func New(engine *gorm.DB) *Dao {
    return &Dao{Engine: engine}
}