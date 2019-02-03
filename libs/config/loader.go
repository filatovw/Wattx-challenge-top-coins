package config

import (
	"encoding/json"
	"fmt"

	consul "github.com/hashicorp/consul/api"
	"github.com/pkg/errors"
)

type ServiceConfig interface {
	ServiceName() string
}

func LoadConfig(dest ServiceConfig) error {
	consulHost := "consul-agent:8500"
	// Get a new client
	defaultCfg := consul.DefaultConfig()
	defaultCfg.Address = consulHost
	client, err := consul.NewClient(defaultCfg)
	if err != nil {
		return errors.Wrapf(err, "LoadConfig, host: %s", consulHost)
	}

	kv := client.KV()
	key := fmt.Sprintf("app/%s/app.json", dest.ServiceName())
	if err := load(kv, key, dest); err != nil {
		return errors.Wrapf(err, "LoadConfig, key: %s", key)
	}
	return nil
}

func load(kv *consul.KV, key string, dest ServiceConfig) error {
	data, _, err := kv.Get(key, nil)
	if err != nil {
		return err
	}
	if data == nil {
		return errors.New("Key not found")
	}
	return json.Unmarshal(data.Value, &dest)
}
