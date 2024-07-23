package models

type Lesson struct {
	LessonId   uint        `gorm:"primaryKey;autoIncrement;column:lesson_id"`
	UnitId     uint        `gorm:"column:unit_id"`
	Title      string      `gorm:"column:title"`
	Sequence   int         `gorm:"column:sequence"`
	Unit       Unit        `gorm:"foreignKey:UnitId"`
	Challenges []Challenge `gorm:"foreignKey:LessonId"`
}

func (Lesson) TableName() string {
	return "lessons"
}
