package config

import (
	"encoding/json"
	"fmt"
	"github.com/Kevin005/flight-go/component"
	"io/ioutil"

	"github.com/go-xorm/xorm"
	"gopkg.in/yaml.v2"
)

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

func init() {
	var err error
	var cfgPar []byte
	cfgPar, err = ioutil.ReadFile("config.yaml")
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
