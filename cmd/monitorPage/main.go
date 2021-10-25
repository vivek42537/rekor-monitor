package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/sigstore/rekor-monitor/internal/monitorGRPC"
	"google.golang.org/grpc"
)

func main() {
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
		log.Fatal(err)
	}
	defer conn.Close()

	client := monitorGRPC.NewMonitorServiceClient(conn)
	log.Println("Initialized GRPC Client")

	// Monitor Page Server
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))

		resp, err := client.GetLastSnapshot(context.Background(), &monitorGRPC.LastSnapshotRequest{})
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		tmpl.Execute(rw, resp)
	})

	serverPort := "8000"
	if val, ok := os.LookupEnv("SERVER_PORT"); ok {
		serverPort = val
	}

	log.Printf("Monitor Page Server listening on localhost:%s\n", serverPort)
	err = http.ListenAndServe(fmt.Sprintf(":%s", serverPort), nil)
	if err != nil {
		log.Fatal(err)
	}
}
