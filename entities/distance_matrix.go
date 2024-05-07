package entities

type DistanceMatrix struct {
	DestinationAddress string
	OriginAddress      string
	Distance           string
	Duration           string
}

type GoogleMapsAPIInterface interface {
	GetDistanceMatrix(origins string, destinations string) (DistanceMatrix, error)
}
