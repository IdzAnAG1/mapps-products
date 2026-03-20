package app

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net"

	productsv1 "mapps_product/generated/mobileapps/proto/products/v1"
	"mapps_product/internal/config"
	"mapps_product/internal/db"
	"mapps_product/internal/interruptor"
	logger "mapps_product/internal/logger"
	"mapps_product/internal/server"

	"google.golang.org/grpc"
)

type App struct {
	tcpPort string
	logger  *slog.Logger
	db      *db.DB
}

func NewApp() (*App, error) {
	cfg, err := config.LoadAndGetConfig()
	if err != nil {
		return nil, err
	}
	log := logger.New(cfg.Logger.Level)
	log.Info("logger is created")

	database, err := db.New(context.Background(), cfg.Database.DSN(), log)
	if err != nil {
		return nil, fmt.Errorf("init database: %w", err)
	}
	log.Info("database connected")

	return &App{
		tcpPort: cfg.Server.Port,
		logger:  log,
		db:      database,
	}, nil
}

func (app *App) Run() error {
	app.logger.Info("Starting Product gRPC Server", "port", app.tcpPort)

	tcp, err := net.Listen(
		"tcp",
		fmt.Sprintf(":%s", app.tcpPort),
	)
	if err != nil {
		return err
	}
	defer func() {
		err = tcp.Close()
		if err == nil || errors.Is(err, net.ErrClosed) {
			app.logger.Info("tcp listener is closed")
			return
		}
		app.logger.Error("closing tcp listener is failed", "error", err)
	}()

	srv := grpc.NewServer()
	app.logger.Info("grpc server is created")

	iter := interruptor.NewInterruptor(srv, app.logger)
	iter.Run()

	productServer := &server.GrpcProductServer{
		Logger: app.logger,
		DB:     app.db,
	}
	productsv1.RegisterProductsServer(srv, productServer)
	app.logger.Info("product gRPC server registered")

	err = srv.Serve(tcp)
	if err != nil {
		return err
	}
	return nil
}
