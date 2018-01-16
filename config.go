package main

import (
	"github.com/spf13/viper"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type Config struct {
	Address string
	Port    int
	SSLOnly bool

	CertificatePath string
	KeyPath         string
}

func (c Config) Validate() error {
	if c.Port == 0 {
		return errors.New("missing port")
	}

	return nil
}

func ReadConfig(fileName string) (*Config, error) {
	viper.SetConfigName(fileName) // name of config file (without extension)
	viper.AddConfigPath(".")      // optionally look for config in the working directory

	err := viper.ReadInConfig()   // Find and read the config file
	if err != nil {               // Handle errors reading the config file
		return nil, errors.Wrap(err, "Error reading config from file")
	}

	config := Config{
		Address:         viper.GetString("address"),
		Port:            viper.GetInt("port"),
		SSLOnly:         viper.GetBool("ssl_only"),
		CertificatePath: viper.GetString("certificate_path"),
		KeyPath:         viper.GetString("key_path"),
	}

	if err = config.Validate(); err != nil {
		return nil, err
	}

	return &config, nil
}
