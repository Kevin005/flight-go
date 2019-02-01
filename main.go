package main

import (
	"github.com/Kevin005/flight-go/center"
)

func main() {
	f := center.InitFlight()
	f.AddAction("/db/write", f.WriteDB)
	f.WDB() // to do
	f.Run()
}
