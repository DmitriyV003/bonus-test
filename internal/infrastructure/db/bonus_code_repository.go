package db

import (
	"bonus-test/internal/domain/models"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type BonusCodeRepository struct {
	db *gorm.DB
}

func NewBonusCodeRepository(db *gorm.DB) *BonusCodeRepository {
	return &BonusCodeRepository{
		db: db,
	}
}

func (r *BonusCodeRepository) Save(board *models.BonusCode) error {
	if board == nil {
		return errors.New("boardRepository.bonus_code: bonus_code is nil")
	}

	if err := r.db.Model(board).Save(board).Error; err != nil {
		return errors.Wrap(err, "boardRepository.Create: error to create bonus_code")
	}

	return nil
}

func (s *BonusCodeRepository) FindBonusCodeByName(name string) (*models.BonusCode, error) {
	var bonusCode models.BonusCode
	if err := s.db.Where("name = ?", name).First(&bonusCode).Error; err != nil {
		return nil, err
	}
	return &bonusCode, nil
}
