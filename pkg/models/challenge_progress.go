package models

type ChallengeProgress struct {
	ChallengeProgressId uint      `gorm:"primaryKey;autoIncrement;column:challenge_progress_id"`
	UserId              string    `gorm:"column:user_id"`
	ChallengeId         uint      `gorm:"column:challenge_id"`
	Completed           bool      `gorm:"column:completed"`
	User                User      `gorm:"foreignKey:UserId"`
	Challenge           Challenge `gorm:"foreignKey:ChallengeId"`
}

func (ChallengeProgress) TableName() string {
	return "challenge_progresses"
}
