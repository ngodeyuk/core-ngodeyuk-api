package models

type User struct {
	UserID   string `gorm:"primaryKey;column:user_id"`
	Name     string `gorm:"column:name"`
	ImgURL   string `gorm:"column:img_url"`
	Username string `gorm:"unique;column:username"`
	Password string `gorm:"column:password"`
	Heart    string `gorm:"column:heart"`
	Points   string `gorm:"column:points"`
}

func (User) TableName() string {
	return "users"
}
