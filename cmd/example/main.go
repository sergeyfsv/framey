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
	log.Printf("Upload speed for %v: %v\n", measurement, speed)

	if err != nil {
		return
	}

	speed, err = Download(&measurement)
	log.Printf("Download speed for %v: %v\n", measurement, speed)

	if err != nil {
		return
	}
}
