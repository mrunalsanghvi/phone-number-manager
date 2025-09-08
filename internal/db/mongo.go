package db

import (
	"context"
	"phone-number-manager/internal/models"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type PhoneBookRepository interface {
	CreateEntry(ctx context.Context, entry *models.PhoneBook) error
	GetEntry(ctx context.Context, uuid string) (*models.PhoneBook, error)
}

type MongoClient struct {
	NumbersCollection *mongo.Collection
	Client            *mongo.Client
}

func NewMongoClient(ctx context.Context, uri string) (*MongoClient, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	return &MongoClient{
		NumbersCollection: client.Database("phonebook").Collection("numbers"),
		Client:            client,
	}, nil
}

func (m *MongoClient) DisconnectMongoClient(ctx context.Context) error {
	return m.Client.Disconnect(ctx)
}

func (m *MongoClient) CreateEntry(ctx context.Context, entry *models.PhoneBook) error {
	_, err := m.NumbersCollection.InsertOne(ctx, entry)
	if err != nil {
		return err
	}
	return nil
}
func (m *MongoClient) GetEntry(ctx context.Context, uuid string) (*models.PhoneBook, error) {
	var entry models.PhoneBook
	err := m.NumbersCollection.FindOne(ctx, map[string]interface{}{"id": uuid}).Decode(&entry)
	if err != nil {
		return nil, err
	}
	return &entry, nil
}
