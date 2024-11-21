package models

import (
	"database/sql"
	"time"
)

type BonusCodeType string

const ActiveType BonusCodeType = "active"
const DeletedType BonusCodeType = "deleted"

type BonusCode struct {
	Id                uint64        `gorm:"id"`
	Name              string        `gorm:"name"`
	MaxUsage          uint64        `gorm:"max_usage"`
	CurrentUsage      uint64        `gorm:"current_usage"`
	BonusCodeRewardID uint64        `gorm:"bonus_code_reward_id"`
	Status            BonusCodeType `gorm:"status"`
	ValidSince        sql.NullTime  `gorm:"valid_since"`
	ValidTill         sql.NullTime  `gorm:"valid_till"`
	CreatedAt         time.Time     `gorm:"created_at"`
	UpdatedAt         time.Time     `gorm:"updated_at"`
	DeletedAt         sql.NullTime  `gorm:"deleted_at"`
}

func (*BonusCode) TableName() string {
	return "bonus_codes"
}

func (b *BonusCode) IsActive() bool {
	return b.Status == ActiveType
}
