package models

type Unit struct {
	UnitId      uint     `gorm:"primaryKey;autoIncrement;column:unit_id"`
	CourseId    uint     `gorm:"column:course_id"`
	Title       string   `gorm:"column:title"`
	Description string   `gorm:"column:description"`
	Sequence    int      `gorm:"column:sequence"`
	Course      Course   `gorm:"foreignKey:CourseId"`
	Lessons     []Lesson `gorm:"foreignKey:UnitId"`
}

func (Unit) TableName() string {
	return "units"
}
