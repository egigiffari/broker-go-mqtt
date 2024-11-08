package pkg

import (
	"fmt"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Config struct {
	Host                 string
	Port                 uint
	ClientId             string
	Username             string
	Password             string
	PingTimeout          uint
	KeepAlive            uint
	ReconnectInterval    uint
	MaxReconnectInterval uint
}

func (c *Config) Validate() error {
	if c.Host == "" {
		return ErrRequireHost
	}

	if c.Port == 0 {
		return ErrInvalidPort
	}

	if c.ClientId == "" {
		return ErrRequireClientId
	}

	if c.Username == "" {
		return ErrRequireUsername
	}

	if c.Password == "" {
		return ErrRequirePassword
	}

	if c.ReconnectInterval < 5 {
		return ErrInvalidReconnectInterval
	}

	if c.MaxReconnectInterval < 1 {
		return ErrInvalidMaxReconnectInterval
	}

	return nil
}

func (c *Config) Get() *mqtt.ClientOptions {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(c.dns()).
		SetClientID(c.ClientId).
		SetKeepAlive(60 * time.Second).
		SetPingTimeout(5 * time.Second).
		SetAutoReconnect(true).
		SetConnectRetry(true).
		SetConnectRetryInterval(time.Duration(c.ReconnectInterval) * time.Second).
		SetMaxReconnectInterval(time.Duration(c.MaxReconnectInterval) * time.Hour)
	return opts
}

func (c *Config) dns() string {
	return fmt.Sprintf("tcp://%s:%d", c.Host, c.Port)
}
