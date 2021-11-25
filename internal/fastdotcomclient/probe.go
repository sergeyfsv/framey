package fastdotcomclient

import (
	"context"
	//	"fmt"
	"log"
	//	"time"

	"go.jonnrb.io/speedtest/fastdotcom"
	//	"go.jonnrb.io/speedtest/oututil"
	"go.jonnrb.io/speedtest/units"
	"golang.org/x/sync/errgroup"

	"go.framey.io/speedtest/internal/types"
)

type FastdotcomSpeedtestEndpoint struct{}

func (instance FastdotcomSpeedtestEndpoint) Download(
	ctx context.Context,
	cfg *types.Configuration,
	measurementCallback types.MeasurementCallback) (float64, error) {

	var client fastdotcom.Client

	m, err := fastdotcom.GetManifest(ctx, 5)
	if err != nil {
		log.Fatalf("Error loading fast.com configuration: %v", err)
		return 0, err
	}

	stream, finalize := createMetricsStream(cfg.MeasurementRatio, measurementCallback)
	speed, err := m.ProbeDownloadSpeed(ctx, &client, stream)
	if err != nil {
		log.Fatalf("Error probing download speed: %v", err)
		return 0, err
	}
	result := float64(speed) * cfg.MeasurementRatio

	finalize(units.BytesPerSecond(result))
	return result, nil
}

func (instance FastdotcomSpeedtestEndpoint) Upload(
	ctx context.Context,
	cfg *types.Configuration,
	measurementCallback types.MeasurementCallback) (float64, error) {

	var client fastdotcom.Client

	m, err := fastdotcom.GetManifest(ctx, 5)
	if err != nil {
		log.Fatalf("Error loading fast.com configuration: %v", err)
		return 0, err
	}

	stream, finalize := createMetricsStream(cfg.MeasurementRatio, measurementCallback)
	speed, err := m.ProbeUploadSpeed(ctx, &client, stream)
	if err != nil {
		log.Fatalf("Error probing upload speed: %v", err)
		return 0, err
	}

	result := float64(speed) * cfg.MeasurementRatio

	finalize(units.BytesPerSecond(result))
	return result, nil
}

func createMetricsStream(measurementRatio float64, measurementCallback types.MeasurementCallback) (
	stream chan units.BytesPerSecond,
	finalize func(units.BytesPerSecond),
) {
	measurementCallback(float64(0) * measurementRatio)
	/*
		p := oututil.StartPrinting()
		p.Println(measurementCallback(units.BytesPerSecond(0)))
	*/
	stream = make(chan units.BytesPerSecond)
	var g errgroup.Group
	g.Go(func() error {
		for speed := range stream {
			measurementCallback(float64(speed) * measurementRatio)
			//fmt.Println(measurementCallback(float64(speed)))
			//p.Println(format(speed))
		}
		return nil
	})

	finalize = func(s units.BytesPerSecond) {
		g.Wait()
		measurementCallback(float64(s))
		//p.Finalize(measurementCallback(s))
	}
	return
}
