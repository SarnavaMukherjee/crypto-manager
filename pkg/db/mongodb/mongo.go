/**
 * File: mongo.go
 * Author: Kedar Kekan
 * Contact: (support@airavana.ai)
 * Copyright (c) 2020 - 2021 Airavana Inc.
 */

package mongodb

import (
	"context"
	"fmt"
	"sync"
	"time"

	log "github.com/SarnavaMukherjee/crypto-manager/pkg/logger"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	connectTimeout = 5
	queryTimeout   = 30

	readPreference = "secondary"
	//connectionStringTemplate = "mongodb+srv://%s:%s@%s/admin?readPreference=%s&ssl=false"
	connectionStringTemplate = "mongodb://%s:%s@%s:27888"
)

type MongoService struct {
	Collection *mongo.Collection
}

var db *mongo.Client

// Singleton
var dial sync.Once

func (c *MongoService) BindCollection(database string, collection string) {
	c.Collection = db.Database(database).Collection(collection)
}

func new(connectionURI string) *mongo.Client {

	clientOptions := options.Client().ApplyURI(connectionURI)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal("", "", "Failed to create new mongodb client: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal("", "", "Failed to connect to cluster: %v", err)
	}

	// Force a connection to verify our connection string
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("", "", "Failed to ping cluster: %v", err)
	}

	return client
}

// NewMongoDB ...
func NewMongoDB(username, password, host string) {

	//connectionURI := fmt.Sprintf(connectionStringTemplate, username, password, host, readPreference)
	connectionURI := fmt.Sprintf(connectionStringTemplate, username, password, host)

	dial.Do(func() {
		db = new(connectionURI)
		log.Debug("", "", "MongoDB Connected. host: %s", host)
	})
}
