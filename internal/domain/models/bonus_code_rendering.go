package models

import (
	"time"
)

type BonusCodeRendering struct {
	Id          uint64    `gorm:"id"`
	UserID      uint64    `gorm:"user_id"`
	BonusCodeID uint64    `gorm:"bonus_code_id"`
	CreatedAt   time.Time `gorm:"created_at"`
	UpdatedAt   time.Time `gorm:"updated_at"`
}

func (*BonusCodeRendering) TableName() string {
	return "bonus_code_renderings"
}
