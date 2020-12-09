package signal

import (
	wg "Week03/workgroup"
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
)

func New(sig ...os.Signal) wg.Run  {
	return func(ctx context.Context) error {
		if len(sig) == 0 {
			sig = append(sig, os.Interrupt)
		}
		done := make(chan os.Signal, len(sig))
		signal.Notify(done, sig...)
		is_done := false
		select {
		case <-done:
			is_done = true
		case <- ctx.Done():
			is_done = false
		}
		signal.Stop(done)
		close(done)
		fmt.Printf("\nsignal quit\n")
		if !is_done {
			return nil
		}
		return errors.New("quit signal")
	}
	
}
