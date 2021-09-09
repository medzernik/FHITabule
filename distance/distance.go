package distance

import (
	"FHITabule/config"
	"context"
	"fmt"
	"googlemaps.github.io/maps"
	"log"
)

func GetTimeToHOPA(from string, to string) float64 {
	if config.Cfg.Nastavenie.APIkluc == "" {
		return 0
	} else {

		c, err := maps.NewClient(maps.WithAPIKey(config.Cfg.Nastavenie.APIkluc))
		if err != nil {
			fmt.Println("ERROR: ", err)
			return 0
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
}
