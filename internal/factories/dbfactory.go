package factories

import (
	"context"
	"fmt"
	"phone-number-manager/internal/db"
)

func NewDBClient(ctx context.Context, uri string, dbType string) (db.PhoneBookRepository, error) {
	switch dbType {
	case "mongo":
		return db.NewMongoClient(ctx, uri)
	case "memory":
		return db.NewInMemoryPhoneBookRepository(ctx), nil
	default:
		return nil, fmt.Errorf("unsupported database type: %s", dbType)
	}
}
