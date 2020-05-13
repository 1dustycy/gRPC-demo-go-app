package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/betterchen/go-project-tmpl/internal/web"
	"github.com/betterchen/go-project-tmpl/pkg/multiservices"
)

func main() {
	multiservices.Init()
	defer multiservices.Shutdown()

	if err := web.RunServices(); err != nil {
		log.Panic(err)
	}

	time.Sleep(time.Second * 5)

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
