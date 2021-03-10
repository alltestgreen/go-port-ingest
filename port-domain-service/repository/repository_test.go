package repository

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	portV1 = PortEntity{ID: "ID", Name: "Port V1"}
	portV2 = PortEntity{ID: "ID", Name: "Port V2"}
)

func Test_portRepository_PersistPort(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name  string
		ports []PortEntity
	}{
		{
			"store single port",
			[]PortEntity{portV1},
		}, {
			"store and update port",
			[]PortEntity{portV1, portV2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := NewRepository()
			for _, port := range tt.ports {
				// save port
				assert.NoError(t, repo.PersistPort(ctx, port))

				// check database content
				portByID, err := repo.GetPortByID(ctx, port.ID)
				assert.NoError(t, err)
				assert.NotNil(t, t, portByID)
				assert.Equal(t, *portByID, port)
			}
		})
	}
}
