package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/sigstore/rekor-monitor/internal/monitorGRPC"
	"google.golang.org/grpc"
)

var grpcClient monitorGRPC.MonitorServiceClient

func initClient() error {
	// gRPC client
	grpcPort := "9000"
	if val, ok := os.LookupEnv("GRPC_PORT"); ok {
		grpcPort = val
	}

	grpcContainer := "monitor-grpc"
	if val, ok := os.LookupEnv("GRPC_CONTAINER"); ok {
		grpcContainer = val
	}

	grpcEndpoint := fmt.Sprintf("%s:%s", grpcContainer, grpcPort)

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithBlock())
	opts = append(opts, grpc.FailOnNonTempDialError(true))

	conn, err := grpc.Dial(grpcEndpoint, opts...)
	if err != nil {
		return err
	}
	defer conn.Close()

	grpcClient = monitorGRPC.NewMonitorServiceClient(conn)
	return nil
}

func getLastSnapshot() {
	// TODO: Handle error
	resp, _ := grpcClient.GetLastSnapshot(context.Background(), &monitorGRPC.LastSnapshotRequest{})

}

func main() {
	err := initClient()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Initialized GRPC Client")

}
