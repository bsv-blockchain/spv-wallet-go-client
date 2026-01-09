package main

import (
	"context"
	"log"

	"github.com/bsv-blockchain/spv-wallet/models/filter"

	wallet "github.com/bsv-blockchain/spv-wallet-go-client"
	"github.com/bsv-blockchain/spv-wallet-go-client/examples"
	"github.com/bsv-blockchain/spv-wallet-go-client/examples/exampleutil"
	"github.com/bsv-blockchain/spv-wallet-go-client/queries"
)

func main() {
	usersAPI, err := wallet.NewUserAPIWithXPriv(exampleutil.NewDefaultConfig(), examples.UserXPriv)
	if err != nil {
		log.Fatalf("Failed to initialize user API with XPriv: %v", err)
	}

	page, err := usersAPI.Contacts(context.Background(), queries.QueryWithPageFilter[filter.ContactFilter](filter.Page{
		Size: 1,
		Sort: "asc",
	}))
	if err != nil {
		log.Fatalf("Failed to fetch contacts: %v", err)
	}
	exampleutil.PrettyPrint("Fetched contacts", page.Content)
}
