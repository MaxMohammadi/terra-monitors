# terra-monitors

There are required three .env files to start service:
`./docker/env/.telegram.env` - telegramm bot credentials
`./docker/env/.lido_terra.env` - terra-monitor service settings
`./docker/env/.grafana.env` - grafana related variables

Before running the service you have to fill .env file `./docker/env/.telegram.env` with the following variables:
```shell
cat ./docker/env/.telegram.env 
TELEGRAM_BOTTOKEN=<telegram_bottoken>
TELEGRAM_CHAT_ID=<telegram_chat_id>
```

Before running the service you have to fill .env file `./docker/env/.lido_terra.env` with the following variables:
```shell
cat ./docker/env/.lido_terra.env
# Grafana frontend port - optional value, default value is 3000
GRAFANA_PORT=3000 

# LCD - light client daemon, https://docs.terra.money/terracli/lcd.html
LCD_ENDPOINT=fcd.terra.dev
LCD_SCHEMES=https

# terra-monitor data update interval
UPDATE_DATA_INTERVAL=30s

# monitored contracts
ADDRESSES_HUB_CONTRACT=terra1mtwph2juhj0rvjz7dy92gvl6xvukaxu8rfv8ts
ADDRESSES_REWARD_CONTRACT=terra17yap3mhph35pcwvhza38c2lkj7gzywzy05h7l0
ADDRESSES_BLUNA_TOKEN_INFO_CONTRACT=terra1kc87mu460fwkqte29rquh4hc20m54fxwtsx7gp
ADDRESSES_VALIDATORS_REGISTRY_CONTRACT=terra_dummy_validators_registry
ADDRESSES_REWARDS_DISPATCHER_CONTRACT=terra_dummy_rewards_dispatcher
ADDRESSES_AIR_DROP_REGISTRY_CONTRACT=terra_dummy_airdrop

# monitored bot, executing update_global_index message on the hub contract
# https://www.notion.so/bAsset-index-updating-bot-f64ebb5ec6704f05a840d93f28b1e3be
ADDRESSES_UPDATE_GLOBAL_INDEX_BOT_ADDRESS=terra1eqpx4zr2vm9jwu2vas5rh6704f6zzglsayf2fy
```

To set relevant grafana address in alert messages, put env variable `GF_SERVER_ROOT_URL` with URL in `./docker/env/.grafana.env` file
```shell
cat ./docker/env/.grafana.env 
GF_SERVER_ROOT_URL=http://some-host/
```


To run the service
```shell
make start
```

To stop the service:
```shell
make stop
```

Grafana dashboards avaliable at http://127.0.0.1:3000.
To change default port of the grafana host, pass variable `GRAFANA_PORT` to start command, lile:
```shell
GRAFANA_PORT=3001 make start
#or
GRAFANA_PORT=3001 docker-compose up -d
```

Default login/pass: `admin/admin`.

Directory `openapi` is generated by `go-swagger` tool according to the instruction at https://goswagger.io/generate/client.html from the `swagger.yaml` file.

To install the `go-swagger` CLI tool run the following command: 
```shell
go install github.com/go-swagger/go-swagger/cmd/swagger@latest`
```
