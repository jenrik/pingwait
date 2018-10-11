package main

import (
	"flag"
	"fmt"
	"github.com/sparrc/go-ping"
	"os"
	"time"
)

/**
 * Exit codes:
 * 0: received ping response
 * 1: invalid arguments
 * 2: timeout
 * 50: internal error
 */

func main() {
	timeout := flag.Int("t", 0, "Timeout")
	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Println("Missing address")
		os.Exit(1)
	}
	if flag.NArg() > 1 {
		fmt.Println("Too many address")
		os.Exit(1)
	}
	addr := flag.Arg(len(flag.Args()) - 1)

	if timeout != nil && *timeout > 0 {
		timeoutHandler := time.NewTimer(time.Second * time.Duration(*timeout))
		go func() {
			<-timeoutHandler.C
			os.Exit(2)
		}()
	}

	pingSendTicker := time.NewTicker(time.Second)
	for _ = range pingSendTicker.C {
		// Send ping
		pinger, err := ping.NewPinger(addr)
		if err != nil {
			os.Exit(50)
		}

		pinger.SetPrivileged(true)
		pinger.Timeout = time.Second * 5
		pinger.Count = 1
		pinger.OnRecv = func(_ *ping.Packet) {
			os.Exit(0)
		}

		pinger.Run()
	}
}
