package center

import (
	"fmt"
	"net/http"
	"testing"
	"time"
)

func (f *Flight) writeDB(response http.ResponseWriter, request *http.Request) {
	fmt.Println("flit-go, writeDB start")
	//sql := "update `userinfo` set username=? where id=?"
	sql := "insert into `test` (`name`,`age`) values (?,?)"
	f.FlightDB.DB.Exec(sql, "Davie", 30)

	go func() {
		for {
			f.FlightDB.DB.Exec(sql, "Davie", 30)
			time.Sleep(1 * time.Millisecond)
		}
	}()

	<-time.After(1 * time.Hour)
	response.Write([]byte("day day up"))
}

func TestRun(t *testing.T) {
	f := InitFlight()
	f.AddAction("/db/write", f.writeDB)
	f.Run()
}

func TestFlight_WDB(t *testing.T) {
	sql := "insert into `test` (`name`,`age`) values (?,?)"
	f := &Flight{}
	f.FlightDB.DB.Exec(sql, "Davie", 30)
}
