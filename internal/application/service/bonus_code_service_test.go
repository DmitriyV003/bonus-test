package service

import (
	"bonus-test/internal/domain/models"
	db2 "bonus-test/internal/infrastructure/db"
	"database/sql"
	"gorm.io/gorm"
	"testing"
	"time"
)

func TestBonusCodeService_checkBonusCodeRestrictions(t *testing.T) {
	type fields struct {
		bonusCodeRepository          *db2.BonusCodeRepository
		bonusCodeRenderingRepository *db2.BonusCodeRenderingRepository
		db                           *gorm.DB
	}
	type args struct {
		bonusCode *models.BonusCode
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "",
			fields: struct {
				bonusCodeRepository          *db2.BonusCodeRepository
				bonusCodeRenderingRepository *db2.BonusCodeRenderingRepository
				db                           *gorm.DB
			}{bonusCodeRepository: nil, bonusCodeRenderingRepository: nil, db: nil},
			args: struct{ bonusCode *models.BonusCode }{bonusCode: &models.BonusCode{
				Name:         "",
				MaxUsage:     1,
				CurrentUsage: 0,
				Status:       "active",
				ValidSince:   sql.NullTime{},
				ValidTill:    sql.NullTime{},
				CreatedAt:    time.Time{},
				UpdatedAt:    time.Time{},
				DeletedAt:    sql.NullTime{},
			}},
			wantErr: false,
		},
		{
			name: "",
			fields: struct {
				bonusCodeRepository          *db2.BonusCodeRepository
				bonusCodeRenderingRepository *db2.BonusCodeRenderingRepository
				db                           *gorm.DB
			}{bonusCodeRepository: nil, bonusCodeRenderingRepository: nil, db: nil},
			args: struct{ bonusCode *models.BonusCode }{bonusCode: &models.BonusCode{
				Name:         "",
				MaxUsage:     1,
				CurrentUsage: 0,
				Status:       "deleted",
				ValidSince:   sql.NullTime{},
				ValidTill:    sql.NullTime{},
				CreatedAt:    time.Time{},
				UpdatedAt:    time.Time{},
				DeletedAt:    sql.NullTime{},
			}},
			wantErr: true,
		},
		{
			name: "",
			fields: struct {
				bonusCodeRepository          *db2.BonusCodeRepository
				bonusCodeRenderingRepository *db2.BonusCodeRenderingRepository
				db                           *gorm.DB
			}{bonusCodeRepository: nil, bonusCodeRenderingRepository: nil, db: nil},
			args: struct{ bonusCode *models.BonusCode }{bonusCode: &models.BonusCode{
				Name:         "",
				MaxUsage:     0,
				CurrentUsage: 0,
				Status:       "active",
				ValidSince: sql.NullTime{
					Valid: true,
					Time:  time.Now().AddDate(0, 0, 10),
				},
				ValidTill: sql.NullTime{},
				CreatedAt: time.Time{},
				UpdatedAt: time.Time{},
				DeletedAt: sql.NullTime{},
			}},
			wantErr: true,
		},
		{
			name: "",
			fields: struct {
				bonusCodeRepository          *db2.BonusCodeRepository
				bonusCodeRenderingRepository *db2.BonusCodeRenderingRepository
				db                           *gorm.DB
			}{bonusCodeRepository: nil, bonusCodeRenderingRepository: nil, db: nil},
			args: struct{ bonusCode *models.BonusCode }{bonusCode: &models.BonusCode{
				Name:         "",
				MaxUsage:     0,
				CurrentUsage: 0,
				Status:       "active",
				ValidSince: sql.NullTime{
					Valid: true,
					Time:  time.Now().AddDate(0, 0, -10),
				},
				ValidTill: sql.NullTime{
					Valid: true,
					Time:  time.Now().AddDate(0, 0, 10),
				},
				CreatedAt: time.Time{},
				UpdatedAt: time.Time{},
				DeletedAt: sql.NullTime{},
			}},
			wantErr: false,
		},
		{
			name: "",
			fields: struct {
				bonusCodeRepository          *db2.BonusCodeRepository
				bonusCodeRenderingRepository *db2.BonusCodeRenderingRepository
				db                           *gorm.DB
			}{bonusCodeRepository: nil, bonusCodeRenderingRepository: nil, db: nil},
			args: struct{ bonusCode *models.BonusCode }{bonusCode: &models.BonusCode{
				Name:         "",
				MaxUsage:     0,
				CurrentUsage: 0,
				Status:       "deleted",
				ValidSince: sql.NullTime{
					Valid: true,
					Time:  time.Now().AddDate(0, 0, -10),
				},
				ValidTill: sql.NullTime{
					Valid: true,
					Time:  time.Now().AddDate(0, 0, 10),
				},
				CreatedAt: time.Time{},
				UpdatedAt: time.Time{},
				DeletedAt: sql.NullTime{},
			}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &BonusCodeService{
				bonusCodeRepository:          tt.fields.bonusCodeRepository,
				bonusCodeRenderingRepository: tt.fields.bonusCodeRenderingRepository,
				db:                           tt.fields.db,
			}
			if err := s.checkBonusCodeRestrictions(tt.args.bonusCode); (err != nil) != tt.wantErr {
				t.Errorf("checkBonusCodeRestrictions() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
