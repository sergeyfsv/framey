package speedtestdotnetclient

import (
	"context"

	"go.framey.io/speedtest/internal/core"
	"go.jonnrb.io/speedtest/speedtestdotnet"
)

func selectServerWithMinLatency(
	ctx context.Context,
	measurement *core.SpeedMeasurement,
	client *speedtestdotnet.Client,
	servers []speedtestdotnet.Server) (speedtestdotnet.Server, error) {

	const MAX_SERVERS_TO_TEST_LATENCY = 5
	serversToTestLatency := func() []speedtestdotnet.Server {
		if len(servers) > MAX_SERVERS_TO_TEST_LATENCY {
			return servers[:MAX_SERVERS_TO_TEST_LATENCY]
		} else {
			return servers
		}
	}()

	_, err := speedtestdotnet.StableSortServersByAverageLatency(
		serversToTestLatency, ctx, client, speedtestdotnet.DefaultLatencySamples)
	if err != nil {
		measurement.Logger.Errorf("Error getting server latencies: %v", err)
		return speedtestdotnet.Server{}, err
	}

	return serversToTestLatency[0], nil
}
