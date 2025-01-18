package model

type Interest struct {
	ID   int    `gorm:"primaryKey"`
	Name string `gorm:"column:name"`
}

func (Interest) TableName() string {
	return "users.interests"
}
