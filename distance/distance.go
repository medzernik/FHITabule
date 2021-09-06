package distance

import (
	"FHITabule/config"
	"context"
	"googlemaps.github.io/maps"
	"log"
)

func GetTimeToHOPA(from string, to string) float64 {

	c, err := maps.NewClient(maps.WithAPIKey(config.Cfg.Nastavenie.APIkluc))
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}

	r := &maps.DistanceMatrixRequest{
		Origins:      []string{from},
		Destinations: []string{to},
		Mode:         "transit",
	}
	route, err := c.DistanceMatrix(context.Background(), r)
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}

	minutesTravel := route.Rows[0].Elements[0].Duration.Minutes()

	return minutesTravel
}
