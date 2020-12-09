package main

import (
	"Week03/shutdown"
	s "Week03/signal"
	"Week03/workgroup"
	"context"
	"fmt"
	"net/http"
	"os"
	"time"
)


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