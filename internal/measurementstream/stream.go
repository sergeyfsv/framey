package measurementstream

import (
	"go.framey.io/speedtest/internal/core"
	"go.jonnrb.io/speedtest/units"
	"golang.org/x/sync/errgroup"
)

func CreateMeasurementStream(measurementRatio float64, measurementCallback core.MeasurementCallback) (
	stream chan units.BytesPerSecond,
	finalize func(units.BytesPerSecond),
) {
	measurementCallback(float64(0) * measurementRatio)
	stream = make(chan units.BytesPerSecond)
	var g errgroup.Group
	g.Go(func() error {
		for speed := range stream {
			measurementCallback(float64(speed) * measurementRatio)
		}
		return nil
	})

	finalize = func(s units.BytesPerSecond) {
		g.Wait()
		measurementCallback(float64(s))
	}
	return
}
