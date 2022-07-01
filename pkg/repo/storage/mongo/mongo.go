package mongo

import (
	"context"

	"relay_server/pkg/logger"
	repo "relay_server/pkg/repo/storage"

	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDB struct {
	Collection *mongo.Collection
}

var mongoTest repo.IInfoHost = (*MongoDB)(nil)

func (m *MongoDB) GetInfoHost(ctx context.Context, log logger.ILogger, id string) (*repo.Host, error) { 
	
	return &repo.Host{}, nil
}

func (m *MongoDB) SaveInfoHost(ctx context.Context, h *repo.Host) error {
	return nil
}