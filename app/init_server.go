package app

import (
	"bonus-test/internal/application/service"
	db2 "bonus-test/internal/infrastructure/db"
	"bonus-test/internal/interface/http/bonus_code"
	"bonus-test/internal/interface/http/bonus_reward"
	"bonus-test/internal/interface/http/user"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog/log"
	"net/http"
)

type Server struct {
	e *echo.Echo
}

func NewServer() *Server {
	return &Server{e: echo.New()}
}

func (s *Server) InitServer() {
	s.connectMiddlewares()
	s.e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "ok")
	})
	s.connectApiRoutes()
	data, err := json.MarshalIndent(s.e.Routes(), "", "  ")
	if err != nil {
		log.Err(err).Msg("err")
	}
	log.Info().Bytes("routes", data).Msg("routes")
	log.Err(s.e.Start(":8080")).Msg("error to start server")
}

func (s *Server) connectApiRoutes() {
	gr := s.e.Group("api/v1/")

	bonusCodeService := service.NewBonusCodeService(GetDb())
	bonusRewardRepository := db2.NewBonusCodeRewardRepository(GetDb())
	bonusCodeRepository := db2.NewBonusCodeRepository(GetDb())
	userRepository := db2.NewUserRepository(GetDb())

	validate := validator.New(validator.WithRequiredStructEnabled())

	gr.POST("bonus-code", bonus_code.NewCreateHandler(bonusCodeService, bonusRewardRepository, validate).Handle)
	gr.POST("bonus-code-reward", bonus_reward.NewCreateHandler(bonusRewardRepository, validate).Handle)
	gr.POST("bonus-code/apply", user.NewApplyBonusCadeHandler(
		bonusCodeService,
		bonusCodeRepository,
		userRepository,
		bonusRewardRepository,
		validate,
	).Handle)
}

func (s *Server) connectMiddlewares() {
	s.e.Use(middleware.RemoveTrailingSlash())
	s.e.Use(middleware.RequestID())
	s.e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
		log.Info().
			Str("request_body", string(reqBody)).
			Str("response_body", string(resBody))
	}))
	s.e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:       true,
		LogStatus:    true,
		LogMethod:    true,
		LogError:     true,
		LogRequestID: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			if v.Error != nil {
				log.Err(v.Error).
					Str("Method", v.Method).
					Str("URI", v.URI).
					Int("status", v.Status).
					Time("start_time", v.StartTime).
					Str("request_id", v.RequestID).
					Msg("request")
			} else {
				log.Info().
					Str("Method", v.Method).
					Str("URI", v.URI).
					Int("status", v.Status).
					Time("start_time", v.StartTime).
					Str("request_id", v.RequestID).
					Msg("request")
			}

			return nil
		},
	}))
	s.e.Use(middleware.BodyDumpWithConfig(middleware.BodyDumpConfig{
		Handler: func(context echo.Context, reqBody []byte, resBody []byte) {
			log.Info().
				Str("request_body", string(reqBody)).
				Str("response_body", string(resBody))
		},
	}))
}
