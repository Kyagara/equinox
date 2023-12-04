package ratelimit

import (
	"strings"
	"sync"
	"time"
)

// Limit represents a collection of buckets and the type of limit (application or method).
type Limit struct {
	buckets   []*Bucket
	limitType string
	mutex     sync.Mutex
}

func NewLimit(limitType string) *Limit {
	return &Limit{
		buckets:   make([]*Bucket, 0),
		limitType: limitType,
		mutex:     sync.Mutex{},
	}
}

// Checks if the limits given in the header match the current buckets.
func (l *Limit) limitsDontMatch(limitHeader string) bool {
	if limitHeader == "" {
		return false
	}
	limits := strings.Split(limitHeader, ",")
	if len(l.buckets) != len(limits) {
		return true
	}
	for i, pair := range limits {
		if l.buckets[i] == nil {
			return true
		}
		limit, interval := getNumbersFromPair(pair)
		if l.buckets[i].limit != limit || l.buckets[i].interval != time.Duration(interval)*time.Second {
			return true
		}
	}
	return false
}
