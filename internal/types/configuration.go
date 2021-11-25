package types

import "time"

const (
	KBps float64 = 1000
	MBps         = 1000 * KBps
	GBps         = 1000 * MBps

	Kbps float64 = 1000
	Mbps         = 1000 * Kbps
	Gbps         = 1000 * Mbps
)

type Configuration struct {
	MeasurementRatio float64
	ConfigTimeout    time.Duration
	LatencyTimeout   time.Duration
	DownloadTimeout  time.Duration
	UploadTimeout    time.Duration
}

func DefaultConfiguration() *Configuration {
	return &Configuration{
		MeasurementRatio: 1,
		ConfigTimeout:    1 * time.Second,
		LatencyTimeout:   1 * time.Second,
		DownloadTimeout:  10 * time.Second,
		UploadTimeout:    10 * time.Second,
	}
}
