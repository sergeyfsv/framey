package core

import (
	"context"

	"github.com/Masterminds/log-go"
	"go.framey.io/speedtest/internal/types"
)

type SpeedMeasurement struct {
	Logger        log.Logger
	EndpointType  types.SpeedMeasurementEndpointType
	Endpoint      SpeedTestInterface
	Configuration *types.Configuration
}

type MeasurementCallback func(float64) string

type SpeedTestInterface interface {
	Upload(ctx context.Context, measurement *SpeedMeasurement, callback MeasurementCallback) (float64, error)
	Download(ctx context.Context, measurement *SpeedMeasurement, callback MeasurementCallback) (float64, error)
}

func (instance SpeedMeasurement) String() string {
	s := "speedtest.net"
	if instance.EndpointType == types.FastdotcomEndpoint {
		s = "fast.com"
	}
	return s
}
