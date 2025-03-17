package middleware

import (
	"solver/internal/solver"
)

type Middleware struct {
	apiKeyLimiter *KeyRateLimiter
	solver        solver.Solver
}

func New(s solver.Solver) *Middleware {
	apiKeyLimiter := NewKeyRateLimiter()

	return &Middleware{
		apiKeyLimiter: apiKeyLimiter,
		solver:        s,
	}
}
