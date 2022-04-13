package chain

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

// Info from config file
type Config struct {
	RpcEndpoint string
	Secret      string
	Address     string
	fname       string
}

// Reads info from config file
func ReadConfig(fname string) Config {
	_, err := os.Stat(fname)
	if err != nil {
		log.Fatal("Config file is missing: ", fname)
	}

	var config Config
	_, err = toml.DecodeFile(fname, &config)
	if err != nil {
		log.Fatal(err)
	}
	config.fname = fname
	return config
}

func WriteConfig(config *Config) {
	f, err := os.Create(config.fname)
	defer f.Close()
	if err != nil {
		log.Fatal("Cannot create or update config file: ", config.fname)
	}
	encoder := toml.NewEncoder(f)
	encoder.Encode(config)
	_ = encoder
}
