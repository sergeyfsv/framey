package fastdotcomclient

import (
	"context"
	//"log"
	"fmt"
	"time"

	"go.framey.io/speedtest/internal/types"
	"go.jonnrb.io/speedtest/units"
)

func formatSpeed(s float64) string {
	return fmt.Sprintf("%v", units.BytesPerSecond(s).BitsPerSecond())
}

func Main(args []string) {
	/*
		err := flagSet.Parse(args[1:])
		if err != nil {
			panic(err)
		}
	*/

	endpoint := FastdotcomSpeedtestEndpoint{}
	cfg := types.DefaultConfiguration()

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.ConfigTimeout+cfg.DownloadTimeout))
	defer cancel()

	endpoint.Download(ctx, cfg, func(s float64) string {
		return formatSpeed(s)
	})

	/*
		var client fastdotcom.Client

		ctx, cancel := context.WithTimeout(context.Background(), *cfgTime)
		defer cancel()

		m, err := fastdotcom.GetManifest(ctx, *urlCount)
		if err != nil {
			log.Fatalf("Error loading fast.com configuration: %v", err)
		}

		//download(m, &client)
		//upload(m, &client)
	*/
}
