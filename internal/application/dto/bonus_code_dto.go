package dto

import (
	"bonus-test/internal/domain/models"
	"time"
)

type CreateBonusCodedDTO struct {
	UpdateBonusCodedDTO
}

type UpdateBonusCodedDTO struct {
	Name       string
	MaxUsage   uint64
	Status     models.BonusCodeType
	ValidSince *time.Time
	ValidTill  *time.Time
}
