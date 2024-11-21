package db

import (
	"bonus-test/internal/domain/models"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type BonusCodeRewardRepository struct {
	db *gorm.DB
}

func NewBonusCodeRewardRepository(db *gorm.DB) *BonusCodeRewardRepository {
	return &BonusCodeRewardRepository{
		db: db,
	}
}

func (r *BonusCodeRewardRepository) GetByID(id uint64) (*models.BonusCodeReward, error) {
	if id == 0 {
		return nil, errors.New("boardRepository.bonus_code: bonus_code is nil")
	}

	var model models.BonusCodeReward

	if err := r.db.Model(models.BonusCodeReward{}).Where("id = ?", id).First(&model).Error; err != nil {
		return nil, errors.Wrap(err, "boardRepository.Create: error to create bonus_code")
	}

	return &model, nil
}

func (r *BonusCodeRewardRepository) Save(rew *models.BonusCodeReward) error {
	if rew == nil {
		return errors.New("bonus_code_reward is nil")
	}

	if err := r.db.Model(&models.BonusCodeReward{}).Save(rew).Error; err != nil {
		return errors.Wrap(err, "Create: error to create bonus_code")
	}

	return nil
}
