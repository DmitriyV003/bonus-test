package bonus_code

import (
	"bonus-test/internal/application/dto"
	"bonus-test/internal/application/service"
	"bonus-test/internal/domain/models"
	"bonus-test/internal/infrastructure/db"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type request struct {
	Status     string     `json:"status" validate:"required"`
	MaxUsage   *uint64    `json:"max_usage"`
	Name       string     `json:"name" validate:"required"`
	ValidSince *time.Time `json:"valid_since"`
	ValidTill  *time.Time `json:"valid_till"`
	RewardID   uint64     `json:"reward_id" validate:"required"`
}

type CreateHandler struct {
	bonusCodeService      *service.BonusCodeService
	bonusRewardRepository *db.BonusCodeRewardRepository
	validator             *validator.Validate
}

func NewCreateHandler(bonusCodeService *service.BonusCodeService, bonusRewardRepository *db.BonusCodeRewardRepository, validator *validator.Validate) *CreateHandler {
	return &CreateHandler{bonusCodeService: bonusCodeService, bonusRewardRepository: bonusRewardRepository, validator: validator}
}

func (h *CreateHandler) Handle(c echo.Context) error {
	var (
		req request
	)

	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if err := h.validator.Struct(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	params := dto.CreateBonusCodedDTO{UpdateBonusCodedDTO: dto.UpdateBonusCodedDTO{
		Name:       req.Name,
		MaxUsage:   0,
		Status:     models.BonusCodeType(req.Status),
		ValidSince: req.ValidSince,
		ValidTill:  req.ValidTill,
	}}

	if req.MaxUsage != nil {
		params.MaxUsage = *req.MaxUsage
	}

	reward, err := h.bonusRewardRepository.GetByID(req.RewardID)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, "error to find bonus reward")
	}

	code, err := h.bonusCodeService.Create(params, reward)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, "error to create bonus_code")
	}

	return c.JSON(http.StatusCreated, code)
}
