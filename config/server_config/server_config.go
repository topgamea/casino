package server_config

import (
	"io/ioutil"
	"github.com/BurntSushi/toml"
)

type CasinoConfig struct {
	Title            string           `toml:"title"`
	CQGameConfig     CqGameConfig     `toml:"cqgame"`
	MachinesConfig   MachinesConfig   `toml:"machines"`
	HTTPConfig       HttpConfig       `toml:"http"`
	DBConfig         MysqlConfig      `toml:"mysql"`
	RedisConfig      RedisConfig      `toml:"redis"`
	InitConfig       InitConfig       `toml:"init"`
	BIConfig         BiConfig         `toml:"bi"`
	GameDetailConfig GameDetailConfig `toml:"gamedetail"`
}

type MysqlConfig struct {
	Host        string `toml:"host"`
	Port        int    `toml:"port"`
	DBName      string `toml:"db"`
	User        string `toml:"user"`
	Pwd         string `toml:"password"`
	MaxIdleConn int    `toml:"max_idle_conns"`
}

type BiConfig struct {
	UserStartID  int `toml:"user_start_id"`
	UserCount    int `toml:"user_count"`
	CountPerUser int `toml:"count_per_user"`
}

type InitConfig struct {
	Gold     int `toml:"gold"`
	Debug    int `toml:"debug"`
	ServerId int `toml:"server_id"`
}

type CqGameConfig struct {
	URL  string `toml:"url"`
	Auth string `toml:"auth"`
	Code string `toml:"code"`
}

type MachinesConfig struct {
	Config           string `toml:"config"`
	SpecialConfig    string `toml:"special_config"`
	ENV              string `toml:"env"`
	UnActiveInterval int    `toml:"unactive_interval"`
	RewardUpLimit    int    `toml:"reward_up_limit"`
}

type RedisConfig struct {
	Host string `toml:"host"`
	Port int    `toml:"port"`
	DB   int    `toml:"db"`
	User string `toml:"user"`
	Pwd  string `toml:"password"`
}

type HttpConfig struct {
	Host string `toml:"host"`
	Port int    `toml:"port"`
}

type GameDetailConfig struct {
	Url string `toml:"url"`
}

func ParseConfig(filename string) (*CasinoConfig, error) {
	var config CasinoConfig
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	_, err = toml.Decode(string(data), &config)

	if err != nil {
		return nil, err
	}
	return &config, nil
}

