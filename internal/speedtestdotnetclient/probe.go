package speedtestdotnetclient

import (
	"context"

	"go.jonnrb.io/speedtest/speedtestdotnet"
	"go.jonnrb.io/speedtest/units"

	"go.framey.io/speedtest/internal/core"
	"go.framey.io/speedtest/internal/measurementstream"
)

type SpeedtestdotnetEndpointImpl struct{}

func (instance SpeedtestdotnetEndpointImpl) Download(
	ctx context.Context,
	measurement *core.SpeedMeasurement,
	measurementCallback core.MeasurementCallback) (float64, error) {

	var client speedtestdotnet.Client

	initCtx, cancel := context.WithTimeout(context.Background(), measurement.Configuration.ConfigTimeout)
	defer cancel()

	servers, err := enumerateServers(initCtx, measurement, &client)
	if err != nil {
		return 0, err
	}

	targetServer, err := selectServerWithMinLatency(initCtx, measurement, &client, servers)
	if err != nil {
		return 0, err
	}

	stream, finalize := measurementstream.CreateMeasurementStream(measurement.Configuration.MeasurementRatio, measurementCallback)
	speed, err := targetServer.ProbeDownloadSpeed(ctx, &client, stream)
	if err != nil {
		measurement.Logger.Errorf("Download. Error probing download speed: %v", err)
		return 0, err
	}

	result := float64(speed) * measurement.Configuration.MeasurementRatio

	finalize(units.BytesPerSecond(result))
	return result, nil
}

func (instance SpeedtestdotnetEndpointImpl) Upload(
	ctx context.Context,
	measurement *core.SpeedMeasurement,
	measurementCallback core.MeasurementCallback) (float64, error) {

	var client speedtestdotnet.Client

	initCtx, cancel := context.WithTimeout(context.Background(), measurement.Configuration.ConfigTimeout)
	defer cancel()

	servers, err := enumerateServers(initCtx, measurement, &client)
	if err != nil {
		return 0, err
	}

	targetServer, err := selectServerWithMinLatency(initCtx, measurement, &client, servers)
	if err != nil {
		return 0, err
	}

	stream, finalize := measurementstream.CreateMeasurementStream(measurement.Configuration.MeasurementRatio, measurementCallback)
	speed, err := targetServer.ProbeUploadSpeed(ctx, &client, stream)
	if err != nil {
		measurement.Logger.Errorf("Upload. Error probing speed: %v", err)
		return 0, err
	}

	result := float64(speed) * measurement.Configuration.MeasurementRatio

	finalize(units.BytesPerSecond(result))
	return result, nil
}
