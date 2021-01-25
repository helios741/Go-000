package workgroup

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
)

type Run func(c context.Context) error

type Group struct {
	fns []Run
}

func (g *Group) Add(fn Run)  {
	g.fns = append(g.fns, fn)
}


func (g *Group)Run() error {
	if len(g.fns) == 0 {
		return nil
	}
	eg, ctx := errgroup.WithContext(context.Background())

	go func(ctx context.Context) {
		<- ctx.Done()
		fmt.Println("xxxxx")

	}(ctx)
	for _, f :=  range g.fns {
		fn := f
		eg.Go(func() error {
			return fn(ctx)
		})
	}

	if err := eg.Wait(); err != nil {
		fmt.Println("close all:", err)
		return err
	}
	return nil
}