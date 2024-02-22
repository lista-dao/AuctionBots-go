# AuctionBots-go

## Usage

### Build
Build the project with the following command.
```shell
CGO_ENABLED=0 go build -o auctionbots cmd/bot/main.go
```

### Configuration

Before running, you need to fill in some parameters of the configuration file.

In this project, the configuration file is located in "config/example.config.yaml".

You can rename this file to "config/config.yaml" and fill in the parameters.
```yaml
wallet:
  privateKey: "<your private key>"
rpcNode:
  ws: "<bsc ws node>"
  http: "<bsc http node>"
```

### Run
Run with the configuration file.
```shell
./auctionbots -config config/config.yaml
```

### Run with Docker
The configuration file also needs to be filled in before running.
```shell
docker-compose up -d
```