package logger

import (
	"sync"

	"go.uber.org/zap"
)

var (
	Log  *zap.Logger
	once sync.Once
)

// Initialize structured logger (production or development config)
func init() {
	once.Do(func() {
		var err error
		Log, err = zap.NewProduction()
		if err != nil {
			panic(err)
		}
	})
}
