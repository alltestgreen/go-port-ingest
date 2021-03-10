package repository

import "github.com/alltestgreen/go-port-ingest/proto"

type PortEntity struct {
	ID          string
	Name        string
	City        string
	Country     string
	Alias       []string
	Regions     []string
	Coordinates []float64
	Province    string
	Timezone    string
	Unlocs      []string
	Code        string
}

func PortEntityFromProto(p *proto.Port) PortEntity {
	return PortEntity{
		ID:          p.GetId(),
		Name:        p.GetName(),
		City:        p.GetCity(),
		Country:     p.GetCountry(),
		Alias:       p.GetAlias(),
		Regions:     p.GetRegions(),
		Coordinates: p.GetCoordinates(),
		Province:    p.GetProvince(),
		Timezone:    p.GetTimezone(),
		Unlocs:      p.GetUnlocs(),
		Code:        p.GetCode(),
	}
}
