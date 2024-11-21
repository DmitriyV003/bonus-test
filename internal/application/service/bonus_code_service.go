package service

import (
	"bonus-test/internal/application/dto"
	"bonus-test/internal/domain/models"
	db2 "bonus-test/internal/infrastructure/db"
	"database/sql"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"time"
)

type BonusCodeService struct {
	bonusCodeRepository          *db2.BonusCodeRepository
	bonusCodeRenderingRepository *db2.BonusCodeRenderingRepository
	db                           *gorm.DB
}

func NewBonusCodeService(db *gorm.DB) *BonusCodeService {
	return &BonusCodeService{
		db:                           db,
		bonusCodeRepository:          db2.NewBonusCodeRepository(db),
		bonusCodeRenderingRepository: db2.NewBonusCodeRenderingRepository(db),
	}
}

func (s *BonusCodeService) Create(dto dto.CreateBonusCodedDTO, reward *models.BonusCodeReward) (*models.BonusCode, error) {
	bonusCode := models.BonusCode{
		Name:              dto.Name,
		MaxUsage:          dto.MaxUsage,
		CurrentUsage:      0,
		BonusCodeRewardID: reward.Id,
		Status:            dto.Status,
	}

	if dto.ValidSince != nil {
		bonusCode.ValidSince = sql.NullTime{
			Time:  *dto.ValidSince,
			Valid: true,
		}
	}

	if dto.ValidTill != nil {
		bonusCode.ValidTill = sql.NullTime{
			Time:  *dto.ValidTill,
			Valid: true,
		}
	}

	err := s.bonusCodeRepository.Save(&bonusCode)
	if err != nil {
		return nil, errors.Wrap(err, "bonusCodeService: error to create bonus_code")
	}

	return &bonusCode, nil
}

func (s *BonusCodeService) Update(dto dto.UpdateBonusCodedDTO, bonusCode *models.BonusCode) (*models.BonusCode, error) {
	bonusCode.Name = dto.Name
	bonusCode.MaxUsage = dto.MaxUsage
	bonusCode.Status = dto.Status

	if dto.ValidSince != nil {
		bonusCode.ValidSince = sql.NullTime{
			Time:  *dto.ValidSince,
			Valid: true,
		}
	}

	if dto.ValidTill != nil {
		bonusCode.ValidTill = sql.NullTime{
			Time:  *dto.ValidTill,
			Valid: true,
		}
	}

	err := s.bonusCodeRepository.Save(bonusCode)
	if err != nil {
		return nil, errors.Wrap(err, "bonusCodeService: error to Update bonus_code")
	}

	return bonusCode, nil
}

func (s *BonusCodeService) Delete(bonusCode *models.BonusCode) (*models.BonusCode, error) {
	bonusCode.DeletedAt = sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}
	bonusCode.Status = models.DeletedType

	err := s.bonusCodeRepository.Save(bonusCode)
	if err != nil {
		return nil, errors.Wrap(err, "bonusCodeService: error to Delete bonus_code")
	}

	return bonusCode, nil
}

func (s *BonusCodeService) Render(user *models.User, bonusCode *models.BonusCode) (*models.BonusCodeRendering, error) {
	err := s.checkBonusCodeRestrictions(bonusCode)
	if err != nil {
		return nil, errors.Wrap(err, "bonus code is not applicable")
	}

	rendering := models.BonusCodeRendering{
		UserID:      user.ID,
		BonusCodeID: bonusCode.Id,
	}

	tx := s.db.Begin()
	err = tx.Model(&rendering).Save(&rendering).Error
	if err != nil {
		tx.Rollback()
		return nil, errors.Wrap(err, "error to save bonus code rendering")
	}

	bonusCode.CurrentUsage += 1
	err = tx.Model(bonusCode).Save(bonusCode).Error
	if err != nil {
		tx.Rollback()
		return nil, errors.Wrap(err, "bonusCodeService: error to Delete bonus_code")
	}
	tx.Commit()

	return &rendering, nil
}

func (s *BonusCodeService) checkBonusCodeRestrictions(bonusCode *models.BonusCode) error {
	if !bonusCode.IsActive() {
		return errors.New("bonus code is not active")
	}

	if bonusCode.MaxUsage != 0 {
		if bonusCode.MaxUsage <= bonusCode.CurrentUsage {
			return errors.New("bonus code usage has spent")
		}
	}

	if bonusCode.ValidSince.Valid && bonusCode.ValidSince.Time.After(time.Now()) {
		return errors.New("bonus code ValidSince has spent")
	}

	if bonusCode.ValidTill.Valid && bonusCode.ValidTill.Time.Before(time.Now()) {
		return errors.New("bonus code ValidTill has spent")
	}

	return nil
}
