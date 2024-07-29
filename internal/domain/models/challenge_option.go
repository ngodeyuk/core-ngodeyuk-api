package models

type ChallengeOption struct {
	ChallengeOptionId uint      `gorm:"primaryKey;autoIncrement;column:challenge_option_id"`
	ChallengeId       uint      `gorm:"column:challenge_id"`
	Text              string    `gorm:"column:text"`
	Correct           bool      `gorm:"column:correct"`
	ImgUrl            string    `gorm:"column:img_url"`
	Challenge         Challenge `gorm:"foreignKey:ChallengeId"`
}

func (ChallengeOption) TableName() string {
	return "challenge_options"
}
