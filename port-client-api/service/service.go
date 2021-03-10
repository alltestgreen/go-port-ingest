package service

import (
	"context"
	"google.golang.org/grpc"
	"log"

	"github.com/alltestgreen/go-port-ingest/port-client-api/model"
	"github.com/alltestgreen/go-port-ingest/proto"
)

type PortService interface {
	ProcessPort(context.Context, model.Port)
}

type portService struct {
	portServiceClient proto.PortServiceClient
}

func NewPortService(rpcAddress string) (PortService, error) {
	conn, err := grpc.Dial(rpcAddress, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	client := proto.NewPortServiceClient(conn)

	return &portService{
		portServiceClient: client,
	}, nil
}

// ProcessPort receives a single Port, converts it to protobuf compatible and forwards to the gRPC service
func (s portService) ProcessPort(ctx context.Context, req model.Port) {
	log.Printf("Processing Port: (%s)", req.ID)

	port := req.ToProtoPayload()
	resp, err := s.portServiceClient.SubmitPort(ctx, &port)
	if err != nil {
		log.Fatalf("failed to send port (%s): %v", req.ID, err)
	}

	log.Printf("Port (%s) processed: (%d)", req.ID, resp.Code)
}
