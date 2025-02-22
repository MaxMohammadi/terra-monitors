package config

import (
	"fmt"
	"time"

	"github.com/lidofinance/terra-monitors/openapi/client"
	"github.com/vrischmann/envconfig"
)

const (
	V1Contracts = "1"
	V2Contracts = "2"
)

type CollectorConfig struct {
	BassetContractsVersion string `envconfig:"default=2"` // available values: 1 and 2
	LCD                    LCD
	Addresses              Addresses
	UpdateDataInterval     time.Duration `envconfig:"default=30s"`
}

func NewCollectorConfig() (CollectorConfig, error) {
	config := CollectorConfig{}
	if err := envconfig.Init(&config); err != nil {
		return config, fmt.Errorf("failed to init config: %w", err)
	}

	return config, nil
}

type LCD struct {
	Endpoint string   `envconfig:"default=fcd.terra.dev"`
	Schemes  []string `envconfig:"default=https"`
}

type Addresses struct {
	HubContract                 string `envconfig:"default=terra1mtwph2juhj0rvjz7dy92gvl6xvukaxu8rfv8ts"`
	RewardContract              string `envconfig:"default=terra17yap3mhph35pcwvhza38c2lkj7gzywzy05h7l0"`
	BlunaTokenInfoContract      string `envconfig:"default=terra1kc87mu460fwkqte29rquh4hc20m54fxwtsx7gp"`
	ValidatorsRegistryContract  string `envconfig:"default=terra_dummy_validators_registry"` // TODO: actualize.
	RewardsDispatcherContract   string `envconfig:"default=terra_dummy_rewards_dispatcher"`  // TODO: actualize.
	AirDropRegistryContract     string `envconfig:"default=terra_dummy_airdrop"`             // TODO: actualize.
	UpdateGlobalIndexBotAddress string `envconfig:"default=terra1eqpx4zr2vm9jwu2vas5rh6704f6zzglsayf2fy"`
}

func (c CollectorConfig) GetTerraClient() *client.TerraLiteForTerra {
	transportConfig := &client.TransportConfig{
		Host:     c.LCD.Endpoint,
		Schemes:  c.LCD.Schemes,
		BasePath: client.DefaultBasePath,
	}

	return client.NewHTTPClientWithConfig(nil, transportConfig)
}
