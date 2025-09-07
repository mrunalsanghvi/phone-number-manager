package factories

import (
	"time"
	"phone-number-manager/db"
)

// NewDB initializes the MongoDB client.
func NewDB(uri string) (db.Client, error) {
	ctx := db.NewContextWithTimeout(10 * time.Second)
	return db.NewMongoClient(ctx, uri)
}