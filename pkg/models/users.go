package models

type User struct {
	UserID   uint   `gorm:"primaryKey;autoIncrement;column:user_id"`
	CourseID uint   `gorm:"column:course_id"`
	Name     string `gorm:"column:name"`
	ImgURL   string `gorm:"column:img_url"`
	Username string `gorm:"unique;column:username"`
	Password string `gorm:"column:password"`
	Heart    string `gorm:"column:heart"`
	Points   string `gorm:"column:points"`
	Course   Course `gorm:"foreignKey:CourseId"`
}

func (User) TableName() string {
	return "users"
}
