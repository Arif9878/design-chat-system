package main

import (
	"context"
	"log"
	"os"

	"github.com/urfave/cli/v2"

	"github.com/Arif9878/design-chat-system/go/grpc-uberfx/cmd/grpc"
	"github.com/Arif9878/design-chat-system/go/grpc-uberfx/internal/config"
)

func main() {
	cfg, err := config.LoadConfig("config.json")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	ctx := context.Background()
	app := &cli.App{
		Name: "gRPC server",
		Commands: []*cli.Command{
			{
				Name:            "grpc",
				Usage:           "Start the grpc service",
				Action:          grpc.Exec(ctx, cfg),
				SkipFlagParsing: true,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
