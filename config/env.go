package config

import (
	"encoding/json"
	. "tyrannosaurs/constant"
)

const config = `{
  "alpha" : {
    "mysql": {
      "host": "127.0.0.1",
      "port": "3306",
      "db": "tyrannosaurus",
      "username": "root",
      "password": "king258angel",
      "maxIdleConns": 20,
      "maxOpenConns": 20
    },
    "logPath": "./log"
  },
  "beta" : {
    "mysql": {
      "host": "127.0.0.1",
      "port": "3306",
      "db": "tyrannosaurus",
      "username": "root",
      "password": "king258angel",
      "maxIdleConns": 20,
      "maxOpenConns": 20
    },
    "logPath": "./log"
  },
  "prod" : {
    "mysql": {
      "host": "127.0.0.1",
      "port": "3306",
      "db": "tyrannosaurus",
      "username": "root",
      "password": "king258angel",
      "maxIdleConns": 20,
      "maxOpenConns": 20
    },
    "logPath": "./log"
  }
}`

const (
	Alpha = "alpha"
	Beta  = "beta"
	Prod  = "prod"
)

type Environment struct {
	Alpha Config
	Beta  Config
	Prod  Config
}

type Config struct {
	Mysql   Mysql
	LogPath string
}

type Mysql struct {
	Host         string
	Port         string
	DB           string
	Username     string
	Password     string
	MaxIdleConns int
	MaxOpenConns int
}

var Env Environment

func init() {
	//data, err := ioutil.ReadFile("./env.json")
	//if err != nil {
	//	panic(OpenEnvfileError)
	//}
	if err := json.Unmarshal([]byte(config), &Env); err != nil {
		panic(EnvToStructError)
	}
}

func (e Environment) GetConfig(env string) (Config, error) {
	switch env {
	case Alpha:
		return e.Alpha, nil
	case Beta:
		return e.Beta, nil
	case Prod:
		return e.Prod, nil
	default:
		return Config{}, EnvError
	}
}
