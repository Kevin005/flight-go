package main

import (
	"flag"
	"github.com/Kevin005/flight-go/center"
)

func main() {
	env := flag.String("e", "local", "environment, -e local|dev")
	flag.Parse()
	f := center.InitFlight(*env)
	f.AddAction("/write", f.WriteDB)
	f.Run()
}
