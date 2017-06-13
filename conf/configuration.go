package conf

import (
	"encoding/json"
	"os"
)

type Conf struct {
	UpstreamPort      string
	DownstreamPort    string
	ReadLimit         uint16
	WriteLimit        uint16
	ConnTimeout       uint16
	ConnCheckInterval uint16
	ServerStatistics  uint16
	HttpAddr          string
}

type Configuration struct {
	Configure *Conf
}

func ReadConfig(confpath string) (*Configuration, error) {
	file, _ := os.Open(confpath)
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)

	return &configuration, err
}

var Config *Conf

func SetConfiguration(config *Conf) {
	Config = config
}

func GetConfiguration() *Conf {
	return Config
}
