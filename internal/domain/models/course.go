package models

type Course struct {
	CourseId uint   `gorm:"primaryKey;autoIncrement;column:course_id"`
	Title    string `gorm:"column:title"`
	Img      string `gorm:"column:img"`
	Users    []User `gorm:"foreignKey:CourseId"`
	Units    []Unit `gorm:"foreignKey:CourseId"`
}

func (Course) TableName() string {
	return "courses"
}
