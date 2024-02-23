# AuctionBots-go

## Usage

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
You can get rpcNode from [Alchemy](https://www.alchemy.com/) or [Infura](https://www.infura.io/).


### Run
You can get the binary file from the release page or build it yourself.

Here is the link to the release page: [Releases](https://github.com/lista-dao/AuctionBots-go/releases/).

Run with the configuration file.

For example:

If your configuration file is located in "./config/config.yaml", you can run the following command:
```shell
./auctionbots -config ./config/config.yaml
```
