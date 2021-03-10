package server

import (
	"context"
	"log"
	"net/http"

	"github.com/alltestgreen/go-port-ingest/port-domain-service/repository"
	"github.com/alltestgreen/go-port-ingest/proto"
)

// a gRPC server using Protobuf serialisation
type server struct {
	proto.UnimplementedPortServiceServer
	repo repository.Repository
}

func NewServer() *server {
	repo := repository.NewRepository()
	return &server{
		repo: repo,
	}
}

// SubmitPort in the method implementation of PortServiceServer for receiving and storing ports using repository
func (s server) SubmitPort(ctx context.Context, port *proto.Port) (*proto.SubmitPortResponse, error) {
	log.Println("Processing port:", port.Id)

	entity := repository.PortEntityFromProto(port)
	if err := s.repo.PersistPort(ctx, entity); err != nil {
		log.Printf("failed to persist port (%s): %v", entity.ID, err)
		return &proto.SubmitPortResponse{Code: http.StatusInternalServerError, Details: "persist failed"}, err
	}

	return &proto.SubmitPortResponse{Code: http.StatusOK}, nil
}
