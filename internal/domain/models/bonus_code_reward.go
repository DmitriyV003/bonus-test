package models

import (
	"database/sql"
	"time"
)

type BonusCodeReward struct {
	Id        uint64       `gorm:"id"`
	Name      string       `gorm:"name"`
	Type      string       `gorm:"type"`
	Reward    string       `gorm:"reward"`
	CreatedAt time.Time    `gorm:"created_at"`
	UpdatedAt time.Time    `gorm:"updated_at"`
	DeletedAt sql.NullTime `gorm:"deleted_at"`
}

func (*BonusCodeReward) TableName() string {
	return "bonus_code_rewards"
}
