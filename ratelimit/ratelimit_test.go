package ratelimit_test

import (
	"context"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/Kyagara/equinox/ratelimit"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
)

func TestRateLimitMethods(t *testing.T) {
	t.Parallel()

	// Rate limit is disabled
	rateStore := &ratelimit.RateLimit{}
	require.NotNil(t, rateStore)

	ctx := context.Background()

	err := rateStore.Reserve(ctx, zerolog.Nop(), "route", "method")
	require.Equal(t, ratelimit.ErrRateLimitIsDisabled, err)
	err = rateStore.Update(ctx, zerolog.Nop(), "route", "method", http.Header{}, time.Duration(0))
	require.Equal(t, ratelimit.ErrRateLimitIsDisabled, err)

	rateStore.MarshalZerologObject(&zerolog.Event{})

	rateStore.StoreType = ratelimit.InternalRateLimit
	rateStore.Enabled = true

	var logger zerolog.Logger
	logger = logger.Output(zerolog.ConsoleWriter{Out: os.Stderr}).With().Object("ratelimit", rateStore).Logger()
	logger.Info().Msg("Testing rate limit marshal")
}
