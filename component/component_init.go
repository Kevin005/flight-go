package component

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var (
	ComInit = &ComponentInit{}
	Com     = &Component{}
)

type Component struct {
	DB *xorm.Engine
}

type ComponentInit struct{}

func (com *ComponentInit) InitDB(dialect, username, password, dbname, charset string) {
	fmt.Println("start InitDB")
	var err error
	var dbEngine *xorm.Engine

	//fmt.Sprintf("%s:%s@/%s?charset=%s")
	dbEngine, err = xorm.NewEngine(dialect, username+":"+password+"@/"+dbname+"?charset="+charset)
	if err != nil {
		panic(fmt.Errorf("flight-go, InitFlight initDB err, %v", err))
	}
	Com.DB = dbEngine
}
