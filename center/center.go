package center

import (
	"fmt"
	"github.com/Kevin005/flight-go/component"
	"github.com/Kevin005/flight-go/config"
	"github.com/go-xorm/xorm"
	"net/http"
	"time"
)

type Flight struct {
	FlightDB *FlightDB
}

type FlightDB struct {
	DB *xorm.Engine
}

func InitFlight() *Flight {
	return &Flight{
		FlightDB: &FlightDB{
			DB: component.Com.DB,
		},
	}
}

func (f *Flight) AddAction(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	http.HandleFunc(pattern, handler)
}

//type Handler struct {
//	handler func(http.ResponseWriter, *http.Request)
//	DB      *FlightDB
//}

func (f *Flight) Run() {
	fmt.Println("======flight running")
	http.ListenAndServe(":"+config.CommonCfg.Server.Port, nil)
}

func (f *Flight) WriteDB(response http.ResponseWriter, request *http.Request) {
	//fmt.Println("flit-go, writeDB start")
	////sql := "update `userinfo` set username=? where id=?"
	//sql := "insert into `test` (`name`,`age`) values (?,?)"
	//f.FlightDB.DB.Exec(sql, "Davie", 30)
	//
	//go func() {
	//	for {
	//		f.FlightDB.DB.Exec(sql, "Davie", 30)
	//		time.Sleep(1 * time.Second)
	//	}
	//}()
	//
	//<-time.After(1 * time.Hour)

	//f.FlightDB.DB.Where("id >?", 1).Count(user)
	response.Write([]byte("day day up"))
}

func (f *Flight) WDB() {
	fmt.Println("flit-go, writeDB start")
	//sql := "update `userinfo` set username=? where id=?"
	sql := "insert into `test` (`name`,`age`) values (?,?)"
	f.FlightDB.DB.Exec(sql, "Davie", 30)

	go func() {
		for {
			f.FlightDB.DB.Exec(sql, "Davie", 30)
			time.Sleep(1 * time.Second)
		}
	}()
	<-time.After(1 * time.Hour)
}
