package user

import (
	"bonus-test/internal/application/service"
	"bonus-test/internal/domain/models"
	"bonus-test/internal/infrastructure/db"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type request struct {
	UserID    uint64 `json:"user_id" validate:"required"`
	BonusCode string `json:"bonus_code" validate:"required"`
}

type ApplyBonusCadeHandler struct {
	bonusCodeService          *service.BonusCodeService
	bonusCodeRepository       *db.BonusCodeRepository
	userRepository            *db.UserRepository
	bonusCodeRewardRepository *db.BonusCodeRewardRepository
	validator                 *validator.Validate
}

func NewApplyBonusCadeHandler(
	bonusCodeService *service.BonusCodeService,
	bonusCodeRepository *db.BonusCodeRepository,
	userRepository *db.UserRepository,
	bonusCodeRewardRepository *db.BonusCodeRewardRepository,
	validator *validator.Validate,
) *ApplyBonusCadeHandler {
	return &ApplyBonusCadeHandler{
		bonusCodeService:          bonusCodeService,
		bonusCodeRepository:       bonusCodeRepository,
		userRepository:            userRepository,
		bonusCodeRewardRepository: bonusCodeRewardRepository,
		validator:                 validator,
	}
}

func (h *ApplyBonusCadeHandler) Handle(c echo.Context) error {
	payload := new(request)
	if err := c.Bind(payload); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := h.validator.Struct(payload); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	user, err := h.userRepository.FindUserByID(payload.UserID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
	}

	bonusCode, err := h.bonusCodeRepository.FindBonusCodeByName(payload.BonusCode)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Bonus code not found"})
	}

	bonusCodeReward, err := h.bonusCodeRewardRepository.GetByID(bonusCode.BonusCodeRewardID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Bonus code not found"})
	}

	result, err := h.bonusCodeService.Render(user, bonusCode)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, mapResult(result, bonusCode, user, bonusCodeReward))
}

func mapResult(res *models.BonusCodeRendering, bonusCode *models.BonusCode, user *models.User, rew *models.BonusCodeReward) responseRendering {
	return responseRendering{
		Id: res.Id,
		User: responseUser{
			ID:        user.ID,
			Name:      user.Name,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
			DeletedAt: returnTime(user.DeletedAt),
		},
		BonusCode: responseBonusCode{
			Id:           bonusCode.Id,
			Name:         bonusCode.Name,
			MaxUsage:     bonusCode.MaxUsage,
			CurrentUsage: bonusCode.CurrentUsage,
			BonusCodeReward: responseReward{
				Id:        rew.Id,
				Name:      rew.Name,
				Type:      rew.Type,
				Reward:    rew.Reward,
				CreatedAt: rew.CreatedAt,
				UpdatedAt: rew.UpdatedAt,
				DeletedAt: returnTime(rew.DeletedAt),
			},
			Status:     string(bonusCode.Status),
			ValidSince: returnTime(bonusCode.ValidSince),
			ValidTill:  returnTime(bonusCode.ValidTill),
			CreatedAt:  bonusCode.CreatedAt,
			UpdatedAt:  bonusCode.UpdatedAt,
			DeletedAt:  returnTime(bonusCode.DeletedAt),
		},
		CreatedAt: res.CreatedAt,
		UpdatedAt: res.UpdatedAt,
	}
}

type responseRendering struct {
	Id        uint64            `json:"id"`
	User      responseUser      `json:"user"`
	BonusCode responseBonusCode `json:"bonus_code"`
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at"`
}

type responseBonusCode struct {
	Id              uint64         `json:"id"`
	Name            string         `json:"name"`
	MaxUsage        uint64         `json:"max_usage"`
	CurrentUsage    uint64         `json:"current_usage"`
	BonusCodeReward responseReward `json:"bonus_code_reward"`
	Status          string         `json:"status"`
	ValidSince      *time.Time     `json:"valid_since"`
	ValidTill       *time.Time     `json:"valid_till"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       *time.Time     `json:"deleted_at"`
}

type responseReward struct {
	Id        uint64     `json:"id"`
	Name      string     `json:"name"`
	Type      string     `json:"type"`
	Reward    string     `json:"reward"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type responseUser struct {
	ID        uint64     `json:"id"`
	Name      string     `json:"name"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func returnTime(nullTime sql.NullTime) *time.Time {
	if nullTime.Valid {
		return &nullTime.Time
	}

	return nil
}
