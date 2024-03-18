package main

import (
	"errors"
	"log"
	"net/http"

	"try-on/internal/middleware"
	"try-on/internal/pkg/app_errors"
	"try-on/internal/pkg/config"
	clothes "try-on/internal/pkg/delivery/clothes"
	session "try-on/internal/pkg/delivery/session"
	tryOn "try-on/internal/pkg/delivery/try_on"
	"try-on/internal/pkg/delivery/user_images"
	"try-on/internal/pkg/ml"
	clothesRepo "try-on/internal/pkg/repository/sqlc/clothes"
	clothesUsecase "try-on/internal/pkg/usecase/clothes"
	"try-on/internal/pkg/utils"

	"github.com/wagslane/go-rabbitmq"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
<<<<<<< Updated upstream
=======
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/wagslane/go-rabbitmq"
>>>>>>> Stashed changes
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type App struct {
	api    *fiber.App
	cfg    *config.Config
	logger *zap.SugaredLogger
}

func (app *App) Run() error {
	db, pg, err := app.getDB()
	if err != nil {
		return err
	}

	err = applyMigrations(app.cfg.Sql, db)
	if err != nil {
		return err
	}

	log.Println("Connecting to rabbit", app.cfg.Rabbit.DSN())
	rabbitConn, err := rabbitmq.NewConn(app.cfg.Rabbit.DSN())
	if err != nil {
		return err
	}

	clothesProcessor, err := ml.New(
		app.cfg.Rabbit.RequestQueue,
		app.cfg.Rabbit.ResponseQueue,
		rabbitConn,
	)
	if err != nil {
		return err
	}

	recover := recover.New(recover.Config{
		EnableStackTrace: true,
	})

	logger := logger.New(logger.Config{
		Format:     config.JsonLogFormat,
		TimeFormat: config.TimeFormat,
	})

	cors := cors.New(cors.Config{
		AllowOrigins:     app.cfg.Cors.Domain,
		AllowCredentials: app.cfg.Cors.AllowCredentials,
		MaxAge:           app.cfg.Cors.MaxAge,
	})

	sessionHandler := session.New(db, &app.cfg.Session)

	checkSession := middleware.CheckSession(middleware.SessionConfig{
		TokenName:    app.cfg.Session.TokenName,
		Sessions:     sessionHandler.Sessions,
		NoAuthRoutes: []string{"/register", "/login"},
		// SecureRoutes: []string{"/renew", "/clothes"},
	})

<<<<<<< Updated upstream
	clothesUsecase := clothesUsecase.New(clothesRepo.New(db))
=======
	clothesUsecase := clothesUsecase.New(clothesRepo.New(pg))
>>>>>>> Stashed changes

	clothesHandler := clothes.New(clothesUsecase, clothesProcessor, &app.cfg.Static)

	tryOnHandler := tryOn.New(db, clothesProcessor, clothesUsecase, app.logger, &app.cfg.Static)

	userImageHandler := user_images.New(db, &app.cfg.Static)

<<<<<<< Updated upstream
=======
	typeHandler := types.New(pg)

>>>>>>> Stashed changes
	app.api.Use(recover, logger, cors, middleware.AddLogger(app.logger), checkSession)

	app.api.Post("/register", sessionHandler.Register)
	app.api.Post("/login", sessionHandler.Login)
	app.api.Post("/renew", sessionHandler.Renew)

	app.api.Post("/clothes", clothesHandler.Upload)
	app.api.Get("/clothes", clothesHandler.GetOwn)
	app.api.Get("/clothes/:id", clothesHandler.GetByID)
	app.api.Delete("/clothes/:id", clothesHandler.Delete)
	app.api.Put("/clothes/:id", clothesHandler.Update)
	app.api.Get("/user/:id/clothes", clothesHandler.GetByUser)

	app.api.Get("/photos", userImageHandler.GetByUser)
	app.api.Get("/photos/:id", userImageHandler.GetByID)
	app.api.Post("/photos", userImageHandler.Upload)
	app.api.Delete("/photos/:id", userImageHandler.Delete)

	app.api.Post("/try-on", tryOnHandler.TryOn)
	app.api.Get("/try-on/:id", tryOnHandler.GetTryOnResult)

	app.api.Static("/static", app.cfg.Static.Dir)

	tryOnHandler.ListenTryOnResults()

	return app.api.Listen(app.cfg.Addr)
}

func NewApp(cfg *config.Config, logger *zap.SugaredLogger) *App {
	return &App{
		api: fiber.New(
			fiber.Config{
				ErrorHandler: errorHandler,
				JSONEncoder:  utils.EasyJsonMarshal,
				JSONDecoder:  utils.EasyJsonUnmarshal,
			},
		),
		cfg:    cfg,
		logger: logger,
	}
}

func (app *App) getDB() (*gorm.DB, *pgxpool.Pool, error) {
	pg, err := initPostgres(&app.cfg.Postgres)
	if err != nil {
		return nil, nil, err
	}

	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: stdlib.OpenDBFromPool(pg),
	}), &gorm.Config{
		// Logger: gormLogger.Discard,
		// FullSaveAssociations: true,
		TranslateError: true,
	})
	if err != nil {
		return nil, nil, err
	}

	return db, pg, nil
}

func errorHandler(ctx *fiber.Ctx, err error) error {
	var e *fiber.Error
	if errors.As(err, &e) {
		return ctx.Status(e.Code).JSON(
			&app_errors.ResponseError{
				Msg: e.Message,
			},
		)
	}

	msg := "Internal Server Error"

	var errorMsg *app_errors.ResponseError
	if errors.As(err, &errorMsg) {
		return ctx.Status(errorMsg.Code).JSON(errorMsg)
	}

	middleware.LogError(ctx, err)

	return ctx.Status(http.StatusInternalServerError).JSON(
		&app_errors.ResponseError{
			Msg: msg,
		},
	)
}
