package service

import (
	"context"
	"log"

	"github.com/alltestgreen/go-port-ingest/port-client-api/model"
)

type PortService interface {
	ProcessPort(context.Context, model.Port)
}

type portService struct {
}

func NewPortService() PortService {
	return &portService{}
}

func (s portService) ProcessPort(ctx context.Context, req model.Port) {
	log.Printf("Processing Port: (%s)", req.ID)
}
