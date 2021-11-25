package fastdotcomclient

import (
	"context"

	"go.framey.io/speedtest/internal/core"
	"go.framey.io/speedtest/internal/measurementstream"
	"go.jonnrb.io/speedtest/fastdotcom"
	"go.jonnrb.io/speedtest/units"
)

type FastdotcomEndpointImpl struct{}

func (instance FastdotcomEndpointImpl) Download(
	ctx context.Context,
	measurement *core.SpeedMeasurement,
	measurementCallback core.MeasurementCallback) (float64, error) {

	var client fastdotcom.Client

	initCtx, cancel := context.WithTimeout(context.Background(), measurement.Configuration.ConfigTimeout)
	defer cancel()

	manifest, err := fastdotcom.GetManifest(initCtx, 5)
	if err != nil {
		measurement.Logger.Errorf("Download. Error loading configuration: %v", err)
		return 0, err
	}

	stream, finalize := measurementstream.CreateMeasurementStream(measurement.Configuration.MeasurementRatio, measurementCallback)
	speed, err := manifest.ProbeDownloadSpeed(ctx, &client, stream)
	if err != nil {
		measurement.Logger.Errorf("Download. Error probing speed: %v", err)
		return 0, err
	}

	result := float64(speed) * measurement.Configuration.MeasurementRatio

	finalize(units.BytesPerSecond(result))
	return result, nil
}

func (instance FastdotcomEndpointImpl) Upload(
	ctx context.Context,
	measurement *core.SpeedMeasurement,
	measurementCallback core.MeasurementCallback) (float64, error) {

	var client fastdotcom.Client

	initCtx, cancel := context.WithTimeout(context.Background(), measurement.Configuration.ConfigTimeout)
	defer cancel()

	manifest, err := fastdotcom.GetManifest(initCtx, 5)
	if err != nil {
		measurement.Logger.Errorf("Upload. Error loading configuration: %v", err)
		return 0, err
	}

	stream, finalize := measurementstream.CreateMeasurementStream(measurement.Configuration.MeasurementRatio, measurementCallback)
	speed, err := manifest.ProbeUploadSpeed(ctx, &client, stream)
	if err != nil {
		measurement.Logger.Errorf("Upload. Error probing speed: %v", err)
		return 0, err
	}

	result := float64(speed) * measurement.Configuration.MeasurementRatio

	finalize(units.BytesPerSecond(result))
	return result, nil
}
