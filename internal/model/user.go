package model
 
import "gorm.io/gorm"
 
type User struct {
    gorm.Model
    Name string `gorm:"not null" json:"name,omitempty"`
    Username string `gorm:"not null; unique" json:"username,omitempty"`
    Password string `gorm:"not null" json:"-"`
}

func (u User) Create(db *gorm.DB) (User, error) {
    err := db.Create(&u).Error
    return u, err
}