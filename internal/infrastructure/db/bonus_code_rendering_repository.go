package db

import (
	"bonus-test/internal/domain/models"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type BonusCodeRenderingRepository struct {
	db *gorm.DB
}

func NewBonusCodeRenderingRepository(db *gorm.DB) *BonusCodeRenderingRepository {
	return &BonusCodeRenderingRepository{
		db: db,
	}
}

func (r *BonusCodeRenderingRepository) Save(bcr *models.BonusCodeRendering) error {
	if bcr == nil {
		return errors.New("BonusCodeRenderingRepository.bonus_code: bonus_code_rendering is nil")
	}

	if err := r.db.Model(bcr).Save(bcr).Error; err != nil {
		return errors.Wrap(err, "BonusCodeRenderingRepository.Create: error to create bonus_code_rendering")
	}

	return nil
}
