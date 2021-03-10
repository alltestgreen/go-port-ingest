package repository

import (
	"context"
)

type Repository interface {
	PersistPort(context.Context, PortEntity) error
	GetPortByID(ctx context.Context, id string) (*PortEntity, error)
}

type portRepository struct {
	database map[string]PortEntity
}

func NewRepository() Repository {
	return &portRepository{
		database: make(map[string]PortEntity),
	}
}

func (p portRepository) PersistPort(ctx context.Context, port PortEntity) error {

	p.database[port.ID] = port

	return nil
}

func (p portRepository) GetPortByID(ctx context.Context, id string) (*PortEntity, error) {

	port, ok := p.database[id]
	if ok {
		return &port, nil
	}

	return nil, nil
}
