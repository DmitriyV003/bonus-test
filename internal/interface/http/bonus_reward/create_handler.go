package bonus_reward

import (
	"bonus-test/internal/domain/models"
	"bonus-test/internal/infrastructure/db"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
)

type request struct {
	Name   string `json:"name" validate:"required"`
	Type   string `json:"type" validate:"required"`
	Reward string `json:"reward" validate:"required"`
}

type CreateHandler struct {
	bonusRewardRepository *db.BonusCodeRewardRepository
	validator             *validator.Validate
}

func NewCreateHandler(bonusRewardRepository *db.BonusCodeRewardRepository, validator *validator.Validate) *CreateHandler {
	return &CreateHandler{bonusRewardRepository: bonusRewardRepository, validator: validator}
}

func (h *CreateHandler) Handle(c echo.Context) error {
	payload := new(request)
	if err := c.Bind(payload); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	if err := h.validator.Struct(payload); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	reward := models.BonusCodeReward{
		Name:   payload.Name,
		Type:   payload.Type,
		Reward: payload.Reward,
	}

	if err := h.bonusRewardRepository.Save(&reward); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create reward"})
	}

	return c.JSON(http.StatusCreated, reward)
}
