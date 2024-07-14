package models

type Course struct {
	CourseId uint   `gorm:"primaryKey;autoIncrement;column:course_id"`
	Title    string `gorm:"column:title"`
	Img      string `gorm:"column:img"`
	Users    []User `gorm:"foreignKey:CourseID"`
}

func (Course) TableName() string {
	return "courses"
}
