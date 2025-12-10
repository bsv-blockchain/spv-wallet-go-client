package main

import (
	"fmt"
	"log"

	"github.com/bitcoin-sv/spv-wallet-go-client/walletkeys"
)

func main() {
	keys, err := walletkeys.RandomKeysWithMnemonic()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("XPriv: ", keys.XPriv())       //nolint: forbidigo // example output
	fmt.Println("XPub: ", keys.XPub())         //nolint: forbidigo // example output
	fmt.Println("Mnemonic: ", keys.Mnemonic()) //nolint: forbidigo // example output
}
