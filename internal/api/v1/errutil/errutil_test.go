package errutil_test

import (
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/bsv-blockchain/spv-wallet-go-client/internal/api/v1/errutil"
)

var errServerError = errors.New(http.StatusText(http.StatusInternalServerError))

func TestHTTPErrorFormatter_Format(t *testing.T) {
	// given:
	const (
		API    = "Users API"
		action = "retrieve users page"
	)
	expectedErr := fmt.Errorf("failed to send HTTP %s request to %s via %s: %w", http.MethodPost, action, API, errServerError)

	formatter := errutil.HTTPErrorFormatter{
		Action: action,
		API:    API,
		Err:    errServerError,
	}

	// when:
	got := formatter.Format(http.MethodPost)

	// then:
	require.Equal(t, expectedErr, got)
}
