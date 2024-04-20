package grpcclient

import (
	"context"
	"fmt"
	"github.com/nezhahq/agent/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"log"
)

// SendSystemState sends system state to the server
func SendSystemState(ctx context.Context, client proto.NezhaServiceClient, results []ResultItem, clientSecret string) error {
	// Construct authentication metadata
	md := metadata.Pairs("client_secret", clientSecret)
	ctx = metadata.NewOutgoingContext(ctx, md)

	// Send each result item to the server
	for _, result := range results {
		uptime := uint64(time.Since(time.Unix(int64(result.LastActive), 0)).Seconds())
		result.Status.Uptime = uptime

		randomizeStatus(&result.Status)

		_, err := client.ReportSystemState(ctx, &result.Status)
		if err != nil {
			return fmt.Errorf("failed to report system state: %v", err)
		}
	}

	return nil
}

// SendHostInfo sends host info to the server
func SendHostInfo(ctx context.Context, client proto.NezhaServiceClient, results []ResultItem, clientSecret string) error {
	// Construct authentication metadata
	md := metadata.Pairs("client_secret", clientSecret)
	ctx = metadata.NewOutgoingContext(ctx, md)

	// Send each result item to the server
	for _, result := range results {
		result.Host.BootTime = result.LastActive

		_, err := client.ReportSystemInfo(ctx, &result.Host)
		if err != nil {
			return fmt.Errorf("failed to report host info: %v", err)
		}
	}

	return nil
}