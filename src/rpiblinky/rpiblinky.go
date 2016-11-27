package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/kidoman/embd"
	_ "github.com/kidoman/embd/host/rpi"
)

func main() {
	flag.Parse()

	// Init the GPIO lib
	if err := embd.InitGPIO(); err != nil {
		panic(err)
	}
	defer embd.CloseGPIO()

	// Create a new interface to Pin "GPIO_4"
	gpio4_pin, err := embd.NewDigitalPin("GPIO_4")
	if err != nil {
		panic(err)
	}
	// Set Pin "GPIO_4" as an out pin
	if err := gpio4_pin.SetDirection(embd.Out); err != nil {
		panic(err)
	}

	// Create a channel to receive all interrupt and kill(ctrl+c) signals
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Kill)
	defer signal.Stop(quit)

	for {
		select {
		// Turn On for 1 sec, then turn off for 1 sec
		case <-time.After(1 * time.Second):
			if err := gpio4_pin.Write(embd.High); err != nil {
				panic(err)
			}
			fmt.Println("Toggled")
			time.Sleep(1 * time.Second)
			if err := gpio4_pin.Write(embd.Low); err != nil {
				panic(err)
			}
		// When a signal is caught the program will exit
		case <-quit:
			return
		}
	}
}
