package fib

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/evergreen-ci/poplar"
)

func BenchmarkFib(b *testing.B) {
	b.Run("Using Go", func(b *testing.B) {
		for _, length := range []int{5, 10} {
			b.Run(fmt.Sprintf("fib_%d", length), func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					Fib(length)
				}
			})
		}
	})

	suite := poplar.BenchmarkSuite{newTestBenchmark(5), newTestBenchmark(10)}
	b.Run("Using Poplar", suite.Standard(poplar.NewRegistry()))
}

func newTestBenchmark(length int) *poplar.BenchmarkCase {
	return &poplar.BenchmarkCase{
		CaseName: fmt.Sprintf("fib_%d", length),
		Bench: func(ctx context.Context, r poplar.Recorder, count int) error {
			for i := 0; i < count; i++ {
				startAt := time.Now()
				r.BeginIteration()
				Fib(length)
				r.EndIteration(time.Since(startAt))
				r.IncOperations(int64(length))
				// r.IncSize(len(val))
			}
			os.Remove(fmt.Sprintf("fib_%d.ftdc", length))
			return nil
		},
		MinIterations: 200,
		MaxIterations: 500,
		MinRuntime:    1 * time.Second,
		MaxRuntime:    30 * time.Second,
		Recorder:      poplar.RecorderPerf,
	}
}
