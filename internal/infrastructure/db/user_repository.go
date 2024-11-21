package db

import (
	"bonus-test/internal/domain/models"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Save(board *models.BonusCode) error {
	if board == nil {
		return errors.New("boardRepository.bonus_code: bonus_code is nil")
	}

	if err := r.db.Model(board).Save(board).Error; err != nil {
		return errors.Wrap(err, "boardRepository.Create: error to create bonus_code")
	}

	return nil
}

func (r *UserRepository) FindUserByID(userID uint64) (*models.User, error) {
	var user models.User
	if err := r.db.Model(&user).Where("id = ?", userID).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
