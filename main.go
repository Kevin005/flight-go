package main

import (
	"github.com/Kevin005/flight-go/center"
)

func main() {
	f := center.InitFlight()
	f.AddAction("/write", f.WriteDB)
	f.Run()
}
