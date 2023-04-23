package main

import (
	"flag"
	"fmt"

	"github.com/spf13/viper"
)

// Config represents the configuration values
type Config struct {
	Host string
	Port int
	// Add more configuration values as needed
}

var cfg Config

func init() {

	// Set default
	viper.SetDefault("Host", "localhost")
	viper.SetDefault("Port", 1948)

	// Set up global viper instance
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			panic(fmt.Errorf("error reading config file: %s", err))
		}
	}
	viper.Unmarshal(&cfg)
}

func parseFlags() {
	// Parse command-line options
	host := flag.String("host", "", "Host value")
	port := flag.Int("port", 0, "Port number")
	flag.Parse()

	// Read environment variables
	viper.SetEnvPrefix("MYAPP")
	viper.AutomaticEnv()
	hostEnv := viper.GetString("SERVER_HOST")
	portEnv := viper.GetInt("SERVER_PORT")

	// Determine host value
	if *host != "" {
		cfg.Host = *host
	} else if hostEnv != "" {
		cfg.Host = hostEnv
	}

	if *port != 0 {
		cfg.Port = *port
	} else if portEnv != 0 {
		cfg.Port = portEnv
	}
}
