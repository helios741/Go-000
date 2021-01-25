package main

import (
	"Week03/shutdown"
	s "Week03/signal"
	"Week03/workgroup"
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"
)

type Foo struct {
	X int
}

type Bar struct {
	X int
}

type Baz struct {
	X int
}

// ProvideFoo returns a Foo.
func ProvideFoo() Foo {
	return Foo{X: 42}
}

// ProvideBar returns a Bar: a negative Foo.
func ProvideBar(foo Foo) Bar {
	return Bar{X: -foo.X}
}

// ProvideBaz returns a value if Bar is not zero.
func ProvideBaz(ctx context.Context, bar Bar) (Baz, error) {
	if bar.X == 0 {
		return Baz{}, errors.New("cannot provide baz when bar is zero")
	}
	return Baz{X: bar.X}, nil
}

//func initializeBaz(ctx context.Context) (Baz, error) {
//	var SuperSet = wire.NewSet(ProvideFoo, ProvideBar, ProvideBaz)
//	wire.Build(SuperSet)
//	return Baz{}, nil
//}
//
//func main() {
//	ctx := context.Background()
//	aa, _ := initializeBaz(ctx)
//	fmt.Println(aa.X)
//}
func main() {
	var wg workgroup.Group

	wg.Add(s.New(os.Interrupt))
	srv1 := http.Server{Addr: "127.0.0.1:8080"}
	srv2 := http.Server{Addr: "127.0.0.1:8081"}
	wg.Add(shutdown.New(func() error {
			fmt.Printf("Server1 is about to listen at %v\n", srv1.Addr)
			return srv1.ListenAndServe()
		},
		func() {
			fmt.Printf("Server1 is get quit signal \n")
			ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
			defer cancel()

			_ = srv1.Shutdown(ctx)
			fmt.Printf("Server1 shutdown\n")
		},
	))

	//wg.Add(shutdown.New(func() error {
	//	time.Sleep(6 * time.Second)
	//	return errors.New("fake news")
	//}, func() {
	//	fmt.Println("fake news fail")
	//}))

	wg.Add(shutdown.New(func() error {
			fmt.Printf("Server2 is about to listen at %v\n", srv2.Addr)
			return srv2.ListenAndServe()
		},
		func() {
			fmt.Printf("Server2 is get quit signal \n")
			ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
			defer cancel()

			_ = srv2.Shutdown(ctx)
			fmt.Printf("Server2 shutdown \n")
		},
	))

	err := wg.Run()
	fmt.Printf("Workgroup quit with error: %v\n", err)
}