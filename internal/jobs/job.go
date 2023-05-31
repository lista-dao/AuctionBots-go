package jobs

import (
	"context"
	"sync"
)

type Job interface {
	Run(ctx context.Context, wg *sync.WaitGroup)
}
