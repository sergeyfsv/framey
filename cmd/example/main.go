package main

import (
	"log"

	. "go.framey.io/speedtest"
	"go.framey.io/speedtest/internal/types"
)

func main() {
	measurementConfiguration := types.DefaultConfiguration()
	measurementConfiguration.MeasurementRatio = 8 / float64(types.Mbps)

	measurement := NewMeasurement(types.SpeedtestdotnetEndpoint, &measurementConfiguration)

	speed, err := Upload(&measurement)
	if err != nil {
		log.Fatalf("(%v) Upload speed measurement error. %v\n", measurement, err)
		return
	}
	log.Printf("Upload speed for %v: %v\n", measurement, speed)

	speed, err = Download(&measurement)
	if err != nil {
		log.Fatalf("(%v) Download speed measurement error. %v\n", measurement, err)
		return
	}
	log.Printf("(%v) Download speed: %v\n", measurement, speed)
}
