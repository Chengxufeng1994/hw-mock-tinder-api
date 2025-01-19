package ratelimiter

import "go.uber.org/ratelimit"

type RateLimiter interface {
	Allow(key string) bool
}

type UberRateLimiter struct {
	ratelimit.Limiter
}

var _ RateLimiter = (*UberRateLimiter)(nil)

func NewUberRateLimiter(ratelimiter ratelimit.Limiter) *UberRateLimiter {
	return &UberRateLimiter{Limiter: ratelimiter}
}

func (u UberRateLimiter) Allow(key string) bool {
	u.Limiter.Take()
	return true
}
