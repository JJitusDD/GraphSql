package pagination

import (
	"context"
	"time"
)

var (
	defaultLimit = 10
	maxLimit     = 10000
)

func getPaginationLimit(limit int64) int64 {
	if limit <= 0 {
		return int64(defaultLimit)
	}

	if limit > int64(maxLimit) {
		return int64(maxLimit)
	}

	return limit
}

func getSkip(page, limit int64) int64 {
	if page > 0 {
		return (page - 1) * limit
	}
	return page
}

func getCtx() context.Context {
	ctx, _ := context.WithTimeout(context.Background(), 60*time.Second)
	return ctx
}
