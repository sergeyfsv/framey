package speedtest

import (
	"context"
	"github.com/Masterminds/log-go/impl/logrus"
	"go.framey.io/speedtest/internal/core"
	"go.framey.io/speedtest/internal/fastdotcomclient"
	"go.framey.io/speedtest/internal/speedtestdotnetclient"
	"go.framey.io/speedtest/internal/types"
	"time"
)

func NewMeasurement(endpointType types.SpeedMeasurementEndpointType, cfg *types.Configuration) core.SpeedMeasurement {
	endpoints := make([]interface{}, 2)
	endpoints[types.FastdotcomEndpoint] = fastdotcomclient.FastdotcomEndpointImpl{}
	endpoints[types.SpeedtestdotnetEndpoint] = speedtestdotnetclient.SpeedtestdotnetEndpointImpl{}

	return core.SpeedMeasurement{
		Logger:        logrus.NewStandard(),
		EndpointType:  endpointType,
		Endpoint:      endpoints[endpointType].(core.SpeedTestInterface),
		Configuration: cfg,
	}
}

func Upload(measurement *core.SpeedMeasurement) (float64, error) {
	totalTimeout := time.Duration(measurement.Configuration.ConfigTimeout + measurement.Configuration.UploadTimeout)
	ctx, cancel := context.WithTimeout(context.Background(), totalTimeout)
	defer cancel()

	speed, err := measurement.Endpoint.Upload(ctx, measurement, func(s float64) string {
		return ""
	})
	return speed, err
}

func Download(measurement *core.SpeedMeasurement) (float64, error) {
	totalTimeout := time.Duration(measurement.Configuration.ConfigTimeout + measurement.Configuration.DownloadTimeout)
	ctx, cancel := context.WithTimeout(context.Background(), totalTimeout)
	defer cancel()

	speed, err := measurement.Endpoint.Download(ctx, measurement, func(s float64) string {
		return ""
	})
	return speed, err
}
