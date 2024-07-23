package models

type Challenge struct {
	ChallengeId uint   `gorm:"primaryKey;autoIncrement;column:challenge_id"`
	LessonId    uint   `gorm:"column:lesson_id"`
	Question    string `gorm:"column:question"`
	Sequence    int    `gorm:"column:sequence"`
	Lesson      Lesson `gorm:"foreignKey:LessonId"`
}

func (Challenge) TableName() string {
	return "challenges"
}
