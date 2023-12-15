package ratelimit

import (
	"context"
	"strconv"
	"strings"
	"time"
)

// WaitN waits for the given duration after checking if the context deadline will be exceeded.
func WaitN(ctx context.Context, estimated time.Time, duration time.Duration) error {
	deadline, ok := ctx.Deadline()
	if ok && deadline.Before(estimated) {
		return ErrContextDeadlineExceeded
	}

	select {
	case <-time.After(duration):
	case <-ctx.Done():
		return ctx.Err()
	}

	return nil
}

// getNumbersFromPair returns the limit and interval (in seconds) from a pair of numbers separated by a colon.
func getNumbersFromPair(pair string) (int, time.Duration) {
	numbers := strings.Split(pair, ":")
	interval, _ := strconv.Atoi(numbers[1])
	limitOrCount, _ := strconv.Atoi(numbers[0])
	return limitOrCount, time.Duration(interval) * time.Second
}
