package main

import (
	"github.com/klbud/flight-go/center"
)

func main() {
	env := center.ParseFlag()
	f := center.InitFlight(env)
	f.AddAction("/write", f.WriteDB)
	f.Run()
}
