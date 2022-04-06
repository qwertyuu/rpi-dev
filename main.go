/*

A blinker example using go-rpio library.
Requires administrator rights to run

Toggles a LED on physical pin 19 (mcu pin 10)
Connect a LED with resistor from pin 19 to ground.

*/

package main

import (
	"fmt"
	"os"
	"time"

	"github.com/stianeikeland/go-rpio/v4"
)

var (
	buildTime string
	// Use mcu pin 10, corresponds to physical pin 19 on the pi
	pin = rpio.Pin(3)
)

func main() {
	fmt.Printf("Build Time: %s\n", buildTime)
	// Open and map memory to access gpio, check for errors
	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Unmap gpio memory when done
	defer rpio.Close()

	// Set pin to output mode
	pin.Input()
	start := int32(time.Now().UnixMilli())
	ref_time := start
	iter := 0
	iter_print := 22700000

	// Toggle pin
	for {
		pin.Read()
		//time.Sleep(time.Second / 5)
		iter++
		if iter%iter_print == 0 {
			now := int32(time.Now().UnixMilli())
			fmt.Printf("I toggled %v times in %v ms\n", iter_print, now-ref_time)
			ref_time = now
			iter = 0
		}
	}
}
