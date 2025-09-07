package main

import (
	"context"
	"fmt"
	"net/http"
	"phone-number-manager/factories"
	"phone-number-manager/logging"
	"phone-number-manager/tracing"
	"phone-number-manager/api"
	"go.uber.org/zap"
)

func main() {
	// 1. Load config
	cfg, err := factories.NewConfig("config.yaml")
	if err != nil {
		panic(fmt.Errorf("config load failed: %w", err))
	}

	// 2. Init logging & tracing
	logging.Init(cfg.Logging.Level)
	shutdownTracer := tracing.Init("phone-number-manager")
	defer shutdownTracer(context.Background())
	log := logging.L()
	log.Info("Starting phone-number-manager", zap.String("mongo_uri", cfg.Mongo.URI), zap.Int("port", cfg.Server.Port))

	// 3. DB client
	dbClient, err := factories.NewDB(cfg.Mongo.URI)
	if err != nil {
		log.Fatal("MongoDB connection failed", zap.Error(err))
	}
	defer dbClient.Disconnect(nil)

	// 4. HTTP client for outbound calls
	httpClient := factories.NewHTTPClient()

	// 5. Router + middleware
	router := api.NewRouter()
	routerWithMiddleware := api.LoggingMiddleware(api.TracingMiddleware(router))

	// 6. Start server
	addr := fmt.Sprintf(":%d", cfg.Server.Port)
	log.Info("Server listening", zap.String("addr", addr))
	if err := http.ListenAndServe(addr, routerWithMiddleware); err != nil {
		log.Fatal("Server failed", zap.Error(err))
	}
}