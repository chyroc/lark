package test

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/chyroc/lark"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randInt64() int64 {
	return rand.Int63()
}

func mockGetTenantAccessTokenFailed(ctx context.Context) (*lark.TokenExpire, *lark.Response, error) {
	return nil, nil, fmt.Errorf("failed")
}
