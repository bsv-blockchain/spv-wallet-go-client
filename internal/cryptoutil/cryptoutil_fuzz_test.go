package cryptoutil_test

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/bsv-blockchain/spv-wallet-go-client/internal/cryptoutil"
)

func FuzzParseChildNumsFromHex(f *testing.F) {
	// Seed corpus with known inputs from existing tests
	f.Add("")
	f.Add("ababab")
	f.Add("8bb0cf6eb9b17d0f7d22b456f121257dc1254e1f01665370476383ea776df414")
	f.Add("FFFFFFFF")
	f.Add("test") // Invalid hex
	f.Add("00000000")
	f.Add("a")
	f.Add("ab")

	f.Fuzz(func(t *testing.T, hex string) {
		// Function should not panic on any input
		_, _ = cryptoutil.ParseChildNumsFromHex(hex)
	})
}

func FuzzInt64ToUint32(f *testing.F) {
	// Seed corpus with boundary values
	f.Add(int64(0))
	f.Add(int64(-1))
	f.Add(int64(1))
	f.Add(int64(math.MaxInt32))
	f.Add(int64(math.MaxUint32))
	f.Add(int64(math.MaxUint32 + 1))
	f.Add(int64(math.MinInt64))
	f.Add(int64(math.MaxInt64))

	f.Fuzz(func(t *testing.T, value int64) {
		// Function should not panic on any input
		_, _ = cryptoutil.Int64ToUint32(value)
	})
}

func FuzzHash(f *testing.F) {
	// Seed corpus with various strings
	f.Add("")
	f.Add("hello")
	f.Add("1234567")
	f.Add("xpub661MyMwAqRbcFrBJbKwBGCB7d3fr2SaAuXGM95BA62X41m6eW2ehRQGW4xLi9wkEXUGnQZYxVVj4PxXnyrLk7jdqvBAs1Qq9gf6ykMvjR7J")
	f.Add(string([]byte{0, 1, 2, 255})) // Binary data

	f.Fuzz(func(t *testing.T, input string) {
		result := cryptoutil.Hash(input)
		// Hash should always return 64 char hex string (SHA256 = 32 bytes = 64 hex chars)
		require.Len(t, result, 64)
	})
}
