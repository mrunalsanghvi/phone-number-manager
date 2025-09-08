package main

import (
	"context"
	"fmt"
	"net/http"
	"phone-number-manager/internal/api"
	"phone-number-manager/internal/factories"
	logger "phone-number-manager/internal/logging"
	"phone-number-manager/internal/service"
	"phone-number-manager/internal/tracing"

	"go.uber.org/zap"
)

func main() {
	// 1. Load config
	cfg, err := factories.NewConfig("config.yaml")
	if err != nil {
		panic(fmt.Errorf("config load failed: %w", err))
	}

	// 2. Init logging & tracing
	shutdownTracer := tracing.Init("phone-number-manager")
	defer shutdownTracer(context.Background())

	logger.Log.Info("Starting phone-number-manager", zap.String("mongo_uri", cfg.Mongo.URI), zap.Int("port", cfg.Server.Port))

	// 3. DB client
	repo, err := factories.NewDBClient(context.Background(), "", "memory")
	if err != nil {
		logger.Log.Fatal("DB client creation failed", zap.Error(err))
	}

	// 4. Services
	phoneBookService := service.NewPhoneBookService(repo)

	phoneBookHandler := api.NewHandler(phoneBookService)

	// 5. Router
	router := api.NewRouter(phoneBookHandler)
	router.Use(
		api.LoggingMiddleware,
		api.TracingMiddleware,
	)

	// 6. Start server
	addr := fmt.Sprintf(":%d", cfg.Server.Port)
	logger.Log.Info("Server listening", zap.String("addr", addr))
	if err := http.ListenAndServe(addr, router); err != nil {
		logger.Log.Fatal("Server failed", zap.Error(err))
	}
}
