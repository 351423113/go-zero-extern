package redistest

import (
	"time"

	"github.com/351423113/go-zero-extern/core/lang"
	"github.com/351423113/go-zero-extern/core/stores/redis"
	"github.com/alicebob/miniredis/v2"
)

// CreateRedis returns a in process redis.Redis.
func CreateRedis() (r *redis.Redis, clean func(), err error) {
	mr, err := miniredis.Run()
	if err != nil {
		return nil, nil, err
	}

	return redis.NewRedis(mr.Addr(), redis.NodeType), func() {
		ch := make(chan lang.PlaceholderType)
		go func() {
			mr.Close()
			close(ch)
		}()
		select {
		case <-ch:
		case <-time.After(time.Second):
		}
	}, nil
}
