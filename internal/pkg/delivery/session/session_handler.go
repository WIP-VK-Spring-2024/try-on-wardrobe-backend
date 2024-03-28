package session

import (
	"try-on/internal/middleware"
	"try-on/internal/pkg/app_errors"
	"try-on/internal/pkg/config"
	"try-on/internal/pkg/domain"
	userRepo "try-on/internal/pkg/repository/sqlc/users"
	sessionUsecase "try-on/internal/pkg/usecase/session"
	userUsecase "try-on/internal/pkg/usecase/users"
	"try-on/internal/pkg/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mailru/easyjson"
)

type SessionHandler struct {
	Sessions domain.SessionUsecase
	users    domain.UserUsecase
	cfg      *config.Session
}

//easyjson:json
type tokenResponse struct {
	Token  string
	UserID utils.UUID
}

func New(db *pgxpool.Pool, cfg *config.Session) *SessionHandler {
	userRepo := userRepo.New(db)

	return &SessionHandler{
		Sessions: sessionUsecase.New(
			userRepo,
			cfg,
		),
		users: userUsecase.New(userRepo),
		cfg:   cfg,
	}
}

func (h *SessionHandler) Register(ctx *fiber.Ctx) error {
	var credentials domain.Credentials
	if err := easyjson.Unmarshal(ctx.Body(), &credentials); err != nil {
		middleware.LogWarning(ctx, err)
		return app_errors.ErrBadRequest
	}

	user, err := h.users.Create(credentials)
	if err != nil {
		return app_errors.New(err)
	}

	token, err := h.Sessions.IssueToken(user.ID)
	if err != nil {
		return app_errors.New(err)
	}

	return ctx.JSON(tokenResponse{
		Token:  token,
		UserID: user.ID,
	})
}

func (h *SessionHandler) Login(ctx *fiber.Ctx) error {
	var credentials domain.Credentials
	if err := easyjson.Unmarshal(ctx.Body(), &credentials); err != nil {
		middleware.LogWarning(ctx, err)
		return app_errors.ErrBadRequest
	}

	session, err := h.Sessions.Login(credentials)
	if err != nil {
		return app_errors.New(err)
	}
	return ctx.JSON(tokenResponse{
		Token:  session.ID,
		UserID: session.UserID,
	})
}

func (h *SessionHandler) Renew(ctx *fiber.Ctx) error {
	session := middleware.Session(ctx)
	if session == nil {
		return app_errors.ErrUnauthorized
	}

	token, err := h.Sessions.IssueToken(session.UserID)
	if err != nil {
		return app_errors.New(err)
	}

	return ctx.JSON(tokenResponse{
		Token:  token,
		UserID: session.UserID,
	})
}
