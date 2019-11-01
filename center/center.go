package center

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/klbud/flight-go/component"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type Flight struct {
	FlightDB *FlightDB
}

type FlightDB struct {
	DB *xorm.Engine
}

func ParseFlag() string {
	env := flag.String("e", "local", "environment, -e local|dev")
	flag.Parse()
	return *env
}

func InitFlight(env string) *Flight {
	initConfig(env) // init all config
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
	fmt.Printf("flight-go running port %s", CommonCfg.Server.Port)
	fmt.Println()
	http.ListenAndServe(":"+CommonCfg.Server.Port, nil)
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
	fmt.Println("success...")
	response.Write([]byte("success..."))
}

func (f *Flight) WDB() {
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
}

//////////////////////////////////////// config start
type CommonConfig struct {
	Server serverConfig
	DB     dbConfig
}

type serverConfig struct {
	Ip   string
	Port string
}

type dbConfig struct {
	Dialect  string
	Username string
	Password string
	DBName   string
	Charset  string
}

type ComponentServer struct {
	DBEngine *xorm.Engine
}

var (
	// 公共配置
	CommonCfg = &CommonConfig{}
)

//var InitFlightChannel = make(chan *CommonConfig, 10)

func initConfig(env string) {
	var err error
	var cfgPar []byte
	path, err := os.Getwd()
	cfgPar, err = ioutil.ReadFile(path + "/config/" + env + "/config.yaml")
	fmt.Println("config file path: ", path+"/config/"+env+"/config.yaml")
	if err != nil {
		panic(fmt.Errorf("flight-go, init config.yaml err: %v", err))
	}

	if err = yaml.Unmarshal(cfgPar, CommonCfg); err != nil {
		panic(fmt.Errorf("flight-go, yaml.Unmarshal err: %v", err))
	}

	b, _ := json.Marshal(CommonCfg)
	fmt.Println("flight-go, commoncfg:", string(b))
	component.ComInit.InitDB(CommonCfg.DB.Dialect, CommonCfg.DB.Username, CommonCfg.DB.Password, CommonCfg.DB.DBName, CommonCfg.DB.Charset)
}

//////////////////////////////////////// config end
