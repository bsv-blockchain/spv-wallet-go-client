package main

import (
	"context"
	"log"

	wallet "github.com/bsv-blockchain/spv-wallet-go-client"
	"github.com/bsv-blockchain/spv-wallet-go-client/examples"
	"github.com/bsv-blockchain/spv-wallet-go-client/examples/exampleutil"
)

func main() {
	usersAPI, err := wallet.NewUserAPIWithXPriv(exampleutil.NewDefaultConfig(), examples.UserXPriv)
	if err != nil {
		log.Fatalf("Failed to initialize user API with XPriv: %v", err)
	}

	ctx := context.Background()
	cfg, err := usersAPI.SharedConfig(ctx)
	if err != nil {
		log.Fatalf("Failed to shared config: %v", err)
	}

	exampleutil.PrettyPrint("Fetched shared config", cfg)
}
