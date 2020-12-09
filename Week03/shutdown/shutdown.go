package shutdown

import (
	wg "Week03/workgroup"
	"context"
)

func New(run func() error, quit func()) wg.Run {
	return func(ctx context.Context) error {
		go func() {

			<- ctx.Done()
			quit()
		}()
		return run()
	}
}