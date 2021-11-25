package types

import "time"

type Configuration struct {
	MeasurementRatio float64
	ConfigTimeout    time.Duration
	DownloadTimeout  time.Duration
	UploadTimeout    time.Duration
}

func DefaultConfiguration() Configuration {
	return Configuration{
		MeasurementRatio: 1,
		ConfigTimeout:    5 * time.Second,
		DownloadTimeout:  10 * time.Second,
		UploadTimeout:    10 * time.Second,
	}
}
