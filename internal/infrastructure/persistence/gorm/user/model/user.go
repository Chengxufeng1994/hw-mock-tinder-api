package model

import (
	"gorm.io/gorm"

	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/infrastructure/persistence/gorm/shard"
)

type User struct {
	gorm.Model
	ID        string         `gorm:"primaryKey"`
	AccountID string         `gorm:"column:account_id;unique;not null"`
	Name      string         `gorm:"column:name"`
	Age       uint           `gorm:"column:age"`
	Gender    string         `gorm:"column:gender"`
	Photos    []Photo        `gorm:"many2many:users.user_photos;"`
	Interests []Interest     `gorm:"many2many:users.user_interests;"`
	Location  shard.GeoPoint `gorm:"type:geometry(Point,4326)"`
	Status    string         `gorm:"column:status;not null"`

	Preference Preference `gorm:"foreignKey:UserID"`
}

func (User) TableName() string {
	return "users.users"
}
