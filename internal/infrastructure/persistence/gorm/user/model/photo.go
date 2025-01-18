package model

type Photo struct {
	ID  string `gorm:"primaryKey"`
	URL string `gorm:"column:url"`
}

func (Photo) TableName() string {
	return "users.photos"
}
