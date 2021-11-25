package speedtestdotnetclient

import (
	"context"

	"go.framey.io/speedtest/internal/core"
	"go.jonnrb.io/speedtest/speedtestdotnet"
)

func enumerateServers(
	ctx context.Context,
	measurement *core.SpeedMeasurement,
	client *speedtestdotnet.Client,
) ([]speedtestdotnet.Server, error) {
	servers, err := client.LoadAllServers(ctx)
	if err != nil {
		measurement.Logger.Errorf("Failed to load server list: %v", err)
		return nil, err
	}
	if len(servers) == 0 {
		measurement.Logger.Errorf("No servers found somehow...")
		return nil, err
	}
	return servers, nil
}
