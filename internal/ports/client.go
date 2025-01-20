package ports

import (
	"context"
)

type Cache interface {
	Get(ctx context.Context, key string) (interface{}, bool)
	Set(ctx context.Context, key string, value interface{}, ttlSeconds int)
	Delete(ctx context.Context, key string)
}
