package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"testing"

	"github.com/alltestgreen/go-port-ingest/port-client-api/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const bodyString = `{
    "PortID1": {
        "name": "PortName1",
        "city": "City1",
        "country": "United Arab Emirates",
        "alias": [],
        "regions": [],
        "coordinates": [
            55.5136433,
            25.4052165
        ],
        "province": "Province1",
        "timezone": "Asia/Dubai",
        "unlocs": [
            "PortID1"
        ],
        "code": "10000"
    },
	"PortID2": {
        "name": "PortName2",
        "city": "City2",
        "country": "United Arab Emirates",
        "alias": [],
        "regions": [],
        "coordinates": [
            55.5136433,
            25.4052165
        ],
        "province": "Province2",
        "timezone": "Asia/Dubai",
        "unlocs": [
            "PortID2"
        ],
        "code": "20000"
    }
}`

func Test_parseJsonStream(t *testing.T) {
	tests := []struct {
		name    string
		data    []byte
		expErr  error
		expCall int
	}{
		{
			"good payload",
			[]byte(bodyString),
			nil,
			2,
		}, {
			"bad payload",
			[]byte("asd"),
			&json.SyntaxError{},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			serviceMock := new(portServiceMock)
			h := handler{
				portService: serviceMock,
			}

			r := ioutil.NopCloser(bytes.NewReader(tt.data))
			err := h.parseJsonStream(context.Background(), r)
			if tt.expErr != nil {
				assert.IsType(t, err, tt.expErr)
			}
		})
	}
}

type portServiceMock struct {
	mock.Mock
}

func (m *portServiceMock) ProcessPort(ctx context.Context, port model.Port) {
	log.Println(port)
}
