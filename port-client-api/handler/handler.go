package handler

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/alltestgreen/go-port-ingest/port-client-api/model"
	"github.com/alltestgreen/go-port-ingest/port-client-api/service"
	"github.com/gin-gonic/gin"
)

type handler struct {
	gin         *gin.Engine
	portService service.PortService
}

func ServeEndpoints(portService service.PortService) error {
	g := gin.Default()
	h := &handler{
		gin:         g,
		portService: portService,
	}
	return h.registerRoutes()
}

func (h handler) registerRoutes() error {
	h.gin.GET("/_health", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	h.gin.POST("/ports", h.handlePortsSubmission)

	return h.gin.Run(":8080")
}

func (h handler) handlePortsSubmission(c *gin.Context) {
	req := c.Request
	defer req.Body.Close()

	if err := h.parseJsonStream(c, req.Body); err != nil {
		log.Fatalf("error processing json stream: %v", err)
	}
}

// parseJsonStream tries to parse the json payload and sends off successfully parsed items to the service asynchronously
func (h handler) parseJsonStream(ctx context.Context, body io.ReadCloser) error {
	dec := json.NewDecoder(body)

	// read first delimiter '{'
	if _, err := dec.Token(); err != nil {
		return err
	}

	// parse each item individually, send to service for processing
	for dec.More() {
		portID, err := dec.Token()
		if err != nil {
			log.Printf("error parsing item ID: %v", err)
			continue
		}

		var port model.Port
		if err = dec.Decode(&port); err != nil {
			log.Printf("error parsing item: %v", err)
			continue
		}

		var ok bool
		port.ID, ok = portID.(string)
		if !ok {
			log.Printf("error parsing item ID: %v", err)
			continue
		}

		go h.portService.ProcessPort(ctx, port)
	}

	return nil
}
