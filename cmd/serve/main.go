package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/betterchen/go-project-tmpl/internal/web"
	"github.com/betterchen/go-project-tmpl/pkg/multiservices"
)

func main() {
	ctx, cancle := context.WithCancel(context.Background())
	defer cancle()

	multiservices.Init()
	defer multiservices.Shutdown()

	if err := web.RunServices(ctx); err != nil {
		log.Panic(err)
	}

	done := make(chan error)

	// intercepts with ctrl-c
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT)
		signal.Notify(c, syscall.SIGTERM)
		done <- fmt.Errorf("%s", <-c)
	}()

	// block
	log.Printf("terminated %v", <-done)
}
