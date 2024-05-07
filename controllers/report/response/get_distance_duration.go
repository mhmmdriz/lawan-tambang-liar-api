package response

import "lawan-tambang-liar/entities"

type GetDistanceDuration struct {
	OriginAddress      string `json:"origin_address"`
	DestinationAddress string `json:"destination_address"`
	Distance           string `json:"distance"`
	Duration           string `json:"duration"`
}

func DistanceDurationFromEntitiesToResponse(distanceMatrix entities.DistanceMatrix) *GetDistanceDuration {
	return &GetDistanceDuration{
		OriginAddress:      distanceMatrix.OriginAddress,
		DestinationAddress: distanceMatrix.DestinationAddress,
		Distance:           distanceMatrix.Distance,
		Duration:           distanceMatrix.Duration,
	}
}
