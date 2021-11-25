package types

type BytesPerSecond float64
type BitsPerSecond float64

const (
	KBps BytesPerSecond = 1000
	MBps                = 1000 * KBps
	GBps                = 1000 * MBps

	Kbps BitsPerSecond = 1000
	Mbps               = 1000 * Kbps
	Gbps               = 1000 * Mbps
)

type SpeedMeasurementEndpointType byte

const (
	FastdotcomEndpoint      SpeedMeasurementEndpointType = iota
	SpeedtestdotnetEndpoint                              = iota
)
