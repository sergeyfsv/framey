package types

import (
	"context"
	//	"time"
)

/*
func DefaultConfiguration() *Configuration {
	return &Configuration{
		ConfigTimeout:   1 * time.Second,
		LatencyTimeout:  1 * time.Second,
		DownloadTimeout: 10 * time.Second,
		UploadTimeout:   10 * time.Second,
	}
}
*/

type MeasurementCallback func(float64) string

type SpeedTestInterface interface {
	Upload(ctx context.Context, cfg *Configuration, callback MeasurementCallback) (float64, error)
	Download(ctx context.Context, cfg *Configuration, callback MeasurementCallback) (float64, error)
}
