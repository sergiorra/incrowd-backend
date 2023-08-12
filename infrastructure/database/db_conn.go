package database

import (
	"context"
	"fmt"
	"incrowd-backend/config"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	IncrowdDB = "incrowd"
)

// SetupDatabaseConnection sets up a connection to a MongoDB using the provided configuration. It establishes the connection and sends a ping to ensure connectivity.
func SetupDatabaseConnection(dbConfig config.Database) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(dbConfig.SetupTimeoutInSeconds)*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s@%s:%d", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port)))
	if err != nil {
		return nil, err
	}

	if err = client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	return client, err
}
