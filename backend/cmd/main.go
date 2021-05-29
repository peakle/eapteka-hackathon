package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/peakle/eapteka/internal/server"
)

type Process interface {
	Start(ctx context.Context) error
}
type Processors []Process

func main() {
	var wg sync.WaitGroup
	var ctx, cancel = context.WithCancel(context.Background())
	var sigs = make(chan os.Signal)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigs
		cancel()
	}()

	container := provider.NewContainer()
	ps := Processors{
		consumer.NewConsumer(container),
		server.NewHandler(container),
	}

	for _, p := range ps {
		wg.Add(1)

		go func(process Process) {
			defer wg.Done()

			err := process.Start(ctx)
			if err != nil {
				log.Printf("on main: %s", err)
			}
		}(p)
	}

	wg.Wait()
}
