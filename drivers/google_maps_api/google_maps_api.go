package google_maps_api

import (
	"context"
	"lawan-tambang-liar/entities"

	"googlemaps.github.io/maps"
)

type GoogleMapsAPI struct {
	APIKey string
}

func NewGoogleMapsAPI(apiKey string) *GoogleMapsAPI {
	return &GoogleMapsAPI{
		APIKey: apiKey,
	}
}

func (g *GoogleMapsAPI) GetDistanceMatrix(origins string, destinations string) (entities.DistanceMatrix, error) {
	c, err := maps.NewClient(maps.WithAPIKey(g.APIKey))
	if err != nil {
		return entities.DistanceMatrix{}, err
	}
	r := &maps.DistanceMatrixRequest{
		Origins:      []string{origins},
		Destinations: []string{destinations},
		Units:        maps.UnitsMetric,
	}

	matrix, err := c.DistanceMatrix(context.Background(), r)
	if err != nil {
		return entities.DistanceMatrix{}, err
	}

	distanceMatrix := entities.DistanceMatrix{
		DestinationAddress: matrix.DestinationAddresses[0],
		OriginAddress:      matrix.OriginAddresses[0],
		Distance:           matrix.Rows[0].Elements[0].Distance.HumanReadable,
		Duration:           matrix.Rows[0].Elements[0].Duration.String(),
	}

	return distanceMatrix, nil
}
